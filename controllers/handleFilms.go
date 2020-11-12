package controllers

import (
	"fmt"
	"html/template"
	"net/http"
	"swapigo/lib"
	"swapigo/model"
	"sync"
)

// HandleFilms gets a page of films
func HandleFilms(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("views/index-page.html"))

	var page model.IndexPageResponse

	queryPageNumber := "1"
	keys := r.URL.Query()["id"]
	if len(keys) > 0 {
		queryPageNumber = keys[0]
	}

	var wg sync.WaitGroup
	wg.Add(1)

	filmsPage, err := model.GetInitialFilms(lib.BaseURL+"films/?page="+queryPageNumber, &wg)

	if err != nil {
		fmt.Println(err.Error())
	}
	wg.Wait()

	page.PageTitle = "Films"
	page.NextPage = "/films?id=" + lib.GetIDFromString(filmsPage.Next)
	page.PreviousPage = "/films?id=" + lib.GetIDFromString(filmsPage.Previous)
	fmt.Println("page", page)

	page.Cards = make([]model.SubCard, 0)
	for _, film := range filmsPage.Results {
		page.Cards = append(page.Cards, model.GetFilmCard(film))
	}

	tmpl.Execute(w, page)
}
