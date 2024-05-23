package handlers

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/uvalib/deposit-reg-ws/depositregws/authtoken"
	"github.com/uvalib/deposit-reg-ws/depositregws/config"
	"github.com/uvalib/deposit-reg-ws/depositregws/dao"
	"github.com/uvalib/deposit-reg-ws/depositregws/logger"
	"net/http"
)

// RegistrationDelete - registration delete handler
func RegistrationDelete(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id := vars["id"]
	token := r.URL.Query().Get("auth")

	// parameters OK ?
	if isEmpty(id) == true || isEmpty(token) == true {
		status := http.StatusBadRequest
		encodeRegistrationResponse(w, status, http.StatusText(status), nil)
		return
	}

	// validate the token
	if authtoken.Validate(config.Configuration.SharedSecret, token) == false {
		status := http.StatusForbidden
		encodeRegistrationResponse(w, status, http.StatusText(status), nil)
		return
	}

	// get the request details
	count, err := dao.Store.DeleteDepositRequest(id)
	if err != nil {
		logger.Log(fmt.Sprintf("ERROR: %s", err.Error()))
		status := http.StatusInternalServerError
		encodeRegistrationResponse(w, status,
			fmt.Sprintf("%s (%s)", http.StatusText(status), err),
			nil)
		return
	}

	if count == 0 {
		status := http.StatusNotFound
		encodeRegistrationResponse(w, status, http.StatusText(status), nil)
		return
	}

	status := http.StatusOK
	encodeRegistrationResponse(w, status, http.StatusText(status), nil)
}

//
// end of file
//
