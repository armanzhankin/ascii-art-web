package main

import (
	"asciiartweb/Routes"
	"log"
	"net/http"
)

func main() {
	// connection
	mux := http.NewServeMux()
	fs := http.FileServer(http.Dir("static"))
	mux.Handle("/static/", http.StripPrefix("/static/", fs))
	mux.HandleFunc("/", Routes.Home)
	mux.HandleFunc("/ascii-art", Routes.Ascii)
	// Server loading
	log.Println("launching web-server http://127.0.0.1:8080")
	err := http.ListenAndServe(":8080", mux)
	log.Fatal(err)
}
