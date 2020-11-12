package model

import (
	"swapigo/lib"
	"sync"
)

// Film is a struct of a Star Wars movie
type Film struct {
	Title        string   `json:"title"`
	EpisodeID    int64    `json:"episode_id"`
	OpeningCrawl string   `json:"opening_crawl"`
	Director     string   `json:"director"`
	Producer     string   `json:"producer"`
	Characters   []string `json:"characters"`
	Planets      []string `json:"planets"`
	Starships    []string `json:"starships"`
	Vehicles     []string `json:"vehicles"`
	Species      []string `json:"species"`
	Created      string   `json:"created"`
	Edited       string   `json:"edited"`
	URL          string   `json:"url"`
}

// FilmPageResponse is a page of films
type FilmPageResponse struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []Film
}

// GetInitialFilms gets a page of films
func GetInitialFilms(url string, wg *sync.WaitGroup) (FilmPageResponse, error) {
	var f FilmPageResponse
	return f, lib.GetJSONwg(url, &f, wg)
}
