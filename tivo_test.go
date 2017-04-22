package tivo

import (
	"encoding/xml"
	"sort"
	"testing"
	"time"
)

func TestContainerItemXMLUnmarshal(t *testing.T) {
	container := new(Container)
	err := xml.Unmarshal([]byte(nowPlayingXML), container)
	if err != nil {
		t.Errorf("Failed to Unmarshal Container: %s", err)
	}

	vdRoot := VideoDetailRoot{}
	err = xml.Unmarshal([]byte(movieDetailXML), &vdRoot)
	if err != nil {
		t.Errorf("Failed to Unmarshal VideoDetail: %s", err)
		return
	}
	detail := vdRoot.Showing

	folder(t, container.Items[0])
	tvContainerItem(t, container.Items[12])
	movieVideoDetail(t, detail)

	container = new(Container)
	err = xml.Unmarshal([]byte(containerInProgressXML), container)
	if err != nil {
		t.Errorf("Failed to Unmarshal Container: %s", err)
	}
	tvContainerItemInProgress(t, container.Items[0])

	vdRoot = VideoDetailRoot{}
	err = xml.Unmarshal([]byte(tvDetailXML), &vdRoot)
	if err != nil {
		t.Errorf("Failed to Unmarshal VideoDetail: %s", err)
		return
	}
	detail = vdRoot.Showing
	tvVideoDetail(t, detail)
}

func tvVideoDetail(t *testing.T, detail VideoDetail) {
	var (
		expectedString string
		expectedDesc   description
		expectedInt    int
		expectedBool   bool
		expectedGenres Genres
		expectedPeople People
	)

	expectedString = "1998-05-12T01:00:00Z"
	if detail.OriginalAirDate != expectedString {
		t.Errorf(`Expected OriginalAirDate to be %s, not %s`, expectedString, detail.OriginalAirDate)
	}

	expectedString = "Becoming"
	if detail.EpisodeTitle != expectedString {
		t.Errorf(`Expected EpisodeTitle to be %s, not %s`, expectedString, detail.EpisodeTitle)
	}

	expectedString = "221"
	if detail.EpisodeNumber != expectedString {
		t.Errorf(`Expected EpisodeNumber to be %d, not %d`, expectedString, detail.EpisodeNumber)
	}

	sort.Sort(detail.Actors)
	expectedPeople = People{
		"Gellar|Sarah Michelle",
		"Brendon|Nicholas",
		"Hannigan|Alyson",
		"Carpenter|Charisma",
		"Head|Anthony Stewart",
		"Boreanaz|David",
	}
	for _, p := range expectedPeople {
		f := func(i int) bool {
			return detail.Actors[i] >= p
		}
		i := sort.Search(len(detail.Actors), f)
		if !(i < len(detail.Actors) && detail.Actors[i] == p) {
			t.Errorf(`Expected Actors to contain %s`, p)
		}
	}

	expectedDesc = description("While Angel attempts to unearth a demon who will annihilate the planet, Buffy reluctantly decides to destroy her former lover; Willow unlocks a secret.")
	if detail.Description != expectedDesc {
		t.Errorf(`Expected Description to be %s, not %s`, expectedDesc, detail.Description)
	}

	expectedPeople = People{
		"Whedon|Joss",
	}
	for _, p := range expectedPeople {
		f := func(i int) bool {
			return detail.Directors[i] >= p
		}
		i := sort.Search(len(detail.Directors), f)
		if !(i < len(detail.Directors) && detail.Directors[i] == p) {
			t.Errorf(`Expected Directors to contain %s`, p)
		}
	}

	sort.Sort(detail.ExecProducers)
	expectedPeople = People{
		"Whedon|Joss",
		"Berman|Gail",
		"Gallin|Sandy",
		"Kuzui|Fran",
		"Kuzui|Kaz",
		"Greenwalt|David",
	}
	for _, p := range expectedPeople {
		f := func(i int) bool {
			return detail.ExecProducers[i] >= p
		}
		i := sort.Search(len(detail.ExecProducers), f)
		if !(i < len(detail.ExecProducers) && detail.ExecProducers[i] == p) {
			t.Errorf(`Expected ExecProducers to contain %s`, p)
		}
	}

	expectedBool = true
	if detail.IsEpisode != expectedBool {
		t.Errorf(`Expected IsEpisode to be %t, not %t`, expectedBool, detail.IsEpisode)
	}

	expectedBool = true
	if detail.IsEpisodic != expectedBool {
		t.Errorf(`Expected IsEpisodic to be %t, not %t`, expectedBool, detail.IsEpisodic)
	}

	expectedPeople = People{
		"Davies|Gareth",
	}
	for _, p := range expectedPeople {
		f := func(i int) bool {
			return detail.Producers[i] >= p
		}
		i := sort.Search(len(detail.Producers), f)
		if !(i < len(detail.Producers) && detail.Producers[i] == p) {
			t.Errorf(`Expected Producers to contain %s`, p)
		}
	}

	sort.Sort(detail.SeriesGenres)
	expectedGenres = Genres{
		"Drama",
		"Fantasy",
		"Horror",
		"Drama",
		"Sci-Fi and Fantasy",
	}
	for _, g := range expectedGenres {
		f := func(i int) bool {
			return detail.SeriesGenres[i] >= g
		}
		i := sort.Search(len(detail.SeriesGenres), f)
		if !(i < len(detail.SeriesGenres) && detail.SeriesGenres[i] == g) {
			t.Errorf(`Expected SeriesGenres to contain %s`, g)
		}
	}

	expectedString = "Buffy the Vampire Slayer"
	if detail.Title != expectedString {
		t.Errorf(`Expected Title to be %s, not %s`, expectedString, detail.Title)
	}

	expectedPeople = People{
		"Whedon|Joss",
	}
	for _, p := range expectedPeople {
		f := func(i int) bool {
			return detail.Writers[i] >= p
		}
		i := sort.Search(len(detail.Writers), f)
		if !(i < len(detail.Writers) && detail.Writers[i] == p) {
			t.Errorf(`Expected Writers to contain %s`, p)
		}
	}

	expectedInt = 1
	if detail.PartIndex != expectedInt {
		t.Errorf(`Expected PartIndex to be %d, not %d`, expectedInt, detail.PartIndex)
	}

	expectedInt = 2
	if detail.PartCount != expectedInt {
		t.Errorf(`Expected PartCount to be %d, not %d`, expectedInt, detail.PartCount)
	}

}

// Test XML unmarshaling of TiVoContainer items with InProgress present
func tvContainerItemInProgress(t *testing.T, item ContainerItem) {
	var expectedBool yesNoBool
	expectedBool = true
	if item.InProgress != expectedBool {
		t.Errorf(`Expected InProgress to be %t, not %t`, expectedBool, item.InProgress)
	}

	expectedBool = true
	if item.HighDefinition != expectedBool {
		t.Errorf(`Expected HighDefinition to be %t, not %t`, expectedBool, item.HighDefinition)
	}
}

// Test XML unmarshaling of video details.
func movieVideoDetail(t *testing.T, detail VideoDetail) {
	var (
		expectedString string
		expectedDesc   description
		expectedTime   myTime
		expectedInt    int
		expectedBool   bool
		expectedGenres Genres
		expectedPeople People
	)
	expectedBool = true
	if detail.IsEpisode != expectedBool {
		t.Errorf(`Expected IsEpisode to be %t, not %t`, expectedBool, detail.IsEpisode)
	}

	expectedBool = false
	if detail.IsEpisodic != expectedBool {
		t.Errorf(`Expected IsEpisodic to be %t, not %t`, expectedBool, detail.IsEpisodic)
	}

	expectedString = "Bridge to Terabithia"
	if detail.Title != expectedString {
		t.Errorf(`Expected Title to be %s, not %s`, expectedString, detail.Title)
	}

	expectedString = "Bridge to Terabithia"
	if detail.SeriesTitle != expectedString {
		t.Errorf(`Expected SeriesTitle to be %s, not %s`, expectedString, detail.SeriesTitle)
	}

	expectedDesc = description(`A boy and his new friend, the class outsider, create an imaginary world in which they rule as king and queen. Based on the novel by Katherine Paterson.`)
	if detail.Description != expectedDesc {
		t.Errorf(`Expected Description to be %s, not %s`, expectedDesc, detail.Description)
	}

	// No XML input to test the following:
	//   GuestStars
	//   Choreographers
	//   Hosts

	date, err := time.Parse("2006-01-02T15:04:05Z", "2008-01-08T11:30:00Z")
	if err != nil {
		t.Errorf(`Failed to parse time: %s`, err)
	}
	expectedTime = myTime{date}
	if detail.Time != expectedTime {
		t.Errorf(`Expected Time to be %s, not %s`, expectedTime, detail.Time)
	}
	if detail.Time.String() != expectedTime.String() {
		t.Errorf(`Expected Time string to be %s, not %s`, expectedTime, detail.Time)
	}

	expectedInt = 2007
	if detail.MovieYear != expectedInt {
		t.Errorf(`Expected MovieYear to be %d, not %d`, expectedInt, detail.MovieYear)
	}

	expectedInt = 2007
	if detail.MovieYear != expectedInt {
		t.Errorf(`Expected MovieYear to be %d, not %d`, expectedInt, detail.MovieYear)
	}

	sort.Sort(detail.SeriesGenres)
	expectedGenres = Genres{"Fantasy", "Movies", "Sci-Fi and Fantasy"}
	for _, g := range expectedGenres {
		f := func(i int) bool {
			return detail.SeriesGenres[i] >= g
		}
		i := sort.Search(len(detail.SeriesGenres), f)
		if !(i < len(detail.SeriesGenres) && detail.SeriesGenres[i] == g) {
			t.Errorf(`Expected SeriesGenres to contain %s`, g)
		}
	}

	sort.Sort(detail.Actors)
	expectedPeople = People{
		"Hutcherson|Josh",
		"Robb|AnnaSophia",
		"Deschanel|Zooey",
		"Patrick|Robert",
		"Clinton|Lauren",
		"Cerio|Katrina",
		"Madison|Bailee",
		"Wood|Devon",
		"Wakefield|Cameron",
		"Lawless|Elliott",
		"Kircher|Isabelle Rose",
		"Gaines|Latham",
		"Owen|Carly",
	}
	for _, p := range expectedPeople {
		f := func(i int) bool {
			return detail.Actors[i] >= p
		}
		i := sort.Search(len(detail.Actors), f)
		if !(i < len(detail.Actors) && detail.Actors[i] == p) {
			t.Errorf(`Expected Actors to contain %s`, p)
		}
	}

	expectedPeople = People{
		"Csupo|Gabor",
	}
	for _, p := range expectedPeople {
		f := func(i int) bool {
			return detail.Directors[i] >= p
		}
		i := sort.Search(len(detail.Directors), f)
		if !(i < len(detail.Directors) && detail.Directors[i] == p) {
			t.Errorf(`Expected Directors to contain %s`, p)
		}
	}

	sort.Sort(detail.Producers)
	expectedPeople = People{
		"Lieberman|Hal",
		"Levine|Lauren",
	}
	for _, p := range expectedPeople {
		f := func(i int) bool {
			return detail.Producers[i] >= p
		}
		i := sort.Search(len(detail.Producers), f)
		if !(i < len(detail.Producers) && detail.Producers[i] == p) {
			t.Errorf(`Expected Producers to contain %s`, p)
		}
	}

	sort.Sort(detail.Writers)
	expectedPeople = People{
		"Paterson|David",
		"Stockwell|Jeff",
		"Wade|Kevin",
	}
	for _, p := range expectedPeople {
		f := func(i int) bool {
			return detail.Writers[i] >= p
		}
		i := sort.Search(len(detail.Writers), f)
		if !(i < len(detail.Writers) && detail.Writers[i] == p) {
			t.Errorf(`Expected Writers to contain %s`, p)
		}
	}
}

// Test XML unmarshaling of TiVoContainer items
func tvContainerItem(t *testing.T, item ContainerItem) {
	var (
		expectedString  string
		expectedInt     int
		expectedInt64   int64
		expectedHexDate hexDate
		expectedBool    yesNoBool
	)

	expectedString = "video/x-tivo-raw-pes"
	if item.ContentType != expectedString {
		t.Errorf(`Expected ContentType to be %s, not %s`, expectedString, item.ContentType)
	}

	expectedString = "video/x-tivo-raw-pes"
	if item.ContentType != expectedString {
		t.Errorf(`Expected ContentType to be %s, not %s`, expectedString, item.ContentType)
	}

	expectedString = "Person of Interest"
	if item.Title != expectedString {
		t.Errorf(`Expected Title to be %s, not %s`, expectedString, item.Title)
	}

	expectedInt64 = 1694498816
	if item.SourceSize != expectedInt64 {
		t.Errorf(`Expected SourceSize to be %d, not %d`, expectedInt64, item.SourceSize)
	}

	expectedInt = 3601000
	if item.Duration != expectedInt {
		t.Errorf(`Expected Duration to be %d, not %d`, expectedInt64, item.Duration)
	}

	date, err := time.Parse("2006-01-02 15:04:05 -0700 MST", "2012-10-25 20:00:58 -0700 MST")
	if err != nil {
		t.Errorf(`Failed to parse time: %s`, err)
	}
	expectedHexDate = hexDate{date}
	if item.CaptureDate != expectedHexDate {
		t.Errorf(`Expected CaptureDate to be %s, not %s`, expectedHexDate, item.CaptureDate)
	}

	expectedString = "Triggerman"
	if item.EpisodeTitle != expectedString {
		t.Errorf(`Expected EpisodeTitle to be %s, not %s`, expectedString, item.EpisodeTitle)
	}

	expectedString = `Reese and Finch wonder if they should intervene when they learn that a mob enforcer's life is in danger; Finch turns to an unexpected source for help. Copyright Tribune Media Services, Inc.`
	if item.Description != expectedString {
		t.Errorf(`Expected Description to be %s, not %s`, expectedString, item.Description)
	}

	expectedString = "5"
	if item.SourceChannel != expectedString {
		t.Errorf(`Expected SourceChannel to be %s, not %s`, expectedString, item.SourceChannel)
	}

	expectedString = "KPHO"
	if item.SourceStation != expectedString {
		t.Errorf(`Expected SourceStation to be %s, not %s`, expectedString, item.SourceStation)
	}

	expectedBool = false
	if item.HighDefinition != expectedBool {
		t.Errorf(`Expected HighDefinition to be %t, not %t`, expectedBool, item.HighDefinition)
	}

	expectedString = "EP014198470027"
	if item.ProgramID != expectedString {
		t.Errorf(`Expected ProgramID to be %s, not %s`, expectedString, item.ProgramID)
	}

	expectedString = "SH01419847"
	if item.SeriesID != expectedString {
		t.Errorf(`Expected SeriesID to be %s, not %s`, expectedString, item.SeriesID)
	}

	expectedInt = 204
	if item.EpisodeNumber != expectedInt {
		t.Errorf(`Expected EpisodeNumber to be %d, not %d`, expectedInt, item.EpisodeNumber)
	}

	// ByteOffset is zero in all the input I've seen, so no tests yet.

	expectedString = `http://10.0.0.110:80/download/Person%20of%20Interest.TiVo?Container=%2FNowPlaying&id=2143738`
	if item.ContentURL != expectedString {
		t.Errorf(`Expected ContentURL to be %s, not %s`, expectedString, item.ContentURL)
	}

	expectedString = `video/x-tivo-raw-pes`
	if item.ContentTypeURL != expectedString {
		t.Errorf(`Expected ContentTypeURL to be %s, not %s`, expectedString, item.ContentTypeURL)
	}

	expectedString = `urn:tivo:image:save-until-i-delete-recording`
	if item.CustomIconURL != expectedString {
		t.Errorf(`Expected CustomIconURL to be %s, not %s`, expectedString, item.CustomIconURL)
	}

	expectedString = `https://10.0.0.110:443/TiVoVideoDetails?id=2143738`
	if item.VideoDetailsURL != expectedString {
		t.Errorf(`Expected VideoDetailsURL to be %s, not %s`, expectedString, item.VideoDetailsURL)
	}
}

// Test unmarshaling of folders
func folder(t *testing.T, item ContainerItem) {
	expected := "x-tivo-container/folder"
	if item.ContentType != expected {
		t.Errorf(`Expected ContentType to be %s, not %s`, expected, item.ContentType)
	}
	expected = "x-tivo-container/tivo-dvr"
	if item.SourceFormat != expected {
		t.Errorf(`Expected ContentType to be %s, not %s`, expected, item.SourceFormat)
	}
	expected = "BECKINSALE, KATE"
	if item.Title != expected {
		t.Errorf(`Expected Title to be %s, not %s`, expected, item.Title)
	}
}
