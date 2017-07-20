package handlers

import (
	"depositregws/authtoken"
	"depositregws/config"
	"depositregws/dao"
	"depositregws/logger"
	"fmt"
	"net/http"
)

func RegistrationSearch(w http.ResponseWriter, r *http.Request) {

	token := r.URL.Query().Get("auth")
	id := r.URL.Query().Get("later")

	// parameters OK ?
	if NotEmpty(token) == false || NotEmpty(id) == false {
		status := http.StatusBadRequest
		EncodeStandardResponse(w, status, http.StatusText(status), nil)
		return
	}

	// validate the token
	if authtoken.Validate(config.Configuration.AuthTokenEndpoint, token, config.Configuration.Timeout) == false {
		status := http.StatusForbidden
		EncodeStandardResponse(w, status, http.StatusText(status), nil)
		return
	}

	// get the request details
	reqs, err := dao.Database.SearchDepositRequest(id)
	if err != nil {
		logger.Log(fmt.Sprintf("ERROR: %s\n", err.Error()))
		status := http.StatusInternalServerError
		EncodeStandardResponse(w, status,
			fmt.Sprintf("%s (%s)", http.StatusText(status), err),
			nil)
		return
	}

	if reqs == nil || len(reqs) == 0 {
		status := http.StatusNotFound
		EncodeStandardResponse(w, status, http.StatusText(status), nil)
		return
	}

	status := http.StatusOK
	EncodeStandardResponse(w, status, http.StatusText(status), reqs)
}
