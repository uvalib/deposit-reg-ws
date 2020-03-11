package tests

import (
	"github.com/uvalib/deposit-reg-ws/depositregws/client"
	"net/http"
	"testing"
)

//
// search request tests
//

func TestSearchRequestHappyDay(t *testing.T) {
	expected := http.StatusOK
	status, details := client.SearchDepositRequest(cfg.Endpoint, "0", goodToken(cfg.Secret))
	if status != expected {
		t.Fatalf("Expected %v, got %v\n", expected, status)
	}
	ensureValidRegistrations(t, details)
}

func TestSearchRequestEmptyId(t *testing.T) {
	expected := http.StatusBadRequest
	status, _ := client.SearchDepositRequest(cfg.Endpoint, empty, goodToken(cfg.Secret))
	if status != expected {
		t.Fatalf("Expected %v, got %v\n", expected, status)
	}
}

func TestSearchRequestBadToken(t *testing.T) {
	expected := http.StatusForbidden
	status, _ := client.SearchDepositRequest(cfg.Endpoint, goodID, badToken(cfg.Secret))
	if status != expected {
		t.Fatalf("Expected %v, got %v\n", expected, status)
	}
}

//
// end of file
//
