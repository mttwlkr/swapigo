package controllers

import (
	"fmt"
	"html/template"
	"net/http"
	"swapigo/lib"
	"swapigo/models"
	"sync"
)

// PeoplePageResponse is the response to a request for a page of people
type PeoplePageResponse struct {
	PageNumber   int
	NextPage     string
	PreviousPage string
	People       []model.SubCard
	PageTitle    string
}

// HandlePeople handles a single get request for a page of people
func HandlePeople(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("views/index-page.html"))
	var page PeoplePageResponse

	var wg sync.WaitGroup
	fmt.Println("in handlePeople")

	wg.Add(1)

	queryPageNumber := "1"
	keys := r.URL.Query()["id"]
	if len(keys) > 0 {
		queryPageNumber = keys[0]
	}

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

	page.People = make([]model.SubCard, 0)
	for _, person := range peoplePage.Results {
		page.People = append(page.People, model.SubCard{
			Title:     "Name: " + person.Name,
			SubTitle:  "URL: " + person.URL,
			SubTitle2: "BirthYear: " + person.BirthYear,
			Body:      person.Name + " is a " + person.Gender + " who has " + person.EyeColor + " eyes, " + person.HairColor + " hair and weighs " + person.Mass + " kilograms",
			URL:       "/person?id=" + lib.GetIDFromString(person.URL),
		})
	}

	tmpl.Execute(w, page)
}
