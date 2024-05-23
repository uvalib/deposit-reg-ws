package handlers

import (
	"fmt"
	"github.com/uvalib/deposit-reg-ws/depositregws/authtoken"
	"github.com/uvalib/deposit-reg-ws/depositregws/config"
	"github.com/uvalib/deposit-reg-ws/depositregws/dao"
	"github.com/uvalib/deposit-reg-ws/depositregws/logger"
	"net/http"
)

// RegistrationSearch -- the search registration handler
func RegistrationSearch(w http.ResponseWriter, r *http.Request) {

	token := r.URL.Query().Get("auth")
	id := r.URL.Query().Get("later")

	// parameters OK ?
	if isEmpty(token) == true || isEmpty(id) == true {
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
	reqs, err := dao.Store.SearchDepositRequest(id)
	if err != nil {
		logger.Log(fmt.Sprintf("ERROR: %s", err.Error()))
		status := http.StatusInternalServerError
		encodeRegistrationResponse(w, status,
			fmt.Sprintf("%s (%s)", http.StatusText(status), err),
			nil)
		return
	}

	if reqs == nil || len(reqs) == 0 {
		status := http.StatusNotFound
		encodeRegistrationResponse(w, status, http.StatusText(status), nil)
		return
	}

	status := http.StatusOK
	encodeRegistrationResponse(w, status, http.StatusText(status), reqs)
}

//
// end of file
//
