package main

import "net/http"

func (apicfg *apiConfig) handlerCreateUser(w http.ResponseWriter, r *http.Request) {
	respondWithJSON(w, 400, "something went wrong")
}
