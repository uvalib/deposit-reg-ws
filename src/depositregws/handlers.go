package main

import (
    "log"
    "fmt"
    "strings"
    "encoding/json"
    "net/http"
    "github.com/gorilla/mux"
    "depositregws/api"
    "depositregws/authtoken"
    "depositregws/config"
)

func RegistrationGet( w http.ResponseWriter, r *http.Request ) {

    vars := mux.Vars( r )
    id := vars[ "id" ]
    token := r.URL.Query( ).Get( "auth" )

    // parameters OK ?
    if parameterOK( id ) == false || parameterOK( token ) == false {
        encodeStandardResponse( w, http.StatusBadRequest, http.StatusText( http.StatusBadRequest ), nil )
        return
    }

    // validate the token
    if authtoken.Validate( config.Configuration.AuthTokenEndpoint, token ) == false {
        encodeStandardResponse( w, http.StatusForbidden, http.StatusText( http.StatusForbidden ), nil )
        return
    }

    // get the request details
    reqs, err := env.db.Get( id )
    if err != nil {
        log.Println( err )
        encodeStandardResponse( w, http.StatusInternalServerError,
            fmt.Sprintf( "%s (%s)", http.StatusText( http.StatusInternalServerError ), err ),
            nil )
        return
    }

    if reqs == nil || len( reqs ) == 0 {
        encodeStandardResponse( w, http.StatusNotFound, http.StatusText( http.StatusNotFound ), nil )
        return
    }

    encodeStandardResponse( w, http.StatusOK, http.StatusText( http.StatusOK ), reqs )
}



func HealthCheck( w http.ResponseWriter, r *http.Request ) {
    err := env.db.Check( )
    if err != nil {
        encodeHealthCheckResponse( w, http.StatusInternalServerError, err.Error( ) )
        return
    }
    encodeHealthCheckResponse( w, http.StatusOK, "" )
}

func encodeStandardResponse( w http.ResponseWriter, status int, message string, details [] * api.Registration ) {
    jsonResponse( w )
    w.WriteHeader( status )
    if err := json.NewEncoder(w).Encode( api.StandardResponse{ Status: status, Message: message, Details: details } ); err != nil {
        log.Fatal( err )
    }
}

func encodeHealthCheckResponse( w http.ResponseWriter, status int, message string ) {
    healthy := status == http.StatusOK
    jsonResponse( w )
    w.WriteHeader( status )
    if err := json.NewEncoder(w).Encode( api.HealthCheckResponse { CheckType: api.HealthCheckResult{ Healthy: healthy, Message: message } } ); err != nil {
        log.Fatal( err )
    }
}

func jsonResponse( w http.ResponseWriter ) {
    w.Header( ).Set( "Content-Type", "application/json; charset=UTF-8" )
}

func parameterOK( param string ) bool {
    return len( strings.TrimSpace( param ) ) != 0
}
