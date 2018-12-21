package routing

import (
	"awesomeProject/trip/service"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type MyHandlers struct {
	TripService *service.TripServiceImpl
}

//noinspection GoUnhandledErrorResult
func (h *MyHandlers) GetTripHandler(writer http.ResponseWriter, request *http.Request) {
	parameters := mux.Vars(request)
	id, _ := strconv.Atoi(parameters["id"])
	json.NewEncoder(writer).Encode((*h.TripService).GetTrip(id))
}

