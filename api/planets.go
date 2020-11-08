package api

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

// GetPlanet from url
func GetPlanet(url string) (planet Planet, err error) {
	return planet, lib.GetJSON(url, planet)
}

// GetInitialPlanet gets a single planet with a waitgroup
func GetInitialPlanet(url string, wg *sync.WaitGroup) (Planet, error) {
	var p Planet
	return p, lib.GetJSONwg(url, &p, wg)
}

// func (p Planet) GetResidents() ([]Person, error) {
// var people []Person
//
// }
