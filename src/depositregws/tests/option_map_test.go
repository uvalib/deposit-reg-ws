package tests

import (
	"depositregws/client"
	"net/http"
	"testing"
)

//
// option map tests
//

func TestOptionMapHappyDay(t *testing.T) {
	expected := http.StatusOK
	status, options := client.GetMappedOptions(cfg.Endpoint)
	if status != expected {
		t.Fatalf("Expected %v, got %v\n", expected, status)
	}
	ensureValidMappedOptions(t, options)
}

//
// end of file
//
