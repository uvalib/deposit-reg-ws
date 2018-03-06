package main

import (
	"depositregws/config"
	"depositregws/dao"
	"depositregws/handlers"
	"depositregws/logger"
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {

	logger.Log(fmt.Sprintf("===> version: '%s' <===", handlers.Version()))

	// access the database
	err := dao.NewDB(
		config.Configuration.DbHost,
		config.Configuration.DbName,
		config.Configuration.DbUser,
		config.Configuration.DbPassphrase,
		config.Configuration.DbTimeout)
	if err != nil {
		log.Fatal(err)
	}

	// setup router and server...
	serviceTimeout := 15 * time.Second
	router := NewRouter()
	server := &http.Server{
		Addr:         fmt.Sprintf(":%s", config.Configuration.ServicePort),
		Handler:      router,
		ReadTimeout:  serviceTimeout,
		WriteTimeout: serviceTimeout,
	}
	log.Fatal(server.ListenAndServe())
}

//
// end of file
//
