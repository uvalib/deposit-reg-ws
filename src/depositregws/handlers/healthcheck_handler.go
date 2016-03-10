package handlers

import (
    "net/http"
    "depositregws/dao"
)

func HealthCheck( w http.ResponseWriter, r *http.Request ) {

    err := dao.Database.Check( )
    if err != nil {
        EncodeHealthCheckResponse( w, http.StatusInternalServerError, err.Error( ) )
        return
    }
    EncodeHealthCheckResponse( w, http.StatusOK, "" )
}