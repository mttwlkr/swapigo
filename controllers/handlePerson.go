package controllers

import (
	"fmt"
	"html/template"
	// "log"
	"net/http"
	// "net/url"
	"swapigo/lib"
	"swapigo/models"
	"sync"
)

// PersonPageResponse is the response to a request for an individual user
type PersonPageResponse struct {
	PageTitle string
	Person    model.MainCard
	Vehicles  []model.SubCard
	Species   []model.SubCard
	Starships []model.SubCard
}

// HandlePerson handles a single get request for a person and all their associated values
func HandlePerson(w http.ResponseWriter, r *http.Request) {
	// get id
	keys := r.URL.Query()["id"]
	personID := keys[0]

	var wg sync.WaitGroup
	tmpl := template.Must(template.ParseFiles("views/detail-page.html"))
	var page PersonPageResponse

	fmt.Println("we in here")
	wg.Add(1)
	person, pErr := model.GetInitialPerson(lib.BaseURL+"people/"+personID, &wg)
	if pErr != nil {
		fmt.Println(pErr.Error())
	}
	wg.Wait()

	fmt.Println("person", person)
	page.PageTitle = person.Name
	page.Person.Title = "Name: " + person.Name
	page.Person.Body1 = "Born: " + person.BirthYear
	page.Person.Body2 = "Gender: " + person.Gender
	page.Person.Body3 = "Height: " + person.Height + " CM"
	page.Person.Body4 = "Mass: " + person.Mass + " KG"
	page.Person.Body5 = "Eyes: " + person.EyeColor + " eyes"

	homeworld, hErr := person.GetHomeworld()
	if hErr != nil {
		fmt.Println(hErr.Error())
	}
	page.Person.SubTitle = "Homeworld: " + homeworld.Name
	fmt.Println("Homeworld", homeworld)

	// vehicles
	if len(person.Vehicles) > 0 {
		vehicleChannel := make(chan []model.Vehicle)
		go person.GetVehicles(vehicleChannel)
		vehicles := <-vehicleChannel
		page.Vehicles = make([]model.SubCard, 0)
		for _, vehicle := range vehicles {
			page.Vehicles = append(page.Vehicles, model.SubCard{
				Title:     "Name: " + vehicle.Name,
				SubTitle:  "Manufacturer: " + vehicle.Manufacturer,
				SubTitle2: "Model: " + vehicle.Model,
				Body:      "The vehicle hodls " + vehicle.Crew + " crew & " + vehicle.Passengers + " passengers",
				URL:       "",
			})
		}
	}

	if len(person.Species) > 0 {
		speciesChannel := make(chan []model.Species)
		go person.GetSpecies(speciesChannel)
		species := <-speciesChannel
		page.Species = make([]model.SubCard, 0)
		for _, specie := range species {
			page.Species = append(page.Species, model.SubCard{
				Title:     "Name: " + specie.Name,
				SubTitle:  "Classification: " + specie.Classification,
				SubTitle2: "Designation: " + specie.Designation,
				Body:      "They are from " + specie.Homeworld + " and speak " + specie.Language,
				URL:       "",
			})
		}
	}

	// starships
	if len(person.Starships) > 0 {
		starshipChannel := make(chan []model.Starship)
		go person.GetStarships(starshipChannel)
		starships := <-starshipChannel
		page.Starships = make([]model.SubCard, 0)
		for _, ship := range starships {
			fmt.Println("ship.Name", ship.Name)
			page.Starships = append(page.Starships, model.SubCard{
				Title:     "Name: " + ship.Name,
				SubTitle:  "Manufacturer: " + ship.Manufacturer,
				SubTitle2: "Model: " + ship.Model,
				Body:      "The ship hodls " + ship.Crew + " crew & " + ship.Passengers + " passengers",
				URL:       "",
			})
		}
	}

	tmpl.Execute(w, page)
}
