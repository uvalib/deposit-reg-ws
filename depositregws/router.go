package main

import (
	"github.com/gorilla/mux"
	"github.com/uvalib/deposit-reg-ws/depositregws/handlers"
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
		"FaveIcon",
		"GET",
		"/favicon.ico",
		handlers.FavIconHandler,
	},

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
		"OptionMapGet",
		"GET",
		"/optionmap",
		handlers.OptionMapGet,
	},

	route{
		"OptionMapUpdate",
		"PUT",
		"/optionmap",
		handlers.OptionMapUpdate,
	},

	route{
		"OptionsGet",
		"GET",
		"/options",
		handlers.OptionsGet,
	},

	route{
		"OptionCreate",
		"POST",
		"/options",
		handlers.OptionCreate,
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

// NewRouter -- build and return the router
func NewRouter() *mux.Router {

	router := mux.NewRouter().StrictSlash(true)

	// add the route for the prometheus metrics
	//router.Handle("/metrics", HandlerLogger(promhttp.Handler(), "promhttp.Handler"))

	for _, route := range routes {

		var handler http.Handler = route.HandlerFunc
		handler = HandlerLogger(handler, route.Name)
		//handler = prometheus.InstrumentHandler(route.Name, handler)

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
