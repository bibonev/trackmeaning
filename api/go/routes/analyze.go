package routes

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	// Imports the Google Cloud Natural Language API client package.
	language "cloud.google.com/go/language/apiv1"
	languagepb "google.golang.org/genproto/googleapis/cloud/language/v1"

	"github.com/gin-gonic/gin"
	"gitlab.com/boyan.bonev/intern-hackday/api/go/utils"
)

func getLyrics(trackID string) string {
	// Create the request
	req, err := http.NewRequest("GET", utils.LookupTrackAPIURL+trackID, nil)
	if err != nil {
		log.Fatalln("There is an error with creting a request to the Lookup Track API:", err)
	}

	// Set Headers
	req.Header.Set("Content-Type", "application/json")

	// Execute the request
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatalln("There is an error with making a request to the Lookup Track API:", err)
	}

	// Convert the JSON to usable object
	lookupTrackResp := new(utils.LookupTrackAPIResponse)
	json.NewDecoder(resp.Body).Decode(lookupTrackResp)

	defer resp.Body.Close()

	var stringArray []string
	for _, section := range lookupTrackResp.Sections {
		if section.Type == "LYRICS" {
			stringArray = section.Text
		}
	}

	// For later on it is better to shift all words that end with ' into their formal form ing
	var textSlice []string
	for _, sentance := range stringArray {
		var wordSlice []string
		for _, word := range strings.Split(sentance, " ") {
			if strings.HasSuffix(word, "'") {
				word = strings.Replace(word, "'", "g", 1)
			}
			wordSlice = append(wordSlice, word)
		}
		sentance = strings.Join(wordSlice, " ")
		textSlice = append(textSlice, sentance)
	}
	return strings.Join(textSlice, ". ")
}

func analyzeLyrics(lyrics string, lan string) utils.AnalyzeResponse {
	/*
		Steps:
		1. Create a client for the Google APIs
		2. Analyze the salience of the entities in the lyrics
		3. Get the most common sentiment (the one with highest score)
		4. Analyze the syntax of the lyrics
		5. Get the key phrases - they must start with a VERN and end with a specific combination of NOUN/ADJ
		6. Translate to the desired language and return
	*/
	var phrases []string
	var noun string
	var errorResponse []string

	// Creates a client.
	ctx := context.Background()
	client, err := language.NewClient(ctx)
	if err != nil {
		log.Fatalf("Failed to create client for the GOODLE APIs: %v", err)
		errorResponse = append(errorResponse, "GOOGLE API FAILED")
	}

	// Analyze the salience
	responseEntities, errEntities := client.AnalyzeEntities(ctx, &languagepb.AnalyzeEntitiesRequest{
		Document: &languagepb.Document{
			Source: &languagepb.Document_Content{
				Content: lyrics,
			},
			Type: languagepb.Document_PLAIN_TEXT,
		},
		EncodingType: languagepb.EncodingType_UTF8})

	if errEntities != nil {
		log.Fatalf("Failed to analyze lyrics: %v", errEntities)
		errorResponse = append(errorResponse, "GOOGLE API FAILED")
	}

	// Get the most common sentiment
	if len(responseEntities.Entities) > 0 {
		var maxSalience float32
		for _, entity := range responseEntities.Entities {
			if entity.Salience > maxSalience {
				maxSalience = entity.Salience
				noun = entity.Name
			}
		}
	} else {
		log.Fatalf("Failed to analyze lyrics: No words in the lyrics")
		errorResponse = append(errorResponse, "NO WORDS IN LYRICS")
	}

	// Analyze the syntax
	response, errSyntax := client.AnalyzeSyntax(ctx, &languagepb.AnalyzeSyntaxRequest{
		Document: &languagepb.Document{
			Source: &languagepb.Document_Content{
				Content: lyrics,
			},
			Type: languagepb.Document_PLAIN_TEXT,
		},
		EncodingType: languagepb.EncodingType_UTF8})

	if errSyntax != nil {
		log.Fatalf("Failed to analyze lyrics: %v", errSyntax)
		errorResponse = append(errorResponse, "GOOGLE API FAILED")
	}

	// Get the key phrases
	verbFound := false
	var sentance string
	if len(response.Tokens) > 0 {
		for _, token := range response.Tokens {
			if !verbFound {
				if token.PartOfSpeech.Tag.String() == "VERB" {
					sentance += token.Text.Content + " "
					verbFound = true
				}
			} else {
				if token.PartOfSpeech.Tag.String() == "NOUN" ||
					token.PartOfSpeech.Tag.String() == "ADJ" {
					sentance += token.Text.Content
					if len(phrases) < utils.MaxPhrases {
						phrases = append(phrases, sentance)
					}
					sentance = ""
					verbFound = false
				} else {
					sentance += token.Text.Content + " "
				}
			}
		}
	} else {
		log.Fatalf("Failed to analyze lyrics: No tokens in the lyrics")
		errorResponse = append(errorResponse, "NO TOKENS IN LYRICS")
	}

	// Translate and return
	var translatedPhrases []string
	for _, phrase := range phrases {
		utils.Translate(&phrase, lan)
		translatedPhrases = append(translatedPhrases, phrase)
	}

	nounWithPrefix := "This song is for the " + noun
	keyLyricsPrefix := "The key phrases in the song are"

	utils.Translate(&nounWithPrefix, lan)
	utils.Translate(&keyLyricsPrefix, lan)
	return utils.AnalyzeResponse{
		Noun:            nounWithPrefix,
		KeyLyricsPrefix: keyLyricsPrefix,
		KeyLyrics:       translatedPhrases,
		Error:           errorResponse}
}

//AnalyzeHandler - The handler to analyze the song and produce the summary
//@param {* gin.Context} c - pointer to the context struct where the passed vars are
//@return {AnalyzeResponse} - the data coming out from analyzing the song
func AnalyzeHandler(c *gin.Context) {
	/*
		Steps:
		1. Get the file and the provided language
		2. Create the Recogntion api request
		3. Execute the Recognition api request
		4. Convert to struct from sting (JSON)
		5. Get the lyrics of the song
		6. Translate them to English
		7. Analyze the lyrics using Google API
		8. Format and return the response
	*/

	// Set development headers
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")

	// Get the file
	file, handler, err := c.Request.FormFile("file")

	if err != nil {
		fmt.Println("File error", err)
		c.JSON(400, gin.H{"Error": "File is corruptted or missing"})
	}

	// Create the file locally and pass it to the recognition api
	f, err := os.OpenFile(handler.Filename, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		fmt.Println("File error", err)
		c.JSON(400, gin.H{"Error": "File cannot be saved"})
	}

	io.Copy(f, file)
	f.Close()
	absPathLocalFile, _ := filepath.Abs("../go/" + handler.Filename)
	localFile, _ := os.Open(absPathLocalFile)

	// Get the prefferred language
	languageParam := c.Query("language")
	if len(languageParam) == 0 {
		c.JSON(400, gin.H{"Error": "Missing parameter: 'language'"})
	}

	// Create the request
	req, err := http.NewRequest("POST", utils.RecognitionAudioAPIURL, localFile)
	if err != nil {
		log.Fatalln("There is an error with creting a request to the Recognise Audio API:", err)
	}

	// Set Headers
	req.Header.Set("X-Shazam-Api-Key", utils.GetShazamKey())
	req.Header.Set("Content-Type", "application/octet-stream")

	// Execute the request
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatalln("There is an error with making a request to the Recognise Audio API:", err)
	}

	// Convert the JSON to usable object
	matches := new(utils.RecognitionAudioAPIResponse)
	json.NewDecoder(resp.Body).Decode(matches)

	defer resp.Body.Close()

	var lyrics string
	if len(matches.Matches) > 0 {
		lyrics = getLyrics(matches.Matches[0].TrackID)
	}

	utils.Translate(&lyrics, "en")

	c.JSON(200, analyzeLyrics(lyrics, languageParam))
}
