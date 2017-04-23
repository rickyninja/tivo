package tivo

import (
	"crypto/md5"
	"crypto/tls"
	"encoding/xml"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"

	cache "github.com/robfig/go-cache"
)

var (
	ErrCache = errors.New("cache error")
)

type Client struct {
	Debug     bool
	BaseURI   string
	MAK       int
	Realm     string
	Login     string
	Host      string
	Protocol  string
	UseCache  bool
	Cache     *cache.Cache
	CacheFile string
	*http.Client
}

func NewClient(host string, login string, protocol string, mak int, cachefile string) (*Client, error) {
	c := cache.New(time.Second*3600, time.Minute*60)
	if _, err := os.Stat(cachefile); err == nil {
		err := c.LoadFile(cachefile)
		if err != nil {
			return nil, err
		}
	}

	// downloads get interrupted when the timeout is too short
	timeout := time.Duration(0 * time.Second)
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	jar, err := cookiejar.New(&cookiejar.Options{})
	if err != nil {
		return nil, err
	}

	client := &http.Client{
		Timeout:   timeout,
		Transport: tr,
		Jar:       jar,
	}

	baseuri := fmt.Sprintf("%s://%s/TiVoConnect", protocol, host)
	return &Client{
		Debug:     false,
		BaseURI:   baseuri,
		MAK:       mak,
		Realm:     "TiVo DVR",
		Login:     login,
		Host:      host,
		Protocol:  protocol,
		UseCache:  false,
		Cache:     c,
		CacheFile: cachefile,
		Client:    client,
	}, nil
}

func (c *Client) FetchData(uri *url.URL) ([]byte, error) {
	data, found := c.Cache.Get(uri.String())

	if !found || !c.UseCache {
		if c.Debug {
			log.Print("cache miss: " + uri.String() + "\n")
		}
		response, err := c.Go(uri)
		if err != nil {
			return nil, err
		}
		if response.StatusCode == 200 {
			var err error
			data, err = ioutil.ReadAll(response.Body)
			if err != nil {
				return nil, err
			}
			c.Cache.Set(uri.String(), data, 0)
		}
	} else {
		if c.Debug {
			fmt.Printf("cache hit: %s\n", uri.String())
		}
	}

	return data.([]byte), nil
}

func (c *Client) QueryContainer(param map[string]string) ([]ContainerItem, error) {
	param["Command"] = "QueryContainer"
	uri, err := c.GetURI(param)
	if err != nil {
		return nil, err
	}

	xmldata, err := c.FetchData(uri)
	if err != nil {
		return nil, err
	}
	container := &Container{}
	err = xml.Unmarshal(xmldata, container)
	if err != nil {
		return nil, err
	}
	return container.Items, nil
}

func (c *Client) GetDetail(ci ContainerItem) (VideoDetail, error) {
	vd := VideoDetail{}
	uri, err := url.Parse(ci.VideoDetailsURL)
	if err != nil {
		return vd, err
	}

	xmldata, err := c.FetchData(uri)
	if err != nil {
		return vd, err
	}
	root := &VideoDetailRoot{}
	err = xml.Unmarshal(xmldata, root)
	if err != nil {
		return vd, err
	}
	return root.Showing, nil
}

func (c *Client) DigestAuth(response *http.Response) (*http.Response, error) {
	var (
		err     error
		request *http.Request
	)

	request, err = http.NewRequest("GET", response.Request.URL.String(), nil)
	if err != nil {
		return response, err
	}

	digest := fmt.Sprintf(`Digest username="%s", `, c.Login)

	var re *regexp.Regexp
	var result []string

	re, err = regexp.Compile(`Digest realm="([^"]+)"`)
	if err != nil {
		return response, err
	}
	result = re.FindStringSubmatch(response.Header.Get("Www-Authenticate"))
	realm := result[1]
	digest += fmt.Sprintf(`realm="%s", `, realm)

	re, err = regexp.Compile(`nonce="([^"]+)"`)
	if err != nil {
		return response, err
	}
	result = re.FindStringSubmatch(response.Header.Get("Www-Authenticate"))
	nonce := result[1]
	digest += fmt.Sprintf(`nonce="%s", `, nonce)

	re, err = regexp.Compile(`qop="([^"]+)"`)
	if err != nil {
		return response, err
	}
	result = re.FindStringSubmatch(response.Header.Get("Www-Authenticate"))
	qop := result[1]
	digest += fmt.Sprintf(`qop=%s, `, qop)

	digest += `algorithm="MD5", `

	digest += fmt.Sprintf(`uri="%s", `, response.Request.URL.RequestURI())

	nc := fmt.Sprintf("%08X", 1)
	digest += fmt.Sprintf(`nc="%s", `, nc)

	cnonce := fmt.Sprintf("%8x", time.Now().Unix())
	digest += fmt.Sprintf(`cnonce="%s", `, cnonce)

	HA1 := c.HA1(c.Login, realm, strconv.Itoa(c.MAK))
	HA2 := c.HA2(response.Request.Method, response.Request.URL.RequestURI())

	resp := c.GetDigestResponse(HA1, HA2, nonce, nc, cnonce, qop)
	digest += fmt.Sprintf(`response="%s"`, resp)

	request.Header.Add("Authorization", digest)

	//response.Write(os.Stdout)
	//fmt.Println("")
	//request.Write(os.Stdout)

	authresponse, err := c.Do(request)
	if err != nil {
		return response, err
	}
	return authresponse, nil
}

func (c *Client) GetDigestResponse(ha1, ha2, nonce, nc, cnonce, qop string) string {
	s := []string{ha1, nonce, nc, cnonce, qop, ha2}
	return h(strings.Join(s, ":"))
}

func (c *Client) HA2(method string, uri string) string {
	s := []string{method, uri}
	return h(strings.Join(s, ":"))
}

func (c *Client) HA1(username string, realm string, password string) string {
	s := []string{username, realm, password}
	return h(strings.Join(s, ":"))
}

func h(str string) string {
	sum := md5.Sum([]byte(str))
	return fmt.Sprintf("%x", sum)
}

func (c *Client) GetURI(param map[string]string) (*url.URL, error) {
	uri, err := url.Parse(c.BaseURI)
	if err != nil {
		return nil, err
	}

	p := url.Values{}
	for k, v := range param {
		p.Add(k, v)
	}
	uri.RawQuery = p.Encode()
	return uri, nil
}

func (c *Client) Go(uri *url.URL) (*http.Response, error) {
	request, err := http.NewRequest("GET", uri.String(), nil)
	if err != nil {
		return nil, err
	}

	response, err := c.Do(request)
	if err != nil {
		return nil, err
	}

	return c.CheckAuth(response)
}

// Can the Digest auth be hooked into the transport?
func (c *Client) CheckAuth(response *http.Response) (*http.Response, error) {
	if response.StatusCode == 200 {
		return response, nil
	} else if response.StatusCode == 401 {
		for i := 1; i < 10; i++ {
			authresponse, err := c.DigestAuth(response)
			if err != nil {
				return response, err
			}
			if authresponse.StatusCode == 200 {
				return authresponse, nil
			} else if authresponse.StatusCode == 503 {
				log.Print(fmt.Sprintf("%s : attempt %d/10, retrying\n", authresponse.Status, i))
				time.Sleep(2 * time.Second)
				continue
			} else {
				return authresponse, errors.New("Failed to authenticate to the tivo: " + authresponse.Status)
			}
		}
	}

	return response, nil
}

func (c *Client) WriteCache() error {
	err := c.Cache.SaveFile(c.CacheFile)
	if err != nil {
		return err
	}
	return nil
}
