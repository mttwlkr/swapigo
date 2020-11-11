package model

import (
	"swapigo/lib"
	"sync"
)

// Planet is a struct of a Swapi planet
type Planet struct {
	Name           string   `json:"name"`
	RotationPeriod string   `json:"rotation_period"`
	OrbitalPeriod  string   `json:"orbital_period"`
	Diameter       string   `json:"diameter"`
	Climate        string   `json:"climate"`
	Gravity        string   `json:"gravity"`
	Terrain        string   `json:"terrain"`
	SurfaceWater   string   `json:"surface_water"`
	Population     string   `json:"population"`
	Residents      []string `json:"residents"`
	Films          []string `json:"films"`
	Created        string   `json:"created"`
	Edited         string   `json:"edited"`
	URL            string   `json:"url"`
}

// PlanetPageResponse is the response for a page of planets
type PlanetPageResponse struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []Planet
}

// GetPlanet from url
func GetPlanet(url string) (planet Planet, err error) {
	return planet, lib.GetJSON(url, planet)
}

// GetInitialPlanet gets a single planet with a waitgroup
func GetInitialPlanet(url string, wg *sync.WaitGroup) (Planet, error) {
	var p Planet
	return p, lib.GetJSONwg(url, &p, wg)
}

// GetInitialPlanets gets a page of planets
func GetInitialPlanets(url string, wg *sync.WaitGroup) (PlanetPageResponse, error) {
	var p PlanetPageResponse
	return p, lib.GetJSONwg(url, &p, wg)
}

// GetFilms gets all the films for a planet
func (p Planet) GetFilms(filmChannel chan []Film) {
	var filmArray []Film
	for _, url := range p.Films {
		var f Film
		lib.GetJSON(url, &f)
		filmArray = append(filmArray, f)
		filmChannel <- filmArray
	}
}

// GetResidents gets all the residents for a planet
func (p Planet) GetResidents(residentChannel chan []Person) {
	var residentArray []Person
	for _, url := range p.Residents {
		var resident Person
		lib.GetJSON(url, &resident)
		residentArray = append(residentArray, resident)
		residentChannel <- residentArray
	}
}
