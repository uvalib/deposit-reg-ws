package tests

import (
	"depositregws/client"
	"net/http"
	"testing"
	"depositregws/api"
)

//
// options map update tests
//

func TestOptionsMapUpdateHappyDay(t *testing.T) {
	expected := http.StatusOK
	optionMap := makeOptionMap( )
	status := client.AddOptionMap(cfg.Endpoint, optionMap, goodToken)
	if status != expected {
		t.Fatalf("Expected %v, got %v\n", expected, status)
	}
}

func TestOptionsMapUpdateBadOptionMap(t *testing.T) {
	expected := http.StatusBadRequest
	status := client.AddOptionMap(cfg.Endpoint, api.DepartmentMap{}, goodToken)
	if status != expected {
		t.Fatalf("Expected %v, got %v\n", expected, status)
	}
}

func TestOptionsMapUpdateBadToken(t *testing.T) {
	expected := http.StatusForbidden
	optionMap := makeOptionMap( )
	status := client.AddOptionMap(cfg.Endpoint, optionMap, badToken)
	if status != expected {
		t.Fatalf("Expected %v, got %v\n", expected, status)
	}
}

//
// end of file
//
