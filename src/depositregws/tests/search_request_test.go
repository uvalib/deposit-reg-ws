package tests

import (
	"depositregws/client"
	"net/http"
	"testing"
)

//
// search request tests
//

func TestSearchRequestHappyDay(t *testing.T) {
	expected := http.StatusOK
	status, details := client.SearchDepositRequest(cfg.Endpoint, "0", goodToken)
	if status != expected {
		t.Fatalf("Expected %v, got %v\n", expected, status)
	}
	ensureValidRegistrations(t, details)
}

func TestSearchRequestEmptyId(t *testing.T) {
	expected := http.StatusBadRequest
	status, _ := client.SearchDepositRequest(cfg.Endpoint, empty, goodToken)
	if status != expected {
		t.Fatalf("Expected %v, got %v\n", expected, status)
	}
}

func TestSearchRequestBadToken(t *testing.T) {
	expected := http.StatusForbidden
	status, _ := client.SearchDepositRequest(cfg.Endpoint, goodID, badToken)
	if status != expected {
		t.Fatalf("Expected %v, got %v\n", expected, status)
	}
}

//
// end of file
//
