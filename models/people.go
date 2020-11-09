package model

import (
	// "fmt"
	// "net/http"
	"swapigo/lib"
	"sync"
)

// Person is a star wars person
type Person struct {
	Name      string   `json:"name"`
	Height    string   `json:"height"`
	Mass      string   `json:"mass"`
	HairColor string   `json:"hair_color"`
	SkinColor string   `json:"skin_color"`
	EyeColor  string   `json:"eye_color"`
	BirthYear string   `json:"birth_year"`
	Gender    string   `json:"gender"`
	Created   string   `json:"created"`
	Edited    string   `json:"edited"`
	URL       string   `json:"url"`
	Homeworld string   `json:"homeworld"`
	Films     []string `json:"films"`
	Species   []string `json:"species"`
	Vehicles  []string `json:"vehicles"`
	Starships []string `json:"starships"`
}

// PeoplePageResponse is the response for a page of people
type PeoplePageResponse struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []Person
}

// GetInitialPerson gets a single person using a wait group
func GetInitialPerson(url string, wg *sync.WaitGroup) (Person, error) {
	var p Person
	return p, lib.GetJSONwg(url, &p, wg)
}

// GetInitialPeople fetches a page of people
func GetInitialPeople(url string, wg *sync.WaitGroup) (PeoplePageResponse, error) {
	var page PeoplePageResponse
	return page, lib.GetJSONwg(url, &page, wg)
}

// GetHomeworld gets a persons homeworld
func (person Person) GetHomeworld() (Planet, error) {
	var p Planet
	return p, lib.GetJSON(person.Homeworld, &p)
}

// GetSpecies gets all the species for a person
func (person Person) GetSpecies() (species []Species, e error) {
	for _, url := range person.Species {
		var s Species
		if e := lib.GetJSON(url, &s); e != nil {
			return species, e
		}
		species = append(species, s)
	}
	return
}

// GetStarships gets all the starships for person
func (person Person) GetStarships() (starships []Starship, e error) {
	for _, url := range person.Starships {
		var s Starship
		if e := lib.GetJSON(url, &s); e != nil {
			return starships, e
		}
		starships = append(starships, s)
	}
	return
}

// GetVehicles gets all the vehicles for a person
func (person Person) GetVehicles() (vehicles []Vehicle, err error) {
	for _, url := range person.Vehicles {
		var v Vehicle
		if err := lib.GetJSON(url, &v); err != nil {
			return vehicles, err
		}
		vehicles = append(vehicles, v)
	}
	return
}

// GetFilms gets all the films on a person instance
func (person Person) GetFilms() (films []Film, err error) {
	for _, url := range person.Films {
		var film Film
		if e := lib.GetJSON(url, &film); e != nil {
			return films, err
		}
		films = append(films, film)
	}
	return
}

// func (person Person) GetStarship() (starships []Starship, err, error) {
// for _, url := range person.
// }
