package controllers

import (
	"fmt"
	"html/template"
	// "log"
	"net/http"
	// "net/url"
	"swapigo/lib"
	"swapigo/model"
	"sync"
)

// HandlePerson handles a single get request for a person and all their associated values
func HandlePerson(w http.ResponseWriter, r *http.Request) {
	keys := r.URL.Query()["id"]
	personID := keys[0]

	tmpl := template.Must(template.ParseFiles("views/detail-page.html"))
	var page model.DetailPageResponse

	var wg sync.WaitGroup
	wg.Add(1)
	person, pErr := model.GetInitialPerson(lib.BaseURL+"people/"+personID, &wg)
	if pErr != nil {
		fmt.Println(pErr.Error())
	}
	wg.Wait()

	fmt.Println("person", person)
	page.PageTitle = "Person"
	page.MainCard.Title = "Name: " + person.Name
	page.MainCard.Body1 = "Born: " + person.BirthYear
	page.MainCard.Body2 = "Gender: " + person.Gender
	page.MainCard.Body3 = "Height: " + person.Height + " CM"
	page.MainCard.Body4 = "Mass: " + person.Mass + " KG"
	page.MainCard.Body5 = "Eyes: " + person.EyeColor + " eyes"

	homeworld, hErr := person.GetHomeworld()
	if hErr != nil {
		fmt.Println(hErr.Error())
	}
	page.MainCard.SubTitle = "Homeworld: " + homeworld.Name
	fmt.Println("Homeworld", homeworld)

	// vehicles
	if len(person.Vehicles) > 0 {
		vehicleChannel := make(chan []model.Vehicle)
		go person.GetVehicles(vehicleChannel)
		vehicles := <-vehicleChannel
		page.Cards1Title = "Vehicles"
		page.Cards1 = make([]model.SubCard, 0)
		for _, vehicle := range vehicles {
			page.Cards1 = append(page.Cards1, model.SubCard{
				Title:     "Name: " + vehicle.Name,
				SubTitle:  "Manufacturer: " + vehicle.Manufacturer,
				SubTitle2: "Model: " + vehicle.Model,
				Body:      "The vehicle hodls " + vehicle.Crew + " crew & " + vehicle.Passengers + " passengers",
				URL:       "/vehicle?id=" + lib.GetIDFromString(vehicle.URL),
			})
		}
	}

	if len(person.Species) > 0 {
		speciesChannel := make(chan []model.Species)
		go person.GetSpecies(speciesChannel)
		species := <-speciesChannel
		page.Cards2Title = "Species"
		page.Cards2 = make([]model.SubCard, 0)
		for _, specie := range species {
			page.Cards2 = append(page.Cards2, model.GetSpeciesCard(specie))
		}
	}

	// starships
	if len(person.Starships) > 0 {
		starshipChannel := make(chan []model.Starship)
		go person.GetStarships(starshipChannel)
		starships := <-starshipChannel
		page.Cards3Title = "Starships"
		page.Cards3 = make([]model.SubCard, 0)
		for _, ship := range starships {
			page.Cards3 = append(page.Cards3, model.GetStarshipCard(ship))
		}
	}

	tmpl.Execute(w, page)
}
