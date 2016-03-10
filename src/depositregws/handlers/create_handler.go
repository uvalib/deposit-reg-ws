package handlers

import (
    "log"
    "fmt"
    "net/http"
    "encoding/json"
//    "github.com/gorilla/mux"
    "depositregws/authtoken"
    "depositregws/config"
    "depositregws/dao"
    "depositregws/api"
)

func RegistrationCreate( w http.ResponseWriter, r *http.Request ) {

    token := r.URL.Query( ).Get( "auth" )

    // parameters OK ?
    if NotEmpty( token ) == false {
        EncodeStandardResponse( w, http.StatusBadRequest, http.StatusText( http.StatusBadRequest ), nil )
        return
    }

    // validate the token
    if authtoken.Validate( config.Configuration.AuthTokenEndpoint, token ) == false {
        EncodeStandardResponse( w, http.StatusForbidden, http.StatusText( http.StatusForbidden ), nil )
        return
    }

    decoder := json.NewDecoder( r.Body )
    reg := api.Registration{ }

    if err := decoder.Decode( &reg ); err != nil {
        EncodeStandardResponse( w, http.StatusBadRequest, http.StatusText( http.StatusBadRequest ), nil )
        return
    }

    // get the request details
    rg, err := dao.Database.Create( reg )
    if err != nil {
        log.Println( err )
        EncodeStandardResponse( w, http.StatusInternalServerError,
            fmt.Sprintf( "%s (%s)", http.StatusText( http.StatusInternalServerError ), err ),
            nil )
        return
    }

    results := make([ ] * api.Registration, 0 )
    results = append( results, rg )
    EncodeStandardResponse( w, http.StatusOK, http.StatusText( http.StatusOK ), results )
}