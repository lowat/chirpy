package main

import (
	"log"
	"net/http"
	"sort"
	"strconv"
)

func (cfg *apiConfig) handlerChirpsRetrieve(w http.ResponseWriter, r *http.Request) {
	dbChirps, err := cfg.DB.GetChirps()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't retrieve chirps")
		return
	}

	chirps := []Chirp{}
	for _, dbChirp := range dbChirps {
		chirps = append(chirps, Chirp{
			ID:   dbChirp.ID,
			Body: dbChirp.Body,
		})
	}

	sort.Slice(chirps, func(i, j int) bool {
		return chirps[i].ID < chirps[j].ID
	})

	respondWithJSON(w, http.StatusOK, chirps)
}

func (cfg *apiConfig) handlerChirpsRetrieveByID(w http.ResponseWriter, r *http.Request) {
	i, err := strconv.Atoi(r.PathValue("id"))
	log.Printf("%v", i)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't retrieve chirp with provided id")
		return
	}
	chirp, err := cfg.DB.GetChirpsByID(i)
	if err != nil {
		respondWithError(w, 404, "Couldn't retrieve chirps")
		return
	}
	respondWithJSON(w, http.StatusOK, chirp)
}
