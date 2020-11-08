package main

import (
	"fmt"
	"log"
	"net/http"
	"swapigo/controllers"
)

func main() {
	fmt.Println("Listening on Port 8080")
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("assets"))))
	http.HandleFunc("/people/1", controllers.HandlePerson)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// func handlePerson(w http.ResponseWriter, r *http.Request) {
// 	// get id
// 	fmt.Println("we in here")
// 	wg.Add(1)
// 	person, e := api.GetPerson(baseURL+"people/1/", &wg)
// 	if e != nil {
// 		panic(e)
// 	}
// 	wg.Wait()
// 	// fmt.Println("Returning Response", person.Homeworld)
// 	wg.Add(1)
// 	planet, planetErr := api.GetPlanet(person.Homeworld, &wg)
// 	if planetErr != nil {
// 		panic(planetErr)
// 	}

// 	// films, filmsErr := api.Get

// 	wg.Wait()
// 	fmt.Println("hell yeah", planet)
// 	// if p, e := api.GetPerson(1); e != nil {
// 	// fmt.Println("we in here")
// 	// api.GetPlanet(p.Homeworld)
// 	// defer wg.Done()
// 	// }
// 	// p, e := api.GetPerson(1)
// 	// fmt.Println("p.Homeworld", p.Homeworld)
// 	// fmt.Println("e", e)

// 	// go api.GetPlanet(p.Homeworld)
// }
