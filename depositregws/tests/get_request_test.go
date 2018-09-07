package tests

import (
	"github.com/uvalib/deposit-reg-ws/depositregws/client"
	"net/http"
	"testing"
)

//
// get request tests
//

func TestGetRequestHappyDay(t *testing.T) {
	expected := http.StatusOK
	status, details := client.GetDepositRequest(cfg.Endpoint, goodID, goodToken)
	if status != expected {
		t.Fatalf("Expected %v, got %v\n", expected, status)
	}
	ensureValidRegistrations(t, details)
}

func TestGetRequestEmptyId(t *testing.T) {
	expected := http.StatusBadRequest
	status, _ := client.GetDepositRequest(cfg.Endpoint, empty, goodToken)
	if status != expected {
		t.Fatalf("Expected %v, got %v\n", expected, status)
	}
}

func TestGetRequestNotFoundId(t *testing.T) {
	expected := http.StatusNotFound
	status, _ := client.GetDepositRequest(cfg.Endpoint, notFoundID, goodToken)
	if status != expected {
		t.Fatalf("Expected %v, got %v\n", expected, status)
	}
}

func TestGetRequestBadToken(t *testing.T) {
	expected := http.StatusForbidden
	status, _ := client.GetDepositRequest(cfg.Endpoint, goodID, badToken)
	if status != expected {
		t.Fatalf("Expected %v, got %v\n", expected, status)
	}
}

//
// end of file
//
