package main

import (
	"log"
	"net/http"
)

func main() {
	// Creating a new ServeMux and registering the file server
	mux := http.NewServeMux()
	fileServer := http.FileServer(http.Dir("./static"))

	// Registering routes and handlers
	mux.Handle("/", fileServer)
	mux.HandleFunc("/form", formHandler)
	mux.HandleFunc("/hello", helloHandler)

	log.Printf("Starting server at port 8080\n")

	// Removing if statement and using log.Fatal instead
	err := http.ListenAndServe(":8080", mux)
	log.Fatal(err)
}
