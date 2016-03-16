package handlers

import (
    "log"
    "fmt"
    "net/http"
    //"depositregws/authtoken"
    //"depositregws/config"
    "depositregws/api"
    "depositregws/dao"
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

    // get the request details
    departments, err := dao.Database.GetFieldSet( "department" )
    if err != nil {
        log.Println( err )
        status := http.StatusInternalServerError
        EncodeOptionsResponse( w, status,
            fmt.Sprintf( "%s (%s)", http.StatusText( status ), err ),
            nil )
        return
    }

    degrees := []string{"Graduate","Masters","Ph.D"}
    options := api.Options{ School: departments, Degree: degrees }

    status := http.StatusOK
    EncodeOptionsResponse( w, status, http.StatusText( status ), &options )
}