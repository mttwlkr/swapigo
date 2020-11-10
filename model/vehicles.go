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
