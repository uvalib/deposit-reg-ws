package handlers

import (
   "depositregws/authtoken"
   "depositregws/config"
   "depositregws/dao"
   "depositregws/logger"
   "fmt"
   "net/http"
)

//
// RegistrationSearch -- the search registration handler
//
func RegistrationSearch(w http.ResponseWriter, r *http.Request) {

   token := r.URL.Query().Get("auth")
   id := r.URL.Query().Get("later")

   // parameters OK ?
   if notEmpty(token) == false || notEmpty(id) == false {
      status := http.StatusBadRequest
      encodeStandardResponse(w, status, http.StatusText(status), nil)
      return
   }

   // validate the token
   if authtoken.Validate(config.Configuration.AuthTokenEndpoint, token, config.Configuration.Timeout) == false {
      status := http.StatusForbidden
      encodeStandardResponse(w, status, http.StatusText(status), nil)
      return
   }

   // get the request details
   reqs, err := dao.DB.SearchDepositRequest(id)
   if err != nil {
      logger.Log(fmt.Sprintf("ERROR: %s\n", err.Error()))
      status := http.StatusInternalServerError
      encodeStandardResponse(w, status,
         fmt.Sprintf("%s (%s)", http.StatusText(status), err),
         nil)
      return
   }

   if reqs == nil || len(reqs) == 0 {
      status := http.StatusNotFound
      encodeStandardResponse(w, status, http.StatusText(status), nil)
      return
   }

   status := http.StatusOK
   encodeStandardResponse(w, status, http.StatusText(status), reqs)
}

//
// end of file
//
