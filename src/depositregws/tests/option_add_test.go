package tests

import (
	"depositregws/client"
	"net/http"
	"testing"
	"depositregws/api"
)

//
// option add tests
//

func TestOptionAddHappyDay(t *testing.T) {
	expected := http.StatusOK
	option := makeNewOption( departmentType )
	status := client.AddOption(cfg.Endpoint, option, goodToken)
	if status != expected {
		t.Fatalf("Expected %v, got %v\n", expected, status)
	}
}

func TestOptionAddEmptyOption(t *testing.T) {
	expected := http.StatusBadRequest
	status := client.AddOption(cfg.Endpoint, api.Option{}, goodToken)
	if status != expected {
		t.Fatalf("Expected %v, got %v\n", expected, status)
	}
}

func TestOptionAddBadToken(t *testing.T) {
	option := makeNewOption( departmentType )
	expected := http.StatusForbidden
	status := client.AddOption(cfg.Endpoint, option, badToken)
	if status != expected {
		t.Fatalf("Expected %v, got %v\n", expected, status)
	}
}

func TestOptionAddDuplicateOption(t *testing.T) {
	expected := http.StatusOK
	option := makeNewOption( departmentType )
	status := client.AddOption(cfg.Endpoint, option, goodToken)
	if status != expected {
		t.Fatalf("Expected %v, got %v\n", expected, status)
	}

	expected = http.StatusUnprocessableEntity
	status = client.AddOption(cfg.Endpoint, option, goodToken)
	if status != expected {
		t.Fatalf("Expected %v, got %v\n", expected, status)
	}
}

//
// end of file
//
