package handlers

import (
    "log"
    "strings"
    "encoding/json"
    "net/http"
    "depositregws/api"
)

func EncodeStandardResponse( w http.ResponseWriter, status int, message string, details [] * api.Registration ) {

    log.Printf( "Status: %d (%s)\n", status, message )
    jsonAttributes( w )
    coorsAttributes( w )
    w.WriteHeader( status )
    if err := json.NewEncoder(w).Encode( api.StandardResponse{ Status: status, Message: message, Details: details } ); err != nil {
        log.Fatal( err )
    }
}

func EncodeOptionsResponse( w http.ResponseWriter, status int, message string, options * api.Options ) {

    jsonAttributes( w )
    coorsAttributes( w )
    w.WriteHeader( status )
    if err := json.NewEncoder(w).Encode( api.OptionsResponse{ Status: status, Message: message, Options: options } ); err != nil {
        log.Fatal( err )
    }
}

func EncodeHealthCheckResponse( w http.ResponseWriter, status int, message string ) {
    healthy := status == http.StatusOK
    jsonAttributes( w )
    w.WriteHeader( status )
    if err := json.NewEncoder(w).Encode( api.HealthCheckResponse { CheckType: api.HealthCheckResult{ Healthy: healthy, Message: message } } ); err != nil {
        log.Fatal( err )
    }
}

func jsonAttributes( w http.ResponseWriter ) {
    w.Header( ).Set( "Content-Type", "application/json; charset=UTF-8" )
}

func coorsAttributes( w http.ResponseWriter ) {
    w.Header( ).Set( "Access-Control-Allow-Origin", "*" )
    w.Header( ).Set( "Access-Control-Allow-Headers", "Content-Type" )
}

func jsonResponse( w http.ResponseWriter ) {
    w.Header( ).Set( "Content-Type", "application/json; charset=UTF-8" )
}

func NotEmpty( param string ) bool {
    return len( strings.TrimSpace( param ) ) != 0
}

func StatusHelper( status int ) ( int, string ) {
    return status, http.StatusText( status )
}