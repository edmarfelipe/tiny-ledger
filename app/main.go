package main

import (
	"log"
	"net/http"
)

func main() {
	server := NewServer()
	log.Println("Starting server on port 3000")
	log.Fatal(http.ListenAndServe(":3000", server))
}
