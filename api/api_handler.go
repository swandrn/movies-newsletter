package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

type SearchResult struct {
	Page         int     `json:"page"`
	Results      []Movie `json:"results"`
	TotalPages   int     `json:"total_pages"`
	TotalResults int     `json:"total_results"`
}

type Movie struct {
	Adult         bool   `json:"adult"`
	Id            int    `json:"id"`
	OriginalTitle string `json:"original_title"`
	Title         string `json:"title"`
	Overview      string `json:"overview"`
	GenreId       []int  `json:"genre_ids"`
	ReleaseDate   string `json:"release_date"`
	ImagePath     string `json:"poster_path"`
}

func FetchMoviesByPage(pageNumber int) (SearchResult, error) {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	response, err := http.Get(fmt.Sprintf("https://api.themoviedb.org/3/discover/movie?include_adult=true&include_video=false&language=fr-FR&page=%d&sort_by=primary_release_date.desc&api_key=%s", pageNumber, os.Getenv("TMDB_API_KEY")))
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return SearchResult{}, errors.New(response.Status)
	}

	data, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	var searchResult SearchResult
	err = json.Unmarshal(data, &searchResult)
	if err != nil {
		panic(err)
	}

	return searchResult, nil
}

func FetchMoviesByGenre(pageNumber int, genreId int) (SearchResult, error) {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	response, err := http.Get(fmt.Sprintf("https://api.themoviedb.org/3/discover/movie?include_adult=true&include_video=false&language=fr-FR&page=%d&sort_by=primary_release_date.desc&with_genres=%d&api_key=%s", pageNumber, genreId, os.Getenv("TMDB_API_KEY")))
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return SearchResult{}, errors.New(response.Status)
	}

	data, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	var searchResult SearchResult
	err = json.Unmarshal(data, &searchResult)
	if err != nil {
		panic(err)
	}

	return searchResult, nil
}

type MovieDetails struct {
	Adult         bool    `json:"adult"`
	Id            int     `json:"id"`
	OriginalTitle string  `json:"original_title"`
	Title         string  `json:"title"`
	Overview      string  `json:"overview"`
	Genres        []Genre `json:"genres"`
	ReleaseDate   string  `json:"release_date"`
	ImagePath     string  `json:"poster_path"`
}

type Genre struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func FetchMovieById(movieId int) (MovieDetails, error) {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	response, err := http.Get(fmt.Sprintf("https://api.themoviedb.org/3/movie/%d?language=fr-FR&api_key=%s", movieId, os.Getenv("TMDB_API_KEY")))
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return MovieDetails{}, errors.New(response.Status)
	}

	data, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	var movie MovieDetails
	err = json.Unmarshal(data, &movie)
	if err != nil {
		panic(err)
	}

	return movie, nil
}
