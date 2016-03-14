package handlers

import (
    "net/http"
    //"depositregws/authtoken"
    //"depositregws/config"
    "depositregws/api"
)

func OptionsGet( w http.ResponseWriter, r *http.Request ) {

    //token := r.URL.Query( ).Get( "auth" )

    // parameters OK ?
    //if NotEmpty( token ) == false {
    //    EncodeOptionsResponse( w, http.StatusBadRequest, http.StatusText( http.StatusBadRequest ), nil )
    //    return
    //}

    // validate the token
    //if authtoken.Validate( config.Configuration.AuthTokenEndpoint, token ) == false {
    //    EncodeOptionsResponse( w, http.StatusForbidden, http.StatusText( http.StatusForbidden ), nil )
    //    return
    //}

    // get possible registration options
    schools := []string{"Business","Engineering","Health Sciences"}
    degrees := []string{"Graduate","Masters","Ph.D"}
    options := api.Options{ School: schools, Degree: degrees }

    status, msg := StatusHelper( http.StatusOK )
    EncodeOptionsResponse( w, status, msg, &options )
}