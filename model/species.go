package model

// Species is a creature in star wars
type Species struct {
	Name            string   `json:"name"`
	Classification  string   `json:"classification"`
	Designation     string   `json:"designation"`
	AverageHeight   string   `json:"average_height"`
	SkinColors      string   `json:"skin_colors"`
	HairColors      string   `json:"hair_colors"`
	EyeColors       string   `json:"eye_colors"`
	AverageLifespan string   `json:"average_lifespan"`
	Homeworld       string   `json:"homeworld"`
	Language        string   `json:"language"`
	People          []string `json:"people"`
	Films           []string `json:"films"`
	Created         string   `json:"created"`
	Edited          string   `json:"edited"`
	URL             string   `json:"url"`
}

// GetSpeciesCard returns a SubCard of specie values
func GetSpeciesCard(s Species) SubCard {
	return SubCard{
		Title:     "Name: " + s.Name,
		SubTitle:  "Classification: " + s.Classification,
		SubTitle2: "Designation: " + s.Designation,
		Body:      "They speak " + s.Language,
		URL:       "",
	}
}
