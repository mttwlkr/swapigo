package controllers

import (
	"fmt"
	"html/template"
	"net/http"
	"swapigo/lib"
	"swapigo/model"
	"sync"
)

// HandlePeople handles a single get request for a page of people
func HandlePeople(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("views/index-page.html"))
	var page model.IndexPageResponse

	queryPageNumber := "1"
	keys := r.URL.Query()["id"]
	if len(keys) > 0 {
		queryPageNumber = keys[0]
	}

	var wg sync.WaitGroup
	wg.Add(1)

	peoplePage, err := model.GetInitialPeople(lib.BaseURL+"people/?page="+queryPageNumber, &wg)

	if err != nil {
		fmt.Println(err.Error())
	}
	wg.Wait()

	// page.PageNumber = 2
	page.PageTitle = "People"
	// page.PageNumber = int(queryPageNumber)
	page.NextPage = "/people?id=" + lib.GetIDFromString(peoplePage.Next)
	page.PreviousPage = "/people?id=" + lib.GetIDFromString(peoplePage.Previous)

	page.Cards = make([]model.SubCard, 0)
	for _, person := range peoplePage.Results {
		page.Cards = append(page.Cards, model.GetPersonCard(person))
	}

	tmpl.Execute(w, page)
}
