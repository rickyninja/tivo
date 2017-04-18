package tivo

import (
	"encoding/xml"
	"testing"
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
		expectedInt    int
		expectedBool   bool
		expectedGenre  []Genre
		expectedPerson []Person
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

	expectedPerson = []Person{
		"Gellar|Sarah Michelle",
		"Brendon|Nicholas",
		"Hannigan|Alyson",
		"Carpenter|Charisma",
		"Head|Anthony Stewart",
		"Boreanaz|David",
	}
	for _, p := range expectedPerson {
		if !inPerson(p, detail.Actors) {
			t.Errorf(`Expected Actors to contain %s`, p)
		}
	}

	expectedString = "While Angel attempts to unearth a demon who will annihilate the planet, Buffy reluctantly decides to destroy her former lover; Willow unlocks a secret.  Copyright Tribune Media Services, Inc."
	if detail.Description != expectedString {
		t.Errorf(`Expected Description to be %s, not %s`, expectedString, detail.Description)
	}

	expectedPerson = []Person{
		"Whedon|Joss",
	}
	for _, p := range expectedPerson {
		if !inPerson(p, detail.Directors) {
			t.Errorf(`Expected Directors to contain %s`, p)
		}
	}

	expectedPerson = []Person{
		"Whedon|Joss",
		"Berman|Gail",
		"Gallin|Sandy",
		"Kuzui|Fran",
		"Kuzui|Kaz",
		"Greenwalt|David",
	}
	for _, p := range expectedPerson {
		if !inPerson(p, detail.ExecProducers) {
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

	expectedPerson = []Person{
		"Davies|Gareth",
	}
	for _, p := range expectedPerson {
		if !inPerson(p, detail.Producers) {
			t.Errorf(`Expected Producers to contain %s`, p)
		}
	}

	expectedGenre = []Genre{
		"Drama",
		"Fantasy",
		"Horror",
		"Drama",
		"Sci-Fi and Fantasy",
	}
	for _, g := range expectedGenre {
		if !inGenre(g, detail.SeriesGenres) {
			t.Errorf(`Expected SeriesGenres to contain %s`, g)
		}
	}

	expectedString = "Buffy the Vampire Slayer"
	if detail.Title != expectedString {
		t.Errorf(`Expected Title to be %s, not %s`, expectedString, detail.Title)
	}

	expectedPerson = []Person{
		"Whedon|Joss",
	}
	for _, p := range expectedPerson {
		if !inPerson(p, detail.Writers) {
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
	var expectedString string
	expectedString = "Yes"
	if item.InProgress != expectedString {
		t.Errorf(`Expected InProgress to be %s, not %s`, expectedString, item.InProgress)
	}
}

// Test XML unmarshaling of video details.
func movieVideoDetail(t *testing.T, detail VideoDetail) {
	var (
		expectedString string
		expectedInt    int
		expectedBool   bool
		expectedGenre  []Genre
		expectedPerson []Person
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

	expectedString = `A boy and his new friend, the class outsider, create an imaginary world in which they rule as king and queen. Based on the novel by Katherine Paterson. Copyright Tribune Media Services, Inc.`
	if detail.Description != expectedString {
		t.Errorf(`Expected Description to be [%s], not [%s]`, expectedString, detail.Description)
	}

	// No XML input to test the following:
	//   GuestStars
	//   Choreographers
	//   Hosts

	expectedString = `2008-01-08T11:30:00Z`
	if detail.Time != expectedString {
		t.Errorf(`Expected Time to be %s, not %s`, expectedString, detail.Time)
	}

	expectedInt = 2007
	if detail.MovieYear != expectedInt {
		t.Errorf(`Expected MovieYear to be %d, not %d`, expectedInt, detail.MovieYear)
	}

	expectedInt = 2007
	if detail.MovieYear != expectedInt {
		t.Errorf(`Expected MovieYear to be %d, not %d`, expectedInt, detail.MovieYear)
	}

	expectedGenre = []Genre{"Fantasy", "Movies", "Sci-Fi and Fantasy"}
	for _, g := range expectedGenre {
		if !inGenre(g, detail.SeriesGenres) {
			t.Errorf(`Expected SeriesGenres to contain %s`, g)
		}
	}

	expectedPerson = []Person{
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
	for _, p := range expectedPerson {
		if !inPerson(p, detail.Actors) {
			t.Errorf(`Expected Actors to contain %s`, p)
		}
	}

	expectedPerson = []Person{
		"Csupo|Gabor",
	}
	for _, p := range expectedPerson {
		if !inPerson(p, detail.Directors) {
			t.Errorf(`Expected Directors to contain %s`, p)
		}
	}

	expectedPerson = []Person{
		"Lieberman|Hal",
		"Levine|Lauren",
	}
	for _, p := range expectedPerson {
		if !inPerson(p, detail.Producers) {
			t.Errorf(`Expected Producers to contain %s`, p)
		}
	}

	expectedPerson = []Person{
		"Paterson|David",
		"Stockwell|Jeff",
		"Wade|Kevin",
	}
	for _, p := range expectedPerson {
		if !inPerson(p, detail.Writers) {
			t.Errorf(`Expected Writers to contain %s`, p)
		}
	}
}

func inPerson(findit Person, list []Person) bool {
	for _, p := range list {
		if p == findit {
			return true
		}
	}
	return false
}

func inGenre(findit Genre, list []Genre) bool {
	for _, g := range list {
		if g == findit {
			return true
		}
	}
	return false
}

// Test XML unmarshaling of TiVoContainer items
func tvContainerItem(t *testing.T, item ContainerItem) {
	var (
		expectedString string
		expectedInt    int
		expectedInt64  int64
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

	expectedString = "0x5089FCEA"
	if item.CaptureDate != expectedString {
		t.Errorf(`Expected CaptureDate to be %s, not %s`, expectedString, item.CaptureDate)
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

	expectedString = "No"
	if item.HighDefinition != expectedString {
		t.Errorf(`Expected HighDefinition to be %s, not %s`, expectedString, item.HighDefinition)
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
