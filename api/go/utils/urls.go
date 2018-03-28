package utils

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
)

// RecognitionAudioAPIURL - the url for calling the recognition api
var RecognitionAudioAPIURL = "https://beta-amp.shazam.com/partner/recognise"

// LookupTrackAPIURL - the url for calling the lookup track api
var LookupTrackAPIURL = "https://beta-amp.shazam.com/discovery/v5/en/GB/iphone/-/track/"

// GetShazamKey - get the shazam key securely from a file
func GetShazamKey() string {

	absPath, _ := filepath.Abs("../go/shazam-credentials.txt")

	key, err := ioutil.ReadFile(absPath) // just pass the file name
	if err != nil {
		fmt.Print(err)
		return ""
	}

	return string(key)
}
