package routers

import (
	"github.com/gorilla/mux"
	"github.com/roger8b/imc_calc/handlers"
	"net/http"
)

type Route struct {
	Name       string
	Method     string
	Pattern    string
	HandleFunc http.HandlerFunc
}

type Routes []Route

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandleFunc)
	}

	return router
}

var routes = Routes{
	Route{
		Name:       "Index",
		Method:     "GET",
		Pattern:    "/",
		HandleFunc: handlers.Index,
	},

	Route{
		Name:       "ImcCalc",
		Method:     "POST",
		Pattern:    "/calc",
		HandleFunc: handlers.ImcCalc,
	},
}
