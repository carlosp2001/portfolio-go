package main

import (
	"log"
	"net/http"
	"portfolio/backend/handler"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("./frontend")))
	http.HandleFunc("/contact", handler.ContactHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
