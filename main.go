package main

import (
	"log"
	"net/http"
)

func message(w http.ResponseWriter, r *http.Request) {
	log.Print(w, "Welcome to my Go web server!")
}

func main() {
	log.Print("Hello, World!")

	mux := http.NewServeMux()

	mux.HandleFunc("/", message)
	log.Print("starting server on :4000")

	// start a local web server
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)

}
