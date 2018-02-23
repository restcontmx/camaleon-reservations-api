package main

import (
	"log"
	"net/http"
	"os"

	_ "github.com/heroku/x/hmetrics/onload"
	"github.com/restcontmx/camaleon-reservations-api/config"
)

//
// run server
//sets all for the server to run
//
func runServer() {

	graphqlHandler := http.HandlerFunc(config.GraphqlHandlerFunc)

	http.Handle("/graphql", graphqlHandler)

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
