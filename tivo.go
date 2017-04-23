/*
Package tivo is a client library for TiVo HMO service

I'm no longer a TiVo subscriber, and the API won't return anything unless you have an active account.
This means I can no longer test code changes.
*/
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

// Client is a tivo client.
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

// NewClient returns a ready to use Client.
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

// GetCache does an HTTP GET on the uri, and caches the response body.
func (c *Client) GetCache(uri *url.URL) ([]byte, error) {
	data, found := c.Cache.Get(uri.String())

	if !found || !c.UseCache {
		if c.Debug {
			log.Print("cache miss: " + uri.String() + "\n")
		}
		response, err := c.Get(uri)
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

// QueryContainer queries tivo for items in the Container specified in param.
func (c *Client) QueryContainer(param map[string]string) ([]ContainerItem, error) {
	param["Command"] = "QueryContainer"
	uri, err := c.baseURLWithQueryParam(param)
	if err != nil {
		return nil, err
	}

	xmldata, err := c.GetCache(uri)
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

// GetDetail queries tivo for details of a recording with the link provided.
func (c *Client) GetDetail(videoDetailsURL string) (VideoDetail, error) {
	vd := VideoDetail{}
	uri, err := url.Parse(videoDetailsURL)
	if err != nil {
		return vd, err
	}

	xmldata, err := c.GetCache(uri)
	if err != nil {
		return vd, err
	}
	tbe := &TvBusEnvelope{}
	err = xml.Unmarshal(xmldata, tbe)
	if err != nil {
		return vd, err
	}
	return tbe.Showing, nil
}

func (c *Client) digestAuth(response *http.Response) (*http.Response, error) {
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

	HA1 := c.ha1(c.Login, realm, strconv.Itoa(c.MAK))
	HA2 := c.ha2(response.Request.Method, response.Request.URL.RequestURI())

	resp := c.getDigestResponse(HA1, HA2, nonce, nc, cnonce, qop)
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

func (c *Client) getDigestResponse(ha1, ha2, nonce, nc, cnonce, qop string) string {
	s := []string{ha1, nonce, nc, cnonce, qop, ha2}
	return h(strings.Join(s, ":"))
}

func (c *Client) ha2(method string, uri string) string {
	s := []string{method, uri}
	return h(strings.Join(s, ":"))
}

func (c *Client) ha1(username string, realm string, password string) string {
	s := []string{username, realm, password}
	return h(strings.Join(s, ":"))
}

func h(str string) string {
	sum := md5.Sum([]byte(str))
	return fmt.Sprintf("%x", sum)
}

func (c *Client) baseURLWithQueryParam(param map[string]string) (*url.URL, error) {
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

// Get does an HTTP GET request to tivo.
func (c *Client) Get(uri *url.URL) (*http.Response, error) {
	request, err := http.NewRequest("GET", uri.String(), nil)
	if err != nil {
		return nil, err
	}

	response, err := c.Do(request)
	if err != nil {
		return nil, err
	}

	return c.checkAuth(response)
}

func (c *Client) checkAuth(response *http.Response) (*http.Response, error) {
	if response.StatusCode == 200 {
		return response, nil
	} else if response.StatusCode == 401 {
		for i := 1; i < 10; i++ {
			authresponse, err := c.digestAuth(response)
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

// WriteCache writes cache contents to disk.
func (c *Client) WriteCache() error {
	err := c.Cache.SaveFile(c.CacheFile)
	if err != nil {
		return err
	}
	return nil
}
