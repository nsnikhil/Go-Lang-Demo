package routing_test

import (
	"awesomeProject/trip/domain"
	"awesomeProject/trip/routing"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"io/ioutil"
	"net/http/httptest"
	"testing"
)

type mockedService struct {
	mock.Mock
}

func (mock *mockedService) GetTrip(id int) domain.Trip {
	return domain.Trip{Id: 1, Origin: "Kings Landing", Destination: "The wall"}
}

func (mock *mockedService) CreateTrip(trip domain.Trip) {

}

var m = new(mockedService)

func TestAssignHandlers(t *testing.T) {
	handler := routing.MyHandlers{TripService: m}

	req := httptest.NewRequest("GET", "/trip", nil)
	wri := httptest.NewRecorder()

	handler.GetTripHandler(wri, req)

	b := wri.Result().Body

	br, _ := ioutil.ReadAll(b)

	assert.Equal(t, string(br), "{\"Id\":1,\"Origin\":\"Kings Landing\",\"Destination\":\"The wall\"}\n", nil)

}
