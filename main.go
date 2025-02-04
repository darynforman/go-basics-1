package main

import (
	"log"
	"net/http"
)

// handler function
func message(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("welcome to my go webserver"))

}

// main
func main() {
	log.Print("Hello, World!")

	//router
	mux := http.NewServeMux()

	mux.HandleFunc("/", message)
	log.Print("starting server on :4000")

	// start a local web server
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)

}
