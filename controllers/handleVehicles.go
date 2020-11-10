package controllers

import (
	"fmt"
	"html/template"
	"net/http"
	"swapigo/lib"
	"swapigo/models"
	"sync"
)

// HandleVehicles gets a page of vehicles
func HandleVehicles(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("views/index-page.html"))
	var page model.IndexPageResponse

	var wg sync.WaitGroup

	wg.Add(1)
	queryPageNumber := "1"
	keys := r.URL.Query()["id"]
	if len(keys) > 0 {
		queryPageNumber = keys[0]
	}

	vehiclePage, err := model.GetInitialVehicles(lib.BaseURL+"vehicles/?page="+queryPageNumber, &wg)

	if err != nil {
		fmt.Println(err.Error())
	}

	wg.Wait()

	page.PageTitle = "Vehicles"
	page.NextPage = "/vehicles?id=" + lib.GetIDFromString(vehiclePage.Next)
	page.PreviousPage = "/vehicles?id=" + lib.GetIDFromString(vehiclePage.Previous)

	page.Cards = make([]model.SubCard, 0)
	for _, vehicle := range vehiclePage.Results {
		page.Cards = append(page.Cards, model.SubCard{
			Title:     "Name: " + vehicle.Name,
			SubTitle:  "Manufacturer: " + vehicle.Manufacturer,
			SubTitle2: "Model: " + vehicle.Model,
			Body:      "blah",
			URL:       "/vehicle?id=" + lib.GetIDFromString(vehicle.URL),
		})
	}

	tmpl.Execute(w, page)
}
