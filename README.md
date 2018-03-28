# Hackday Project

## Idea

`TrackMeaning` is an application that takes as an input a track (WAV file) and a prefferred language (english, french, german, etc) and it gives you back the message of the song (what is the artist singing about) in the language you specified. 

### Steps
1. Get the id of the track of the specified track using `Recognise Audio API`
2. Get the lyrics of the song using `Lookup Track API`
3. Translate the lyrics in the preffered language using translation API
4. Determine the topic of the text using NLP techniques
5. Present the topic to the user

## Technologies

### Backend

* Programming language - `Go`
* Framework for creating the REST API - `Gin`
* External APIs used - `Recognition Audio API, Lookup Track API, Google NLP API, Google Translate API`

### Frontend

* Programming language - `JavaScript`
* Library - `React`, `Redux`
* Other technologies for building and compiling JS - `Webpack`, `Babel`

## Improvements

 * Most of the http requests can be turned into `Goroutines` so that they run concurrently
 * Improve the NLP techniques used so the result is better
 * Add more languages (currently only four are available - English, French, Italian, Bulgarian)

## Build
To run the application:

1. Get SHAZAM API (.txt) and GOOGLE API credentials (.json) that you need to place inside `/api/go` as wel
2. Run from the main directory `docker build --rm -f api/react/Dockerfile -t boyan.bonev:intern-hackday-react api/react`
3. Run from the main directory `docker build --rm -f api/go/Dockerfile -t boyan.bonev:intern-hackday-go api/go`
4. Run `docker run -d -p 5000:5000 boyan.bonev:intern-hackday-react`
5. Run `docker run -d -p 8080:8080 boyan.bonev:intern-hackday-go`

Go to `localhost:5000` and you will see the app running.
