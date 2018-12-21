package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type Person struct {
	Id        string   `json:"id,omitempty"`
	FirstName string   `json:"firstname,omitempty"`
	LastName  string   `json:"lastname,omitempty"`
	Address   *Address `json:"address,omitempty"`
}

type Address struct {
	City  string `json:"city,omitempty"`
	State string `json:"state,omitempty"`
}

var person []Person

func getPerson(w http.ResponseWriter, req *http.Request) {
	fmt.Println("called")
	json.NewEncoder(w).Encode(person)
}

func getPersonById(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	for _, p := range person {
		if p.Id == params["id"] {
			json.NewEncoder(w).Encode(p)
			return
		}
	}
	json.NewEncoder(w).Encode(&Person{})
}

func createPerson(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	var newPerson Person
	json.NewDecoder(req.Body).Decode(&newPerson)
	newPerson.Id = params["id"]
	person = append(person, newPerson)
	json.NewEncoder(w).Encode(person)
}

func deletePerson(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	for i, p := range person {
		if p.Id == params["id"] {
			person = append(person[:i], person[i+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(person)
}

func main() {
	router := mux.NewRouter()

	person = append(person, Person{Id: "1", FirstName: "nikhil", LastName: "soni", Address: &Address{City: "bangalore", State: "karnataka"}})
	person = append(person, Person{Id: "2", FirstName: "chit", LastName: "singh"})
	router.HandleFunc("/people", getPerson).Methods("GET")
	router.HandleFunc("/people/{id}", getPersonById).Methods("GET")
	router.HandleFunc("/people/{id}", createPerson).Methods("POST")
	router.HandleFunc("/people/{id}", deletePerson).Methods("DELETE")


	fmt.Println(router)

	log.Fatal(http.ListenAndServe(":12345", router))
}
