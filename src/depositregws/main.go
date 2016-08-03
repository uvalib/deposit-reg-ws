package main

import (
    "fmt"
    "log"
    "net/http"
    "depositregws/config"
    "depositregws/dao"
    "depositregws/handlers"
    "depositregws/logger"
)

func main( ) {

    logger.Log( fmt.Sprintf( "===> version: '%s' <===", handlers.Version( ) ) )

    // access the database
    connectStr := fmt.Sprintf( "%s:%s@tcp(%s)/%s?allowOldPasswords=1", config.Configuration.DbUser,
        config.Configuration.DbPassphrase, config.Configuration.DbHost, config.Configuration.DbName )

    err := dao.NewDB( connectStr )
    if err != nil {
        log.Fatal( err )
    }

	// setup router and serve...
    router := NewRouter( )
    log.Fatal( http.ListenAndServe( fmt.Sprintf( ":%s", config.Configuration.ServicePort ), router ) )
}

