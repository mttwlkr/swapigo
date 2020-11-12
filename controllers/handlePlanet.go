package controllers

import (
	"fmt"
	"html/template"
	"net/http"
	"swapigo/lib"
	"swapigo/model"
	"sync"
)

// HandlePlanet gets a page of
func HandlePlanet(w http.ResponseWriter, r *http.Request) {
	keys := r.URL.Query()["id"]
	planetID := keys[0]

	var wg sync.WaitGroup
	tmpl := template.Must(template.ParseFiles("views/detail-page.html"))
	var page model.DetailPageResponse

	wg.Add(1)
	planet, pErr := model.GetInitialPlanet(lib.BaseURL+"planets/"+planetID+"/", &wg)
	if pErr != nil {
		fmt.Println(pErr.Error())
	}
	wg.Wait()

	fmt.Println("planet", planet)

	page.PageTitle = "Planet"
	page.MainCard.Title = "Name: " + planet.Name
	page.MainCard.Body1 = "Rotation/Orbital: " + planet.RotationPeriod + planet.OrbitalPeriod
	page.MainCard.Body2 = "Population: " + planet.Population
	page.MainCard.Body3 = "Climate: " + planet.Climate
	page.MainCard.Body4 = "Gravity: " + planet.Gravity
	page.MainCard.Body5 = "Terrain: " + planet.Terrain

	if len(planet.Films) > 0 {
		filmChannel := make(chan []model.Film)
		go planet.GetFilms(filmChannel)
		films := <-filmChannel
		page.Cards1 = make([]model.SubCard, 0)
		page.Cards1Title = "Films"
		for _, film := range films {
			page.Cards1 = append(page.Cards1, model.GetFilmCard(film))
		}
	}

	if len(planet.Residents) > 0 {
		residentsChannel := make(chan []model.Person)
		go planet.GetResidents(residentsChannel)
		residents := <-residentsChannel
		page.Cards2Title = "Residents"
		page.Cards2 = make([]model.SubCard, 0)
		for _, resident := range residents {
			page.Cards2 = append(page.Cards2, model.GetPersonCard(resident))
		}
	}

	tmpl.Execute(w, page)
}
