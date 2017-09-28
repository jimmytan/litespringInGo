package routers

import (
	"github.com/gorilla/mux"
)

func SetEventRoutes(router *mux.Router) *mux.Router  {
	eventRouter := mux.NewRouter()
	eventRouter.HandleFunc("/events")
}
