package main

import (
    "io/ioutil"
    "log"
    "testing"
    "strings"
    "depositregws/client"
    "depositregws/api"
    "gopkg.in/yaml.v2"
    "net/http"
)

type TestConfig struct {
    Endpoint  string
    Token     string
}

var cfg = loadConfig( )

var goodId = "1"
var notFoundId = "x"
var goodToken = cfg.Token
var badToken = "badness"
var empty = " "

//
// healthcheck tests
//

func TestHealthCheck( t *testing.T ) {
    expected := http.StatusOK
    status := client.HealthCheck( cfg.Endpoint )
    if status != expected {
        t.Fatalf( "Expected %v, got %v\n", expected, status )
    }
}

//
// options tests
//

func TestOptionsHappyDay( t *testing.T ) {
    expected := http.StatusOK
    status, options := client.Options( cfg.Endpoint )
    if status != expected {
        t.Fatalf( "Expected %v, got %v\n", expected, status )
    }
    ensureValidOptions( t, options )
}

//
// get tests
//

func TestGetHappyDay( t *testing.T ) {
    expected := http.StatusOK
    status, details := client.Get( cfg.Endpoint, goodId, goodToken )
    if status != expected {
        t.Fatalf( "Expected %v, got %v\n", expected, status )
    }
    ensureValidRegistrations( t, details )
}

func TestGetEmptyId( t *testing.T ) {
    expected := http.StatusBadRequest
    status, _ := client.Get( cfg.Endpoint, empty, goodToken )
    if status != expected {
        t.Fatalf( "Expected %v, got %v\n", expected, status )
    }
}

func TestGetNotFoundId( t *testing.T ) {
    expected := http.StatusNotFound
    status, _ := client.Get( cfg.Endpoint, notFoundId, goodToken )
    if status != expected {
        t.Fatalf( "Expected %v, got %v\n", expected, status )
    }
}

func TestGetBadToken( t *testing.T ) {
    expected := http.StatusForbidden
    status, _ := client.Get( cfg.Endpoint, goodId, badToken )
    if status != expected {
        t.Fatalf( "Expected %v, got %v\n", expected, status )
    }
}

//
// search tests
//

//
// create tests
//

func TestCreateHappyDay( t *testing.T ) {
    reg := makeRegistration( )
    expected := http.StatusOK
    status, _ := client.Create( cfg.Endpoint, reg, goodToken )
    if status != expected {
        t.Fatalf( "Expected %v, got %v\n", expected, status )
    }
}

func TestCreateBadRegistration( t *testing.T ) {
    expected := http.StatusBadRequest
    status, _ := client.Create( cfg.Endpoint, api.Registration{ }, goodToken )
    if status != expected {
        t.Fatalf( "Expected %v, got %v\n", expected, status )
    }
}

func TestCreateBadToken( t *testing.T ) {
    reg := makeRegistration( )
    expected := http.StatusForbidden
    status, _ := client.Create( cfg.Endpoint, reg, badToken )
    if status != expected {
        t.Fatalf( "Expected %v, got %v\n", expected, status )
    }
}

//
// update tests
//

//
// delete tests
//

func TestDeleteHappyDay( t *testing.T ) {
    newId := createNewReg( t )
    expected := http.StatusOK
    status := client.Delete( cfg.Endpoint, newId, goodToken )
    if status != expected {
        t.Fatalf( "Expected %v, got %v\n", expected, status )
    }
}

func TestDeleteEmptyId( t *testing.T ) {
    expected := http.StatusBadRequest
    status := client.Delete( cfg.Endpoint, empty, goodToken )
    if status != expected {
        t.Fatalf( "Expected %v, got %v\n", expected, status )
    }
}

func TestDeleteNotFoundId( t *testing.T ) {
    expected := http.StatusNotFound
    status := client.Delete( cfg.Endpoint, notFoundId, goodToken )
    if status != expected {
        t.Fatalf( "Expected %v, got %v\n", expected, status )
    }
}

func TestDeleteBadToken( t *testing.T ) {
    expected := http.StatusForbidden
    status := client.Delete( cfg.Endpoint, goodId, badToken )
    if status != expected {
        t.Fatalf( "Expected %v, got %v\n", expected, status )
    }
}

func ensureValidRegistrations( t *testing.T, details [] * api.Registration ) {

    for _, e := range details {
        if emptyField( e.Id ) ||
           emptyField( e.For ) ||
           emptyField( e.School ) ||
           emptyField( e.Degree ) ||
           emptyField( e.RequestDate ) ||
           emptyField( e.Status ) {
            t.Fatalf( "Expected non-empty field but one is empty\n" )
        }
    }
}

func ensureValidOptions( t *testing.T, options * api.Options ) {

    for _, f := range options.School {
        if emptyField( f ) {
            t.Fatalf( "Expected non-empty school field but one is empty\n" )
        }
    }
    for _, f := range options.Degree {
        if emptyField( f ) {
            t.Fatalf( "Expected non-empty degree field but one is empty\n" )
        }
    }
}

func createNewReg( t *testing.T ) string {
    reg := makeRegistration( )
    expected := http.StatusOK
    status, result := client.Create( cfg.Endpoint, reg, goodToken )
    if status != expected {
        t.Fatalf( "Expected %v, got %v\n", expected, status )
    }
    if result == nil {
        t.Fatalf( "No registration details returned" )
    }

    return result.Id
}

func emptyField( field string ) bool {
    return len( strings.TrimSpace( field ) ) == 0
}

func makeRegistration( ) api.Registration {
    return api.Registration{
        For: "dpg3k",
        School: "Engineering",
        Degree: "Ph.D" }
}

func loadConfig( ) TestConfig {

    data, err := ioutil.ReadFile( "service_test.yml" )
    if err != nil {
        log.Fatal( err )
    }

    var c TestConfig
    if err := yaml.Unmarshal( data, &c ); err != nil {
        log.Fatal( err )
    }

    log.Printf( "Test config; endpoint   [%s]\n", c.Endpoint )
    log.Printf( "Test config; auth token [%s]\n", c.Token )

    return c
}