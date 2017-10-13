package handlers

import (
   "depositregws/dao"
   "net/http"
)

//
// HealthCheck -- do the healthcheck
//
func HealthCheck(w http.ResponseWriter, r *http.Request) {

   err := dao.DB.CheckDB()
   if err != nil {
      encodeHealthCheckResponse(w, http.StatusInternalServerError, err.Error())
      return
   }
   encodeHealthCheckResponse(w, http.StatusOK, "")
}

//
// end of file
//
