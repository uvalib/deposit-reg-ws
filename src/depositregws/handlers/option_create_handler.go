package handlers

import (
	"net/http"
	"depositregws/authtoken"
	"depositregws/config"
	"encoding/json"
	"depositregws/logger"
	"fmt"
	"io"
	"io/ioutil"
	"depositregws/api"
	"depositregws/dao"
)

//
// OptionCreate -- create a new option request handler
//
func OptionCreate(w http.ResponseWriter, r *http.Request) {

	token := r.URL.Query().Get("auth")

	// parameters OK ?
	if notEmpty(token) == false {
		status := http.StatusBadRequest
		encodeStandardResponse(w, status, http.StatusText(status))
		return
	}

	// validate the token
	if authtoken.Validate(config.Configuration.AuthTokenEndpoint, token, config.Configuration.ServiceTimeout) == false {
		status := http.StatusForbidden
		encodeStandardResponse(w, status, http.StatusText(status))
		return
	}

	decoder := json.NewDecoder(r.Body)
	option := api.Option{}

	if err := decoder.Decode(&option); err != nil {
		logger.Log(fmt.Sprintf("ERROR: decoding add option request payload %s", err))
		status := http.StatusBadRequest
		encodeStandardResponse(w, status, http.StatusText(status))
		return
	}

	defer io.Copy(ioutil.Discard, r.Body)
	defer r.Body.Close()

	// create the new option
	err := dao.DB.CreateOption( option )
	if err != nil {
		logger.Log(fmt.Sprintf("ERROR: %s\n", err.Error()))
		status := http.StatusInternalServerError
		encodeStandardResponse(w, status,
			fmt.Sprintf("%s (%s)", http.StatusText(status), err))
		return
	}

	status := http.StatusOK
	encodeStandardResponse(w, status, http.StatusText(status))
}

//
// end of file
//
