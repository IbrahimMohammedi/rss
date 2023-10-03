package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"
	"rss/internal/database"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

// open connection to data base
type apiConfig struct {
	DB *database.Queries
}

func main() {
	// loading the .env in our current env
	godotenv.Load(".env")
	//creating the port var by setting equal to our port from .env
	portString := os.Getenv("PORT")
	if portString == "" {
		log.Fatal("PORT is not found")
	}
	// import the database connection
	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		log.Fatal("DB_URL is not found")
	}

	conn, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal("Cant connect to database:", err)
	}
	//convert from sql.db to DB.queris
	dbQuerries := database.New(conn)

	//create a new api cfg and pass querries in it
	apiCfg := apiConfig{
		DB: dbQuerries,
	}

	//setting up the router
	router := chi.NewRouter()
	// cors
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))
	//v1 router to use the handlerReadiness handler
	v1Router := chi.NewRouter()
	//Router
	v1Router.Get("/healthz", handlerReadiness)
	v1Router.Get("/err", handlerErr)
	//User
	v1Router.Post("/users", apiCfg.handlerUsersCreate)
	v1Router.Get("/users", apiCfg.middleWareAuth(apiCfg.GetUser))
	//Feed
	v1Router.Post("/feeds", apiCfg.middleWareAuth(apiCfg.handlerFeedsCreate))
<<<<<<< HEAD
	v1Router.Get("/feeds", apiCfg.handlerFeedsGet)
	//Feed Follows
	v1Router.Post("/feed_Follows", apiCfg.middleWareAuth(apiCfg.handlerCreateFeedFollows))
=======
>>>>>>> parent of 76b53b3 (hooked the get feeds end point)
	//mounting v1 path
	router.Mount("/v1", v1Router)

	srv := &http.Server{
		Handler: router,
		Addr:    ":" + portString,
	}

	log.Printf("Server listening to port %v", portString)
	log.Fatal(srv.ListenAndServe())

}
