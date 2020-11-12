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
	http.HandleFunc("/", controllers.HandlePeople)
	http.HandleFunc("/people", controllers.HandlePeople)
	http.HandleFunc("/person", controllers.HandlePerson)
	http.HandleFunc("/vehicles", controllers.HandleVehicles)
	http.HandleFunc("/vehicle", controllers.HandleVehicle)
	http.HandleFunc("/planets", controllers.HandlePlanets)
	http.HandleFunc("/planet", controllers.HandlePlanet)
	http.HandleFunc("/films", controllers.HandleFilms)
	http.HandleFunc("/film", controllers.HandleFilm)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
