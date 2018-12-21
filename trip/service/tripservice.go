package service

import (
	"awesomeProject/trip/domain"
)

type TripService interface {
	GetTrip(int) *domain.Trip
	CreateTrip(domain.Trip)
}

type TripServiceImpl struct{}

func (*TripServiceImpl) GetTrip(id int) *domain.Trip {
	for _, trip := range *TripPointer {
		if trip.Id == id {
			return &trip
		}
	}
	return &domain.Trip{}
}

func (*TripServiceImpl) CreateTrip(trip domain.Trip) {
	Trips = append(Trips, trip)
}
