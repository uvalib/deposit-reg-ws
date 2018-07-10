package handlers

import (
	"net/http"
	"depositregws/authtoken"
	"depositregws/config"
	"encoding/json"
	"depositregws/api"
	"depositregws/logger"
	"fmt"
	"io"
	"io/ioutil"
	"depositregws/dao"
	"strings"
)

//
// OptionMapUpdate -- update an existing option map request handler
//
func OptionMapUpdate(w http.ResponseWriter, r *http.Request) {

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
	optionMap := api.DepartmentMap{}

	if err := decoder.Decode(&optionMap); err != nil {
		logger.Log(fmt.Sprintf("ERROR: decoding update option map request payload %s", err))
		status := http.StatusBadRequest
		encodeStandardResponse(w, status, http.StatusText(status))
		return
	}

	defer io.Copy(ioutil.Discard, r.Body)
	defer r.Body.Close()

	// payload OK ?
	if isEmpty( optionMap.Department ) == true || anyEmpty( optionMap.Degrees ) == true {
		status := http.StatusBadRequest
		encodeStandardResponse(w, status, http.StatusText(status))
		return
	}

	// update the option map
	err := dao.DB.UpdateOptionMap( optionMap )
	if err != nil {
		logger.Log(fmt.Sprintf("ERROR: %s\n", err.Error()))
		status := http.StatusInternalServerError
		// check for a value not found
		if strings.Contains( err.Error( ), "does not exist" ) == true {
			status = http.StatusNotFound
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
