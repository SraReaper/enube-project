package main

import (
	"SraReaper/enube-me-api/api/src/router"
	"log"
	"net/http"
)

// main initializes the API server and starts listening on port 5050.
func main() {

	conn := router.Route()
	log.Fatal(http.ListenAndServe(":5050", conn))
}
