package handlers

import (
	"depositregws/authtoken"
	"depositregws/config"
	"depositregws/dao"
	"depositregws/logger"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

//
// RegistrationDelete - registration delete handler
//
func RegistrationDelete(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id := vars["id"]
	token := r.URL.Query().Get("auth")

	// parameters OK ?
	if notEmpty(id) == false || notEmpty(token) == false {
		status := http.StatusBadRequest
		encodeStandardResponse(w, status, http.StatusText(status), nil)
		return
	}

	// validate the token
	if authtoken.Validate(config.Configuration.AuthTokenEndpoint, token, config.Configuration.Timeout) == false {
		status := http.StatusForbidden
		encodeStandardResponse(w, status, http.StatusText(status), nil)
		return
	}

	// get the request details
	count, err := dao.DB.DeleteDepositRequest(id)
	if err != nil {
		logger.Log(fmt.Sprintf("ERROR: %s\n", err.Error()))
		status := http.StatusInternalServerError
		encodeStandardResponse(w, status,
			fmt.Sprintf("%s (%s)", http.StatusText(status), err),
			nil)
		return
	}

	if count == 0 {
		status := http.StatusNotFound
		encodeStandardResponse(w, status, http.StatusText(status), nil)
		return
	}

	status := http.StatusOK
	encodeStandardResponse(w, status, http.StatusText(status), nil)
}

//
// end of file
//
