package controllers

import (
	"fmt"
	"html/template"
	"net/http"
	"swapigo/lib"
	"swapigo/models"
	"sync"
)

// HandlePlanets handles a single get request for a page of planets
func HandlePlanets(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("views/index-page.html"))
	var page model.IndexPageResponse

	var wg sync.WaitGroup
	fmt.Println("in handlePeople")

	wg.Add(1)

	queryPageNumber := "1"
	keys := r.URL.Query()["id"]
	if len(keys) > 0 {
		queryPageNumber = keys[0]
	}

	planetPage, err := model.GetInitialPlanets(lib.BaseURL+"planets/?page="+queryPageNumber, &wg)

	if err != nil {
		fmt.Println(err.Error())
	}

	wg.Wait()

	page.PageTitle = "Planets"
	page.NextPage = "/planets?id=" + lib.GetIDFromString(planetPage.Next)
	page.PreviousPage = "/planets?id=" + lib.GetIDFromString(planetPage.Previous)

	for _, planet := range planetPage.Results {
		page.Cards = append(page.Cards, model.SubCard{
			Title:     "Name: " + planet.Name,
			SubTitle:  "Climate: " + planet.Climate,
			SubTitle2: "Gravity: " + planet.Gravity,
			Body:      "Population: " + planet.Population,
			URL:       "planet?id=" + lib.GetIDFromString(planet.URL),
		})
	}
	tmpl.Execute(w, page)
}
