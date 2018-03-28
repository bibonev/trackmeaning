package utils

type metadata struct {
	Title       string `json:"title"`
	Coverart56  string `json:"coverart56"`
	Coverart112 string `json:"coverart112"`
	Coverart500 string `json:"coverart500"`
	Artist      string `json:"artist"`
	ReleaseDate string `json:"releasedate"`
	AlbumTitle  string `json:"albumtitle"`
	ArtistArtHq string `json:"artistarthq"`
}

type match struct {
	Key      string   `json:"key"`
	TrackID  string   `json:"trackid"`
	Metadata metadata `json:"metadata"`
	Type     string   `json:"type"`
}

//RecognitionAudioAPIResponse - The struct that describes the response from the Recognition Audio API
type RecognitionAudioAPIResponse struct {
	Matches []match `json:"matches"`
}

//AnalyzeResponse - The struct that describes the response from the Analyze handler
type AnalyzeResponse struct {
	NounPrefix      string
	Noun            string
	KeyLyricsPrefix string
	KeyLyrics       []string
	Error           []string
}

type section struct {
	Type string   `json:"type"`
	Text []string `json:"text"`
}

//LookupTrackAPIResponse - The struct that describes the response from the Lookup Track API
type LookupTrackAPIResponse struct {
	Sections []section `json:"sections"`
}
