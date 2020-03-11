package tests

import (
	"github.com/uvalib/deposit-reg-ws/depositregws/api"
	"github.com/uvalib/deposit-reg-ws/depositregws/client"
	"net/http"
	"testing"
)

//
// create request tests
//

func TestCreateRequestSingle(t *testing.T) {
	reg := makeSingleRegistration()
	expected := http.StatusOK
	status, details := client.CreateDepositRequest(cfg.Endpoint, reg, goodToken(cfg.Secret))
	if status != expected {
		t.Fatalf("Expected %v, got %v\n", expected, status)
	}

	if details == nil || len(details) != 1 {
		t.Fatalf("Incomplete registration details returned")
	}

	ensureValidRegistrations(t, details)
}

func TestCreateRequestMulti(t *testing.T) {
	reg := makeMultiRegistration()
	expected := http.StatusOK
	status, details := client.CreateDepositRequest(cfg.Endpoint, reg, goodToken(cfg.Secret))
	if status != expected {
		t.Fatalf("Expected %v, got %v\n", expected, status)
	}

	if details == nil || len(details) != 2 {
		t.Fatalf("Incomplete registration details returned")
	}

	ensureValidRegistrations(t, details)
}

func TestCreateRequestBadRegistration(t *testing.T) {
	expected := http.StatusBadRequest
	status, _ := client.CreateDepositRequest(cfg.Endpoint, api.Registration{}, goodToken(cfg.Secret))
	if status != expected {
		t.Fatalf("Expected %v, got %v\n", expected, status)
	}
}

func TestCreateRequestBadToken(t *testing.T) {
	reg := makeSingleRegistration()
	expected := http.StatusForbidden
	status, _ := client.CreateDepositRequest(cfg.Endpoint, reg, badToken(cfg.Secret))
	if status != expected {
		t.Fatalf("Expected %v, got %v\n", expected, status)
	}
}

//
// end of file
//
