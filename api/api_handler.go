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
}

func fetchMoviesByPage(pageNumber int) (SearchResult, error) {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	response, err := http.Get(fmt.Sprintf("https://api.themoviedb.org/3/discover/movie?include_adult=true&include_video=false&language=fr-FR&page=%d&sort_by=popularity.desc&api_key=%s", pageNumber, os.Getenv("TMDB_API_KEY")))
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
