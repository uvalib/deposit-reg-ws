package handlers

import (
	"net/http"
)

//
// OptionMapUpdate -- update an existing option map request handler
//
func OptionMapUpdate(w http.ResponseWriter, r *http.Request) {

	status := http.StatusOK
	encodeStandardResponse(w, status, http.StatusText(status))
}

//
// end of file
//
