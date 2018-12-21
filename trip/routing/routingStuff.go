package routing

import (
	"github.com/gorilla/mux"
)

type MyRouter struct {
	Handlers *MyHandlers
}

func (r *MyRouter) AssignHandlers() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("trip/{id}", r.Handlers.GetTripHandler)
	return router
}
