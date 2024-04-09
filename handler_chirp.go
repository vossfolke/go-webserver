package main

import (
	"net/http"
)

func (cfg *apiConfig) handlerChirps(w http.ResponseWriter, r *http.Request) {

	chirps, err := cfg.DB.GetChirps()
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	respondWithJSON(w, http.StatusOK, chirps)
	w.Header().Set("Content-Type", "application/json")
}
