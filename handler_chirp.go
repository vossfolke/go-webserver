package main

import (
	"net/http"
	"strconv"
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

func (cfg *apiConfig) handlerChirpId(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("chirpID"))
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	chirp, err := cfg.DB.GetChirpById(id)
	if err != nil {
		respondWithError(w, http.StatusNotFound, err.Error())
		return
	}
	respondWithJSON(w, http.StatusOK, chirp)
}
