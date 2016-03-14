package handlers

import (
    "log"
    "fmt"
    "strings"
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

    // create results list
    results := make([ ] * api.Registration, 0 )

    // split the user list of appropriate
    users := strings.Split( reg.For, "," )

    for _, u := range users {

        reg.For = strings.TrimSpace( u )
        rg, err := dao.Database.Create( reg )
        if err != nil {
            log.Println(err)
            EncodeStandardResponse(w, http.StatusInternalServerError,
                fmt.Sprintf("%s (%s)", http.StatusText(http.StatusInternalServerError), err),
                nil)
            return
        }

        results = append(results, rg)
    }

    EncodeStandardResponse( w, http.StatusOK, http.StatusText( http.StatusOK ), results )
}

func RegistrationCreateOptions( w http.ResponseWriter, r *http.Request ) {
    w.Header( ).Add( "Access-Control-Allow-Methods", "POST" )
    EncodeStandardResponse( w, http.StatusOK, http.StatusText( http.StatusOK ), nil )
}