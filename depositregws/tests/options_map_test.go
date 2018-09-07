package tests

import (
	"github.com/uvalib/deposit-reg-ws/depositregws/client"
	"net/http"
	"testing"
)

//
// options map tests
//

func TestOptionsMapHappyDay(t *testing.T) {
	expected := http.StatusOK
	status, options := client.GetMappedOptions(cfg.Endpoint)
	if status != expected {
		t.Fatalf("Expected %v, got %v\n", expected, status)
	}

	ensureValidMappedOptions( t, options )
}

//
// end of file
//
