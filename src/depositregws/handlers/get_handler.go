package handlers

import (
    "log"
    "fmt"
    "net/http"
    "github.com/gorilla/mux"
    "depositregws/authtoken"
    "depositregws/config"
    "depositregws/dao"
)

func RegistrationGet( w http.ResponseWriter, r *http.Request ) {

    vars := mux.Vars( r )
    id := vars[ "id" ]
    token := r.URL.Query( ).Get( "auth" )

    // parameters OK ?
    if NotEmpty( id ) == false || NotEmpty( token ) == false {
        EncodeStandardResponse( w, http.StatusBadRequest, http.StatusText( http.StatusBadRequest ), nil )
        return
    }

    // validate the token
    if authtoken.Validate( config.Configuration.AuthTokenEndpoint, token ) == false {
        EncodeStandardResponse( w, http.StatusForbidden, http.StatusText( http.StatusForbidden ), nil )
        return
    }

    // get the request details
    reqs, err := dao.Database.Get( id )
    if err != nil {
        log.Println( err )
        EncodeStandardResponse( w, http.StatusInternalServerError,
            fmt.Sprintf( "%s (%s)", http.StatusText( http.StatusInternalServerError ), err ),
            nil )
        return
    }

    if reqs == nil || len( reqs ) == 0 {
        EncodeStandardResponse( w, http.StatusNotFound, http.StatusText( http.StatusNotFound ), nil )
        return
    }

    EncodeStandardResponse( w, http.StatusOK, http.StatusText( http.StatusOK ), reqs )
}