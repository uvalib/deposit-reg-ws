package handlers

import (
   "depositregws/api"
   "depositregws/dao"
   "depositregws/logger"
   "fmt"
   "net/http"
)

//
// OptionsGet -- get options request handler
//
func OptionsGet(w http.ResponseWriter, r *http.Request) {

   //token := r.URL.Query( ).Get( "auth" )

   // parameters OK ?
   //if notEmpty( token ) == false {
   //    encodeOptionsResponse( w, http.StatusBadRequest, http.StatusText( http.StatusBadRequest ), nil )
   //    return
   //}

   // validate the token
   //if authtoken.Validate( config.Configuration.AuthTokenEndpoint, token ) == false {
   //    encodeOptionsResponse( w, http.StatusForbidden, http.StatusText( http.StatusForbidden ), nil )
   //    return
   //}

   departments, err := dao.DB.GetFieldSet("department")
   if err != nil {
      logger.Log(fmt.Sprintf("ERROR: %s\n", err.Error()))
      status := http.StatusInternalServerError
      encodeOptionsResponse(w, status,
         fmt.Sprintf("%s (%s)", http.StatusText(status), err),
         nil)
      return
   }

   degrees, err := dao.DB.GetFieldSet("degree")
   if err != nil {
      logger.Log(fmt.Sprintf("ERROR: %s\n", err.Error()))
      status := http.StatusInternalServerError
      encodeOptionsResponse(w, status,
         fmt.Sprintf("%s (%s)", http.StatusText(status), err),
         nil)
      return
   }

   options := api.Options{Department: departments, Degree: degrees}

   status := http.StatusOK
   encodeOptionsResponse(w, status, http.StatusText(status), &options)
}

//
// end of file
//
