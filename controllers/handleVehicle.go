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

// HandleVehicle handles a single get request for a vehicle and all their associated values
func HandleVehicle(w http.ResponseWriter, r *http.Request) {
	keys := r.URL.Query()["id"]
	vehicleID := keys[0]

	var wg sync.WaitGroup
	tmpl := template.Must(template.ParseFiles("views/detail-page.html"))
	var page model.DetailPageResponse

	wg.Add(1)
	vehicle, pErr := model.GetInitialVehicle(lib.BaseURL+"vehicles/"+vehicleID+"/", &wg)
	if pErr != nil {
		fmt.Println(pErr.Error())
	}
	wg.Wait()

	fmt.Println("vehicle", vehicle)

	page.PageTitle = "Vehicle"
	page.MainCard.Title = "Name: " + vehicle.Name
	page.MainCard.SubTitle = "Manufacturer: " + vehicle.Manufacturer + " - " + vehicle.Model
	page.MainCard.Body1 = "Cost: " + vehicle.CostInCredits + " Credits"
	page.MainCard.Body2 = "Length: " + vehicle.Length + " units"
	page.MainCard.Body3 = "Crew: " + vehicle.Crew + " people"
	page.MainCard.Body4 = "Passengers: " + vehicle.Passengers + " people"
	page.MainCard.Body5 = "Cargo Capacity: " + vehicle.CargoCapacity

	if len(vehicle.Films) > 0 {
		filmChannel := make(chan []model.Film)
		go vehicle.GetFilms(filmChannel)
		films := <-filmChannel
		page.Cards1Title = "Films"
		page.Cards1 = make([]model.SubCard, 0)
		for _, film := range films {
			page.Cards1 = append(page.Cards1, model.SubCard{
				Title:     "Title: " + film.Title,
				SubTitle:  "Director: " + film.Director,
				SubTitle2: "Producer: " + film.Producer,
				Body:      "Created: " + film.Created,
				URL:       "",
			})
		}
	}

	if len(vehicle.Pilots) > 0 {
		pilotChannel := make(chan []model.Person)
		go vehicle.GetPilots(pilotChannel)
		pilots := <-pilotChannel
		page.Cards2Title = "Pilots"
		page.Cards2 = make([]model.SubCard, 0)
		for _, pilot := range pilots {
			page.Cards2 = append(page.Cards2, model.GetPersonCard(pilot))
		}
	}

	fmt.Println("page: ", page)

	tmpl.Execute(w, page)
}
