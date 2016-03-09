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

func TestHealthCheck( t *testing.T ) {
    expected := http.StatusOK
    status := client.HealthCheck( cfg.Endpoint )
    if status != expected {
        t.Fatalf( "Expected %v, got %v\n", expected, status )
    }
}

func TestGetHappyDay( t *testing.T ) {
    expected := http.StatusOK
    status, details := client.Get( cfg.Endpoint, goodId, goodToken )
    if status != expected {
        t.Fatalf( "Expected %v, got %v\n", expected, status )
    }
    ensureValidDetails( t, details )
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

func ensureValidDetails( t *testing.T, details [] * api.Registration ) {

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

func emptyField( field string ) bool {
    return len( strings.TrimSpace( field ) ) == 0
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

    log.Printf( "endpoint [%s]\n", c.Endpoint )
    log.Printf( "token    [%s]\n", c.Token )

    return c
}