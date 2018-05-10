package handlers

import (
	"depositregws/dao"
	"net/http"
	"depositregws/logger"
	"fmt"
)

//
// HealthCheck -- do the healthcheck
//
func HealthCheck(w http.ResponseWriter, r *http.Request) {

	err := dao.DB.CheckDB()
	if err != nil {
		logger.Log(fmt.Sprintf( "ERROR: Database reports '%s'", err.Error() ) )
		encodeHealthCheckResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	encodeHealthCheckResponse(w, http.StatusOK, "")
}

//
// end of file
//
