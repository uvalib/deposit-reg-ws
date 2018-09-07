package handlers

import (
	"net/http"
	"github.com/uvalib/deposit-reg-ws/depositregws/authtoken"
	"github.com/uvalib/deposit-reg-ws/depositregws/config"
	"encoding/json"
	"github.com/uvalib/deposit-reg-ws/depositregws/logger"
	"fmt"
	"io"
	"io/ioutil"
	"github.com/uvalib/deposit-reg-ws/depositregws/api"
	"github.com/uvalib/deposit-reg-ws/depositregws/dao"
	"strings"
)

//
// OptionCreate -- create a new option request handler
//
func OptionCreate(w http.ResponseWriter, r *http.Request) {

	token := r.URL.Query().Get("auth")

	// parameters OK ?
	if isEmpty(token) == true {
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

	// payload OK ?
	if isEmpty( option.Option ) == true || isEmpty( option.Value ) == true {
		status := http.StatusBadRequest
		encodeStandardResponse(w, status, http.StatusText(status))
		return
	}

	// create the new option
	err := dao.DB.CreateOption( option )
	if err != nil {
		logger.Log(fmt.Sprintf("ERROR: %s\n", err.Error()))
		status := http.StatusInternalServerError
		// check for a constraint violation
		if strings.Contains( err.Error( ), "Duplicate entry" ) == true {
			status = http.StatusUnprocessableEntity
		}
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
