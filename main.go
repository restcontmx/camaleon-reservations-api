package main

import (
	"log"
	"net/http"
	"os"

	"github.com/rs/cors"

	_ "github.com/heroku/x/hmetrics/onload"
	"github.com/restcontmx/camaleon-reservations-api/config"
)

//
// run server
//sets all for the server to run
//
func runServer() {

	graphqlHandler := http.HandlerFunc(config.GraphqlHandlerFunc)

	handler := cors.Default().Handler(graphqlHandler)
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:4200",
			"http://localhost:4123",
			"https://camaleon-reservations.herokuapp.com",
			"https://camaleon-reservations-dev.herokuapp.com"},
		AllowCredentials: true,
		AllowedMethods:   []string{"POST"},
		AllowedHeaders:   []string{"authorization", "content-type"},
	})
	handler = c.Handler(handler)
	http.Handle("/graphql", handler)

	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	log.Println("Running on port " + port)

	http.ListenAndServe(":"+port, nil)

}

//
// Main function
// Runs all the shit like server and config handler for graphql schema
//
func main() {
	config.InitRepositories()
	runServer()
}
