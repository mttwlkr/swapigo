package controllers

import (
	"fmt"
	"html/template"
	// "log"
	"net/http"
	"swapigo/api"
	"swapigo/lib"
	"sync"
)

// MainCard is the main focus of the page
type MainCard struct {
	Title, SubTitle, Body1, Body2, Body3, Body4, Body5 string
}

// SubCard describes cards lower on the page
type SubCard struct {
	Title, SubTitle, Body string
}

// PersonPageResponse is the response to a request for an individual user
type PersonPageResponse struct {
	PageTitle string
	Person    MainCard
	Vehicles  []SubCard
	Species   []SubCard
	Starships []SubCard
}

var wg sync.WaitGroup

// HandlePerson handles a single get request for a person and all their associated values
func HandlePerson(w http.ResponseWriter, r *http.Request) {
	// get id

	tmpl := template.Must(template.ParseFiles("views/detail-page.html"))
	var page PersonPageResponse

	fmt.Println("we in here")
	wg.Add(1)
	person, pErr := api.GetInitialPerson(lib.BaseURL+"people/1/", &wg)
	if pErr != nil {
		fmt.Println(pErr.Error())
	}
	wg.Wait()

	fmt.Println("person", person)
	page.PageTitle = person.Name
	page.Person.Title = person.Name + " - " + person.Gender
	page.Person.Body1 = "Born: " + person.BirthYear
	page.Person.Body2 = person.Gender
	page.Person.Body3 = person.Height + " CM"
	page.Person.Body4 = person.Mass + " KG"
	page.Person.Body5 = person.EyeColor + " eyes & " + person.HairColor + " hair"

	homeworld, hErr := person.GetHomeworld()
	if hErr != nil {
		fmt.Println(hErr.Error())
	}
	page.Person.SubTitle = homeworld.Name
	fmt.Println("Homeworld", homeworld)

	// vehicles
	if len(person.Vehicles) > 0 {
		vehicles, vErr := person.GetVehicles()
		if vErr != nil {
			fmt.Println(vErr.Error())
		}

		page.Vehicles = make([]SubCard, 0)
		for _, vehicle := range vehicles {
			page.Vehicles = append(page.Vehicles, SubCard{
				Title:    vehicle.Name,
				SubTitle: vehicle.Manufacturer + " - " + vehicle.Model,
				Body:     "It hodls " + vehicle.Crew + " crew & " + vehicle.Passengers + " passengers",
			})
		}
	}

	// species
	if len(person.Species) > 0 {
		species, sErr := person.GetSpecies()
		if sErr != nil {
			fmt.Println(sErr.Error())
		}

		page.Species = make([]SubCard, 0)
		for _, specie := range species {
			page.Species = append(page.Species, SubCard{
				Title:    specie.Name,
				SubTitle: specie.Classification + " - " + specie.Designation,
				Body:     "They are from " + specie.Homeworld + " and speak " + specie.Language,
			})
		}
	}

	// starships
	if len(person.Starships) > 0 {
		starships, sErr := person.GetStarships()
		if sErr != nil {
			fmt.Println(sErr.Error())
		}

		page.Starships = make([]SubCard, 0)
		for _, ship := range starships {
			fmt.Println("ship.Name", ship.Name)
			page.Starships = append(page.Starships, SubCard{
				Title:    ship.Name,
				SubTitle: ship.Manufacturer + " - " + ship.Model,
				Body:     "It hodls " + ship.Crew + " crew & " + ship.Passengers + " passengers",
			})
		}
	}
	
	tmpl.Execute(w, page)
}
