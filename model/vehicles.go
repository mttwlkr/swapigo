package model

import (
	"swapigo/lib"
	"sync"
)

// Vehicle is a star wars vehicle
type Vehicle struct {
	Name                 string   `json:"name"`
	Model                string   `json:"model"`
	Manufacturer         string   `json:"manufacturer"`
	CostInCredits        string   `json:"cost_in_credits"`
	Length               string   `json:"length"`
	MaxAtmospheringSpeed string   `json:"max_atmosphering_speed"`
	Crew                 string   `json:"crew"`
	Passengers           string   `json:"passengers"`
	CargoCapacity        string   `json:"cargo_capacity"`
	Consumables          string   `json:"consumables"`
	VehicleClass         string   `json:"vechicle_class"`
	Pilots               []string `json:"pilots"`
	Films                []string `json:"films"`
	Created              string   `json:"created"`
	Edited               string   `json:"edited"`
	URL                  string   `json:"url"`
}

// VehiclePageResponse is the response for a page of vehicles
type VehiclePageResponse struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []Vehicle
}

// GetInitialVehicles fetches a page of vehicles using a wait group
func GetInitialVehicles(url string, wg *sync.WaitGroup) (VehiclePageResponse, error) {
	var page VehiclePageResponse
	return page, lib.GetJSONwg(url, &page, wg)
}

// GetInitialVehicle gets a vehicle using a wait group
func GetInitialVehicle(url string, wg *sync.WaitGroup) (Vehicle, error) {
	var v Vehicle
	return v, lib.GetJSONwg(url, &v, wg)
}

// GetVehicleCard returns a SubCard of vehicle attributes
func GetVehicleCard(v Vehicle) SubCard {
	return SubCard{
		Title:     "Name: " + v.Name,
		SubTitle:  "Manufacturer: " + v.Manufacturer,
		SubTitle2: "Model: " + v.Model,
		Body:      "The vehicle hodls " + v.Crew + " crew & " + v.Passengers + " passengers",
		URL:       "/vehicle?id=" + lib.GetIDFromString(v.URL),
	}
}

// GetFilms gets all films a vehicle is in
func (v Vehicle) GetFilms(filmChannel chan []Film) {
	var filmArray []Film
	for _, url := range v.Films {
		var f Film
		lib.GetJSON(url, &f)
		filmArray = append(filmArray, f)
		filmChannel <- filmArray
	}
}

// GetPilots gets all pilots of a vehicle
func (v Vehicle) GetPilots(pilotChannel chan []Person) {
	var pilotArray []Person
	for _, url := range v.Pilots {
		var p Person
		lib.GetJSON(url, &p)
		pilotArray = append(pilotArray, p)
		pilotChannel <- pilotArray
	}
}
