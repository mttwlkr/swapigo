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

// GetInitialFilm gets a film using a wait group
func GetInitialFilm(url string, wg *sync.WaitGroup) (Film, error) {
	var f Film
	return f, lib.GetJSONwg(url, &f, wg)
}

// GetFilmCard returns a SubCard of Film attributes
func GetFilmCard(f Film) SubCard {
	return SubCard{
		Title:     "Title: " + f.Title,
		SubTitle:  "Director: " + f.Director,
		SubTitle2: "Producer: " + f.Producer,
		Body:      f.OpeningCrawl,
		URL:       "/film?id=" + lib.GetIDFromString(f.URL),
	}
}

// GetCharacters gets all the characters on a film
func (f Film) GetCharacters(characterChannel chan []Person) {
	var characterArray []Person
	for _, url := range f.Characters {
		var p Person
		lib.GetJSON(url, &p)
		characterArray = append(characterArray, p)
		characterChannel <- characterArray
	}
}

// GetPlanets gets all the planets for a film
func (f Film) GetPlanets(planetChannel chan []Planet) {
	var planetArray []Planet
	for _, url := range f.Planets {
		var p Planet
		lib.GetJSON(url, &p)
		planetArray = append(planetArray, p)
		planetChannel <- planetArray
	}
}

// GetStarships gets all the starships for a film
func (f Film) GetStarships(shipChannel chan []Starship) {
	var shipArray []Starship
	for _, url := range f.Starships {
		var s Starship
		lib.GetJSON(url, &s)
		shipArray = append(shipArray, s)
		shipChannel <- shipArray
	}
}

// GetVehicles gets all the vehicles for a film
func (f Film) GetVehicles(vehicleChannel chan []Vehicle) {
	var vehicleArray []Vehicle
	for _, url := range f.Vehicles {
		var v Vehicle
		lib.GetJSON(url, &v)
		vehicleArray = append(vehicleArray, v)
		vehicleChannel <- vehicleArray
	}
}

// GetSpecies gets all the speices for a film
func (f Film) GetSpecies(speciesChannel chan []Species) {
	var speciesArray []Species
	for _, url := range f.Species {
		var s Species
		lib.GetJSON(url, &s)
		speciesArray = append(speciesArray, s)
		speciesChannel <- speciesArray
	}
}
