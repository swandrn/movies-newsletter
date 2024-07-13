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

func (DiscoverPage) Pagination(page int) []int {
	return []int{
		page - 2,
		page - 1,
		page,
		page + 1,
		page + 2,
	}
}

func (dp DiscoverPage) ShowPagination(page int) bool {
	return (page < dp.Discover.TotalPages-2 && page > 2)
}

func (dp DiscoverPage) ShowFirstPage(page int) bool {
	return page > 3
}

func (dp DiscoverPage) ShowLastPage(page int) bool {
	return page < dp.Discover.TotalPages-2
}

func (dp DiscoverPage) LastPages() []int {
	return []int{
		dp.Discover.TotalPages - 4,
		dp.Discover.TotalPages - 3,
		dp.Discover.TotalPages - 2,
		dp.Discover.TotalPages - 1,
		dp.Discover.TotalPages,
	}
}

func (DiscoverPage) PreviousPage(page int) int {
	return page - 1
}

func (DiscoverPage) NextPage(page int) int {
	return page + 1
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

func (GenrePage) Pagination(page int) []int {
	return []int{
		page - 2,
		page - 1,
		page,
		page + 1,
		page + 2,
	}
}

func (dp GenrePage) ShowPagination(page int) bool {
	return (page < dp.Discover.TotalPages-2 && page > 2)
}

func (dp GenrePage) ShowFirstPage(page int) bool {
	return page > 3
}

func (dp GenrePage) ShowLastPage(page int) bool {
	return page < dp.Discover.TotalPages-2
}

func (dp GenrePage) LastPages() []int {
	return []int{
		dp.Discover.TotalPages - 4,
		dp.Discover.TotalPages - 3,
		dp.Discover.TotalPages - 2,
		dp.Discover.TotalPages - 1,
		dp.Discover.TotalPages,
	}
}

func (GenrePage) PreviousPage(page int) int {
	return page - 1
}

func (GenrePage) NextPage(page int) int {
	return page + 1
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
