package controllers

import (
	"fmt"
	"mode"
)

// HandleVehicle handles a single get request for a vehicle and all their associated values
func HandleVehicle(w http.ResponseWriter, r *http.Request) {
	keys := r.URL.Query()["id"]
	vehicleID := keys[0]

	var wg sync.WaitGroup
	tmpl := template.Must(template.ParseFiles("views/detail-page.html"))
	var page model.DetailPageResponse

	wg.Add(1)
	vehicle, pErr := model.GetInitialVehicle(lib.BaseURL+"vehicles/"+vehicleID, &wg)
	if pErr != nil {
		fmt.Println(pErr.Error())
	}
	wg.Wait()

	page.PageTitle = vehicle.Name
	page.MainCard.Title = "Name: " + vehicle.Name
	page.MainCard.SubTitle = "Manufacturer: " + vehicle.Manufacturer + " - " + vehicle.Model
	page.MainCard.Body1 = "Cost: " + vehicle.CostInCredits + " Credits"
	page.MainCard.Body2 = "Length: " + vehicle.Legth + " units"
	page.MainCard.Body3 = "Crew: " + vehicle.Crew + " people"
	page.MainCard.Body4 = "Passengers: " + vehicle.Passengers + " people"
	page.MainCard.Body5 = "Cargo Capacity: " + vehicle.CargoCapacity

	if len(vehicle.Films) > 0 {
		filmChannel := make(chan []model.vehicle)
	}
}
