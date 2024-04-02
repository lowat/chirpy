package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
)

func (cfg *apiConfig) handlerValidateChirp(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Body string `json:"body"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)

	if err != nil {
		log.Printf("Error decoding parameters: %s", err)
		w.WriteHeader(500)
		return
	}

	type returnVals struct {
		Body string `json:"cleaned_body"`
	}

	// chirpIsValid := chirpIsValid(params.Body)
	badWords := map[string]bool{"kerfuffle": true, "sharbert": true, "fornax": true}
	const redacted = "****"
	cleaned_chirp := replaceBadWordsWithRedacted(params.Body, badWords, redacted)
	respBody := returnVals{
		Body: cleaned_chirp,
	}

	dat, err := json.Marshal(respBody)
	if err != nil {
		log.Printf("Error marshalling JSON: %s", err)
		w.WriteHeader(500)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	if chirpIsValid(params.Body) {
		w.WriteHeader(http.StatusOK)
		w.Write(dat)
	} else {
		w.WriteHeader(400)
	}

}

func chirpIsValid(chirp string) bool {
	return len(chirp) <= 140
}

func replaceBadWordsWithRedacted(original string, badWords map[string]bool, redacted string) string {
	orig_words := strings.Split(original, " ")
	res := []string{}
	for _, word := range orig_words {
		if _, isContained := badWords[strings.ToLower(word)]; isContained {
			res = append(res, redacted)
		} else {
			res = append(res, word)
		}
	}
	return strings.Join(res, " ")
}
