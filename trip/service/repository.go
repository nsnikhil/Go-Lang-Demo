package service

import (
	"awesomeProject/trip/domain"
)

var TripPointer = &Trips

var Trips = []domain.Trip{
	{Id: 1, Origin: "Bangalore", Destination: "Mysore"},
	{Id: 2, Origin: "Bangalore", Destination: "Chennai"},
}
