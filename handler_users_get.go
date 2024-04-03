package main

import (
	"encoding/json"
	"net/http"
)

func (cfg *apiConfig) handlerUsersLogin(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	type responseUser struct {
		Email string `json:"email"`
		ID    int    `json:"id"`
	}
	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, 401, "Unauthorized")
		return
	}

	user, err := cfg.DB.Login(params.Email, params.Password)
	if err != nil {
		respondWithError(w, 401, "Unauthorized")
	}

	response := responseUser{
		Email: user.Email,
		ID:    user.ID,
	}

	respondWithJSON(w, http.StatusOK, response)
}
