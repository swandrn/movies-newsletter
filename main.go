package main

import (
	"fmt"
	"log"
	"net/http"
	tmpl "tmdb-consumer/serving_templates"
)

func main() {
	http.HandleFunc("/", tmpl.HomeHandler)
	http.HandleFunc("/home", tmpl.HomeHandler)
	http.HandleFunc("/contact", tmpl.ContactHandler)
	http.HandleFunc("/discover/", tmpl.DiscoverHandler)
	http.HandleFunc("/movie/", tmpl.MovieHandler)
	http.HandleFunc("/genre/", tmpl.GenreHandler)

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	fmt.Printf("Server started on port 8080\n")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
