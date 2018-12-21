package main

import (
	"awesomeProject/trip/routing"
	"awesomeProject/trip/service"
	"log"
	"net/http"
)

func main() {
	h := routing.MyHandlers{TripService: &service.TripServiceImpl{}}
	r := routing.MyRouter{Handlers: &h}
	router := r.AssignHandlers()
	log.Fatal(http.ListenAndServe(":12345", router))
}

