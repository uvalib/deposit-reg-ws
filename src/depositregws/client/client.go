package client

import (
    "time"
    "fmt"
    "github.com/parnurzeal/gorequest"
    "net/http"
    "depositregws/api"
    "encoding/json"
)

func HealthCheck( endpoint string ) int {

    url := fmt.Sprintf( "%s/healthcheck", endpoint )
    //fmt.Printf( "%s\n", url )

    resp, _, errs := gorequest.New( ).
       SetDebug( false ).
       Get( url ).
       Timeout( time.Duration( 5 ) * time.Second ).
       End( )

    if errs != nil {
        return http.StatusInternalServerError
    }

    defer resp.Body.Close( )

    return resp.StatusCode
}

func Get( endpoint string, id string, token string ) ( int, [] * api.Registration ) {

    url := fmt.Sprintf( "%s/%s?auth=%s", endpoint, id, token )
    //fmt.Printf( "%s\n", url )

    resp, body, errs := gorequest.New( ).
       SetDebug( false ).
       Get( url  ).
       Timeout( time.Duration( 5 ) * time.Second ).
       End( )

    if errs != nil {
       return http.StatusInternalServerError, nil
    }

    defer resp.Body.Close( )

    r := api.StandardResponse{ }
    err := json.Unmarshal( []byte( body ), &r )
    if err != nil {
        return http.StatusInternalServerError, nil
    }

    return resp.StatusCode, r.Details
}