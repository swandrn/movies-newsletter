package tmpl

import (
	"html/template"
	"net/http"
	"strconv"
	"strings"
	"tmdb-consumer/api"
)

type HomePage struct {
	Title   string
	Heading string
	Movies  []api.Movie
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	moviesDiscover, err := api.FetchMoviesByPage(1)
	if err != nil {
		panic(err)
	}

	title := "Home"
	heading := "Movies"
	p := HomePage{
		title,
		heading,
		moviesDiscover.Results,
	}

	t, err := template.ParseFiles("./web/templates/index.html")
	if err != nil {
		panic(err)
	}

	err = t.Execute(w, p)
	if err != nil {
		panic(err)
	}
}

type ContactPage struct {
	Title   string
	Heading string
}

func ContactHandler(w http.ResponseWriter, r *http.Request) {
	title := "Contact"
	heading := "Contact"
	p := ContactPage{title, heading}

	t, err := template.ParseFiles("./web/templates/contact.html")
	if err != nil {
		panic(err)
	}

	err = t.Execute(w, p)
	if err != nil {
		panic(err)
	}
}

type DiscoverPage struct {
	Title    string
	Heading  string
	Discover api.SearchResult
}

func DiscoverHandler(w http.ResponseWriter, r *http.Request) {
	page, err := strconv.Atoi(strings.TrimPrefix(r.URL.Path, "/discover/"))
	if err != nil {
		panic(err)
	}

	moviesDiscover, err := api.FetchMoviesByPage(page)
	if err != nil {
		panic(err)
	}

	title := "Discover"
	heading := "Discover"
	p := DiscoverPage{
		title,
		heading,
		moviesDiscover,
	}

	t, err := template.ParseFiles("./web/templates/discover.html")
	if err != nil {
		panic(err)
	}

	err = t.Execute(w, p)
	if err != nil {
		panic(err)
	}
}

type MoviePage struct {
	Title   string
	Heading string
	Movie   api.MovieDetails
}

func MovieHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(strings.TrimPrefix(r.URL.Path, "/movie/"))
	if err != nil {
		panic(err)
	}

	movie, err := api.FetchMovieById(id)
	if err != nil {
		panic(err)
	}

	title := movie.Title
	heading := movie.Title
	p := MoviePage{
		title,
		heading,
		movie,
	}

	t, err := template.ParseFiles("./web/templates/movie.html")
	if err != nil {
		panic(err)
	}

	err = t.Execute(w, p)
	if err != nil {
		panic(err)
	}
}

type GenrePage struct {
	Title    string
	Heading  string
	Discover api.SearchResult
}

func GenreHandler(w http.ResponseWriter, r *http.Request) {
	gets := strings.Split(strings.TrimPrefix(r.URL.Path, "/genre/"), "/")
	page, err := strconv.Atoi(gets[1])
	if err != nil {
		panic(err)
	}

	genre, err := strconv.Atoi(gets[0])
	if err != nil {
		panic(err)
	}

	moviesGenre, err := api.FetchMoviesByGenre(page, genre)
	if err != nil {
		panic(err)
	}

	title := "Genre"
	heading := "Genre"
	p := DiscoverPage{
		title,
		heading,
		moviesGenre,
	}

	t, err := template.ParseFiles("./web/templates/genre.html")
	if err != nil {
		panic(err)
	}

	err = t.Execute(w, p)
	if err != nil {
		panic(err)
	}
}
