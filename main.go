package main

import (
	"fmt"
	"log"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
)

func main() {
	// loading the .env in our current env
	godotenv.Load(".env")
	//creating the port var by setting equal to our port from .env
	portString := os.Getenv("PORT")
	if portString == "" {
		log.Fatal("PORT is not found")
	}

	//setting up the router
	router := chi.NewRouter()

	srv := &http.Server{
		Handler: router,
		Addr:    ":" + portString,
	}

	log.Printf("Port listening to port %v", portString)
	err := srv.ListenAndServe(){
		if err != nil {
			log.Fatal(err)
		}
	}

}
