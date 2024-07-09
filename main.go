package main

import (
	"log"
	"net/http"
	tmpl "tmdb-consumer/serving_templates"
)

func main() {
	http.HandleFunc("/", tmpl.HomeHandler)
	http.HandleFunc("/home", tmpl.HomeHandler)
	http.HandleFunc("/contact", tmpl.ContactHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
