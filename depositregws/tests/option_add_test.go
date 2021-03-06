package tests

import (
	"github.com/uvalib/deposit-reg-ws/depositregws/api"
	"github.com/uvalib/deposit-reg-ws/depositregws/client"
	"net/http"
	"testing"
)

//
// option add tests
//

func TestOptionAddHappyDay(t *testing.T) {
	expected := http.StatusOK
	option := makeNewOption(departmentType)
	status := client.AddOption(cfg.Endpoint, option, goodToken(cfg.Secret))
	if status != expected {
		t.Fatalf("Expected %v, got %v\n", expected, status)
	}
}

func TestOptionAddEmptyOption(t *testing.T) {
	expected := http.StatusBadRequest
	status := client.AddOption(cfg.Endpoint, api.Option{}, goodToken(cfg.Secret))
	if status != expected {
		t.Fatalf("Expected %v, got %v\n", expected, status)
	}
}

func TestOptionAddBadToken(t *testing.T) {
	option := makeNewOption(departmentType)
	expected := http.StatusForbidden
	status := client.AddOption(cfg.Endpoint, option, badToken(cfg.Secret))
	if status != expected {
		t.Fatalf("Expected %v, got %v\n", expected, status)
	}
}

func TestOptionAddDuplicateOption(t *testing.T) {
	expected := http.StatusOK
	option := makeNewOption(departmentType)
	status := client.AddOption(cfg.Endpoint, option, goodToken(cfg.Secret))
	if status != expected {
		t.Fatalf("Expected %v, got %v\n", expected, status)
	}

	expected = http.StatusUnprocessableEntity
	status = client.AddOption(cfg.Endpoint, option, goodToken(cfg.Secret))
	if status != expected {
		t.Fatalf("Expected %v, got %v\n", expected, status)
	}
}

//
// end of file
//
