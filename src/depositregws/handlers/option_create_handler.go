package handlers

import (
	"net/http"
)

//
// OptionCreate -- create a new option request handler
//
func OptionCreate(w http.ResponseWriter, r *http.Request) {

	status := http.StatusOK
	encodeStandardResponse(w, status, http.StatusText(status))
}

//
// end of file
//
