package model

// Starship is a star wars starship
type Starship struct {
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
	HyperdriveRating     string   `json:"hyperdrive_rating"`
	MGLT                 string   `json:"MGLT"`
	StarshipClass        string   `json:"starship_class"`
	Pilots               []string `json:"pilots"`
	Films                []string `json:"films"`
	Created              string   `json:"created"`
	Edited               string   `json:"edited"`
	URL                  string   `json:"url"`
}

// GetStarshipCard returns a SubCard of Starship attributes
func GetStarshipCard(s Starship) SubCard {
	return SubCard{
		Title:     "Name: " + s.Name,
		SubTitle:  "Manufacturer: " + s.Manufacturer,
		SubTitle2: "Model: " + s.Model,
		Body:      "The ship hodls " + s.Crew + " crew & " + s.Passengers + " passengers",
		URL:       "",
	}
}
