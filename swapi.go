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
	log.Fatal(http.ListenAndServe(":8080", nil))
}
