# tivo
--
    import "github.com/rickyninja/tivo"

Package tivo is a client library for TiVo HMO service

I'm no longer a TiVo subscriber, and the API won't return anything unless you
have an active account. This means I can no longer test code changes.

## Usage

```go
var (
	ErrCache = errors.New("cache error")
)
```

#### type Client

```go
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
```

Client is a tivo client.

#### func  NewClient

```go
func NewClient(host string, login string, protocol string, mak int, cachefile string) (*Client, error)
```
NewClient returns a ready to use Client.

#### func (*Client) Get

```go
func (c *Client) Get(uri *url.URL) (*http.Response, error)
```
Get does an HTTP GET request to tivo.

#### func (*Client) GetCache

```go
func (c *Client) GetCache(uri *url.URL) ([]byte, error)
```
GetCache does an HTTP GET on the uri, and caches the response body.

#### func (*Client) GetDetail

```go
func (c *Client) GetDetail(videoDetailsURL string) (VideoDetail, error)
```
GetDetail queries tivo for details of a recording with the link provided.

#### func (*Client) QueryContainer

```go
func (c *Client) QueryContainer(param map[string]string) ([]ContainerItem, error)
```
QueryContainer queries tivo for items in the Container specified in param.

#### func (*Client) WriteCache

```go
func (c *Client) WriteCache() error
```
WriteCache writes cache contents to disk.

#### type Container

```go
type Container struct {
	XMLName xml.Name        `xml:"TiVoContainer"`
	Items   []ContainerItem `xml:"Item"`
}
```

TiVoContainer represents the response returned by QueryContainer.

#### type ContainerItem

```go
type ContainerItem struct {
	InProgress      yesNoBool `xml:"Details>InProgress"`
	ContentType     string    `xml:"Details>ContentType"`
	SourceFormat    string    `xml:"Details>SourceFormat"`
	Title           string    `xml:"Details>Title"`
	SourceSize      int64     `xml:"Details>SourceSize"`
	Duration        int       `xml:"Details>Duration"`
	CaptureDate     hexDate   `xml:"Details>CaptureDate"`
	EpisodeTitle    string    `xml:"Details>EpisodeTitle"`
	Description     string    `xml:"Details>Description"`
	SourceChannel   string    `xml:"Details>SourceChannel"`
	SourceStation   string    `xml:"Details>SourceStation"`
	HighDefinition  yesNoBool `xml:"Details>HighDefinition"`
	ProgramID       string    `xml:"Details>ProgramId"`
	SeriesID        string    `xml:"Details>SeriesId"`
	EpisodeNumber   int       `xml:"Details>EpisodeNumber"`
	ByteOffset      int       `xml:"Details>ByteOffset"`
	ContentURL      string    `xml:"Links>Content>Url"`
	ContentTypeURL  string    `xml:"Links>Content>ContentType"`
	CustomIconURL   string    `xml:"Links>CustomIcon>Url"`
	VideoDetailsURL string    `xml:"Links>TiVoVideoDetails>Url"`
	Detail          VideoDetail
}
```

ContainerItem represents info about a tivo recording.

#### type Genre

```go
type Genre string
```

Genre represents a genre in responses from tivo.

#### type Genres

```go
type Genres []Genre
```

Genres represents multiple genres in responses from tivo, and is used to help
implement sort.Interface.

#### func (Genres) Len

```go
func (g Genres) Len() int
```
Len helps implement sort.Interface.

#### func (Genres) Less

```go
func (g Genres) Less(i, j int) bool
```
Less helps implement sort.Interface.

#### func (Genres) Swap

```go
func (g Genres) Swap(i, j int)
```
Swap helps implement sort.Interface.

#### type People

```go
type People []Person
```

People represents multiple persons in responses from tivo, and is used to help
implement sort.Interface.

#### func (People) Len

```go
func (p People) Len() int
```
Len helps implement sort.Interface.

#### func (People) Less

```go
func (p People) Less(i, j int) bool
```
Less helps implement sort.Interface.

#### func (People) Swap

```go
func (p People) Swap(i, j int)
```
Swap helps implement sort.Interface.

#### type Person

```go
type Person string
```

Person represents a person in responses from tivo.

#### type TvBusEnvelope

```go
type TvBusEnvelope struct {
	XMLName xml.Name `xml:"TvBusEnvelope"`
	Showing VideoDetail
}
```

TvBusEnvelope represents the response returned by GetDetail.

#### type VideoDetail

```go
type VideoDetail struct {
	XMLName         xml.Name    `xml:"showing"`
	IsEpisode       bool        `xml:"program>isEpisode"`
	IsEpisodic      bool        `xml:"program>series>isEpisodic"`
	Title           string      `xml:"program>title"`
	SeriesTitle     string      `xml:"program>series>seriesTitle"`
	Description     description `xml:"program>description"`
	OriginalAirDate string      `xml:"program>originalAirDate"`
	// episode Number is frequently incorrect in tivo metadata,
	// prefer the value acquired from TV meta data API.
	EpisodeNumber  string `xml:"program>episodeNumber"`
	EpisodeTitle   string `xml:"program>episodeTitle"`
	Time           myTime `xml:"time"`
	MovieYear      int    `xml:"program>movieYear"`
	PartCount      int    `xml:"partCount"`
	PartIndex      int    `xml:"partIndex"`
	SeriesGenres   Genres `xml:"program>series>vSeriesGenre>element"`
	Actors         People `xml:"program>vActor>element"`
	GuestStars     People `xml:"program>vGuestStar>element"`
	Directors      People `xml:"program>vDirector>element"`
	ExecProducers  People `xml:"program>vExecProducer>element"`
	Producers      People `xml:"program>vProducer>element"`
	Choreographers People `xml:"program>vChoreographer>element"`
	Writers        People `xml:"program>vWriter>element"`
	Hosts          People `xml:"program>vHost>element"`
}
```

VideoDetail represents detailed info about a tivo recording.
