//Package utils - code for translation taken from https://cloud.google.com/translate/docs/reference/libraries
package utils

import (
	"context"
	"log"

	"cloud.google.com/go/translate"
	"golang.org/x/text/language"
)

//Translate translate the given text to the preffered language
func Translate(text *string, l string) {
	ctx := context.Background()

	// Creates a client.
	client, err := translate.NewClient(ctx)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	// Sets the target language.
	target, err := language.Parse(l)
	if err != nil {
		log.Fatalf("Failed to parse target language: %v", err)
	}

	// Translates the text into preffered lanugage.
	translations, err := client.Translate(ctx, []string{*text}, target, nil)
	if err != nil {
		log.Fatalf("Failed to translate text: %v", err)
	}

	*text = translations[0].Text
}
