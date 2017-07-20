package handlers

import (
	"depositregws/dao"
	"net/http"
)

func HealthCheck(w http.ResponseWriter, r *http.Request) {

	err := dao.Database.Check()
	if err != nil {
		EncodeHealthCheckResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	EncodeHealthCheckResponse(w, http.StatusOK, "")
}
