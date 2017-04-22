package tivo

import (
	"encoding/xml"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
)

type Container struct {
	XMLName xml.Name        `xml:"TiVoContainer"`
	Items   []ContainerItem `xml:"Item"`
}

type yesNoBool bool

func (m *yesNoBool) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var yesno string
	err := d.DecodeElement(&yesno, &start)
	if err != nil {
		return err
	}
	switch yesno {
	case "y", "Y", "yes", "YES", "Yes":
		*m = true
	case "n", "N", "no", "NO", "No":
		*m = false
	default:
		return errors.New(fmt.Sprintf("Failed to parse %s as a bool", yesno))
	}
	return nil
}

type hexDate struct {
	time.Time
}

func (m *hexDate) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var date string // eg: 0x51ED10AE -> 2013-07-22 03:59:58 -0700 MST
	err := d.DecodeElement(&date, &start)
	if err != nil {
		return err
	}
	epoch, err := strconv.ParseInt(date[2:], 16, 64)
	if err != nil {
		return err
	}
	*m = hexDate{time.Unix(epoch, 0)}
	return nil
}

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

type Person string
type People []Person

func (p People) Len() int {
	return len(p)
}

func (p People) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p People) Less(i, j int) bool {
	return p[i] < p[j]
}

type Genre string
type Genres []Genre

func (g Genres) Len() int {
	return len(g)
}

func (g Genres) Swap(i, j int) {
	g[i], g[j] = g[j], g[i]
}

func (g Genres) Less(i, j int) bool {
	return g[i] < g[j]
}

type VideoDetailRoot struct {
	XMLName xml.Name `xml:"TvBusEnvelope"`
	Showing VideoDetail
}

type description string

func (desc *description) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var str string
	err := d.DecodeElement(&str, &start)
	if err != nil {
		return err
	}
	str = strings.TrimSuffix(str, " Copyright Tribune Media Services, Inc.")
	*desc = description(str)
	return nil
}

type myTime struct {
	time.Time
}

func (t *myTime) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var date string // eg: 2008-01-08T11:30:00Z
	err := d.DecodeElement(&date, &start)
	if err != nil {
		return err
	}
	tt, err := time.Parse("2006-01-02T15:04:05Z", date)
	if err != nil {
		return err
	}
	*t = myTime{tt}
	return nil
}

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
