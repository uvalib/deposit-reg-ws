package main

import (
   "net/http"
   "github.com/gorilla/mux"
)

type Route struct {
   Name        string
   Method      string
   Pattern     string
   HandlerFunc http.HandlerFunc
}

type Routes [] Route

var routes = Routes{

    Route{
        "HealthCheck",
        "GET",
        "/healthcheck",
        HealthCheck,
    },

    Route{
        "RegistrationGet",
        "GET",
        "/{id}",
        RegistrationGet,
    },
/*
    Route{
        "RegistrationSearch",
        "GET",
        "/",
        RegistrationSearch,
    },

    Route{
        "RegistrationCreate",
        "POST",
        "/",
        RegistrationCreate,
    },

    Route{
        "RegistrationUpdate",
        "PUT",
        "/{id}",
        RegistrationUpdate,
    },

    Route{
        "RegistrationDelete",
        "DEL",
        "/{id}",
        RegistrationDelete,
    },
    */
}

func NewRouter( ) *mux.Router {

   router := mux.NewRouter().StrictSlash( true )
   for _, route := range routes {

      var handler http.Handler

      handler = route.HandlerFunc
      handler = Logger( handler, route.Name )

      router.
         Methods( route.Method ).
         Path( route.Pattern ).
         Name( route.Name ).
         Handler( handler )
   }

   return router
}
