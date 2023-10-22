package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	fileServer := http.FileServer(http.Dir("./static"))

	mux.Handle("/", fileServer)
	mux.HandleFunc("/form", formHandler)
	mux.HandleFunc("/hello", helloHandler)

	log.Printf("Starting server at port 8080\n")
	err := http.ListenAndServe(":8080", mux)
	log.Fatal(err)
}
