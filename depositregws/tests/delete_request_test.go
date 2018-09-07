package tests

import (
	"github.com/uvalib/deposit-reg-ws/depositregws/client"
	"net/http"
	"testing"
)

//
// delete request tests
//

func TestDeleteRequestHappyDay(t *testing.T) {
	newID := createNewReg(t)
	expected := http.StatusOK
	status := client.DeleteDepositRequest(cfg.Endpoint, newID, goodToken)
	if status != expected {
		t.Fatalf("Expected %v, got %v\n", expected, status)
	}
}

func TestDeleteRequestEmptyId(t *testing.T) {
	expected := http.StatusBadRequest
	status := client.DeleteDepositRequest(cfg.Endpoint, empty, goodToken)
	if status != expected {
		t.Fatalf("Expected %v, got %v\n", expected, status)
	}
}

func TestDeleteRequestNotFoundId(t *testing.T) {
	expected := http.StatusNotFound
	status := client.DeleteDepositRequest(cfg.Endpoint, notFoundID, goodToken)
	if status != expected {
		t.Fatalf("Expected %v, got %v\n", expected, status)
	}
}

func TestDeleteRequestBadToken(t *testing.T) {
	expected := http.StatusForbidden
	status := client.DeleteDepositRequest(cfg.Endpoint, goodID, badToken)
	if status != expected {
		t.Fatalf("Expected %v, got %v\n", expected, status)
	}
}

//
// end of file
//
