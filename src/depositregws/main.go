package main

import (
    "fmt"
    "log"
    "net/http"
    "depositregws/config"
    "depositregws/dao"
)

type Env struct {
    db dao.Datastore
}

var env = &Env{ }

func main( ) {

    // access the database

    connectStr := fmt.Sprintf( "%s:%s@tcp(%s)/%s", config.Configuration.DbUser,
        config.Configuration.DbPassphrase, config.Configuration.DbHost, config.Configuration.DbName )

    db, err := dao.NewDB( connectStr )
    if err != nil {
        log.Fatal( err )
    }

    env.db = db

	// setup router and serve...
    router := NewRouter( )
    log.Fatal( http.ListenAndServe( fmt.Sprintf( ":%s", config.Configuration.ServicePort ), router ) )
}

