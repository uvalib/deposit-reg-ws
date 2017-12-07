package main

import (
	"depositregws/handlers"
	"github.com/gorilla/mux"
	"net/http"
)

type route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type routeSlice []route

var routes = routeSlice{

	route{
		"HealthCheck",
		"GET",
		"/healthcheck",
		handlers.HealthCheck,
	},

	route{
		"VersionInfo",
		"GET",
		"/version",
		handlers.VersionInfo,
	},

	route{
		"RuntimeInfo",
		"GET",
		"/runtime",
		handlers.RuntimeInfo,
	},

	route{
		"OptionMapGet",
		"GET",
		"/optionmap",
		handlers.OptionMapGet,
	},

	route{
		"OptionMapGet",
		"GET",
		"/options",
		handlers.OptionsGet,
	},

	route{
		"RegistrationGet",
		"GET",
		"/{id}",
		handlers.RegistrationGet,
	},

	route{
		"RegistrationSearch",
		"GET",
		"/",
		handlers.RegistrationSearch,
	},

	route{
		"RegistrationCreate",
		"POST",
		"/",
		handlers.RegistrationCreate,
	},

	route{
		"RegistrationCreate",
		"OPTIONS",
		"/",
		handlers.RegistrationCreateOptions,
	},

	/*
	   route{
	       "RegistrationUpdate",
	       "PUT",
	       "/{id}",
	       handlers.RegistrationUpdate,
	   },

	*/
	route{
		"RegistrationDelete",
		"DELETE",
		"/{id}",
		handlers.RegistrationDelete,
	},
}

//
// NewRouter -- build and return the router
//
func NewRouter() *mux.Router {

	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {

		var handler http.Handler

		handler = route.HandlerFunc
		handler = HandlerLogger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	return router
}

//
// end of file
//
