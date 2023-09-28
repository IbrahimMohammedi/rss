package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	dat, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Failed to Marshal JSON respone: %v", payload)
		w.WriteHeader(500)
		return
	}
	w.Header().Add("Conetent-Type", "application/json")
	w.WriteHeader(200)
	w.Write(dat)
}
