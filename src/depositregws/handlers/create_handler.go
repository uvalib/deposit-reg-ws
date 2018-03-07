package handlers

import (
	"depositregws/api"
	"depositregws/authtoken"
	"depositregws/config"
	"depositregws/dao"
	"depositregws/logger"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
)

//
// RegistrationCreate -- create registration handler
//
func RegistrationCreate(w http.ResponseWriter, r *http.Request) {

	token := r.URL.Query().Get("auth")

	// parameters OK ?
	if notEmpty(token) == false {
		status := http.StatusBadRequest
		encodeStandardResponse(w, status, http.StatusText(status), nil)
		return
	}

	// validate the token
	if authtoken.Validate(config.Configuration.AuthTokenEndpoint, token, config.Configuration.ServiceTimeout) == false {
		status := http.StatusForbidden
		encodeStandardResponse(w, status, http.StatusText(status), nil)
		return
	}

	decoder := json.NewDecoder(r.Body)
	reg := api.Registration{}

	if err := decoder.Decode(&reg); err != nil {
		status := http.StatusBadRequest
		encodeStandardResponse(w, status, http.StatusText(status), nil)
		return
	}

	defer io.Copy(ioutil.Discard, r.Body)
	defer r.Body.Close()

	// create results list
	results := make([]*api.Registration, 0)

	// split the user list of appropriate
	users := strings.Split(reg.For, ",")

	for _, u := range users {

		reg.For = strings.TrimSpace(u)
		rg, err := dao.DB.CreateDepositRequest(reg)
		if err != nil {
			logger.Log(fmt.Sprintf("ERROR: %s\n", err.Error()))
			status := http.StatusInternalServerError
			encodeStandardResponse(w, status,
				fmt.Sprintf("%s (%s)", http.StatusText(status), err),
				nil)
			return
		}

		results = append(results, rg)
	}

	status := http.StatusOK
	encodeStandardResponse(w, status, http.StatusText(status), results)
}

//
// RegistrationCreateOptions -- create registration options handler
//
func RegistrationCreateOptions(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Access-Control-Allow-Methods", "POST")
	encodeStandardResponse(w, http.StatusOK, http.StatusText(http.StatusOK), nil)
}

//
// end of file
//
