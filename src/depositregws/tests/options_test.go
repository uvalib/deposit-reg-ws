package tests

import (
   "depositregws/client"
   "net/http"
   "testing"
)

//
// options tests
//

func TestOptionsHappyDay(t *testing.T) {
   expected := http.StatusOK
   status, options := client.GetOptions(cfg.Endpoint)
   if status != expected {
      t.Fatalf("Expected %v, got %v\n", expected, status)
   }
   ensureValidOptions(t, options)
}

//
// end of file
//