package handlers

import (
	"fmt"
	"github.com/uvalib/deposit-reg-ws/depositregws/dao"
	"github.com/uvalib/deposit-reg-ws/depositregws/logger"
	"net/http"
)

// HealthCheck -- do the healthcheck
func HealthCheck(w http.ResponseWriter, r *http.Request) {

	err := dao.Store.Check()
	if err != nil {
		logger.Log(fmt.Sprintf("ERROR: Datastore reports '%s'", err.Error()))
		encodeHealthCheckResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	encodeHealthCheckResponse(w, http.StatusOK, "")
}

//
// end of file
//
