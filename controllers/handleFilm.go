package controllers

import (
	"fmt"
	"html/template"
	"net/http"
	"swapigo/lib"
	"swapigo/model"
	"sync"
)

// HandleFilm gets a film and all of its associated values
func HandleFilm(w http.ResponseWriter, r *http.Request) {
	filmID := "1"
	keys := r.URL.Query()["id"]
	if len(keys) > 0 {
		filmID = keys[0]
	}

	tmpl := template.Must(template.ParseFiles("views/detail-page.html"))
	var page model.DetailPageResponse

	fmt.Println("FilmId", filmID)
	var wg sync.WaitGroup
	wg.Add(1)
	film, pErr := model.GetInitialFilm(lib.BaseURL+"films/"+filmID+"/", &wg)
	if pErr != nil {
		fmt.Println("we in the fetch error")
		fmt.Println(pErr.Error())
	}
	wg.Wait()

	page.PageTitle = "Film"
	page.MainCard.Title = "Title: " + film.Title
	page.MainCard.Body1 = "Director: " + film.Director
	page.MainCard.Body2 = "Producer: " + film.Producer
	page.MainCard.Body3 = film.OpeningCrawl

	if len(film.Characters) > 0 {
		characterChannel := make(chan []model.Person)
		go film.GetCharacters(characterChannel)
		characters := <-characterChannel
		page.Cards1Title = "Characters"
		page.Cards1 = make([]model.SubCard, 0)
		for _, character := range characters {
			page.Cards1 = append(page.Cards1, model.GetPersonCard(character))
		}
	}

	// planets
	if len(film.Planets) > 0 {
		planetChannel := make(chan []model.Planet)
		go film.GetPlanets(planetChannel)
		planets := <-planetChannel
		page.Cards2Title = "Planets"
		page.Cards2 = make([]model.SubCard, 0)
		for _, planet := range planets {
			page.Cards2 = append(page.Cards2, model.GetPlanetCard(planet))
		}
	}

	if len(film.Starships) > 0 {
		starshipChannel := make(chan []model.Starship)
		go film.GetStarships(starshipChannel)
		starships := <-starshipChannel
		page.Cards3Title = "Starships"
		page.Cards3 = make([]model.SubCard, 0)
		for _, ship := range starships {
			page.Cards3 = append(page.Cards3, model.GetStarshipCard(ship))
		}
	}

	// vehicles
	if len(film.Vehicles) > 0 {
		vehicleChannel := make(chan []model.Vehicle)
		go film.GetVehicles(vehicleChannel)
		vehicles := <-vehicleChannel
		page.Cards4Title = "Vehicles"
		page.Cards4 = make([]model.SubCard, 0)
		for _, vehicle := range vehicles {
			page.Cards4 = append(page.Cards4, model.GetVehicleCard(vehicle))
		}
	}

	if len(film.Species) > 0 {
		speciesChannel := make(chan []model.Species)
		go film.GetSpecies(speciesChannel)
		species := <-speciesChannel
		page.Cards5Title = "Species"
		page.Cards5 = make([]model.SubCard, 0)
		for _, specie := range species {
			page.Cards5 = append(page.Cards5, model.GetSpeciesCard(specie))
		}
	}

	tmpl.Execute(w, page)
}
