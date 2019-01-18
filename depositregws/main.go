package main

import (
	"fmt"
	"github.com/uvalib/deposit-reg-ws/depositregws/config"
	"github.com/uvalib/deposit-reg-ws/depositregws/dao"
	"github.com/uvalib/deposit-reg-ws/depositregws/handlers"
	"github.com/uvalib/deposit-reg-ws/depositregws/logger"
	"log"
	"net/http"
	"time"
)

func main() {

	logger.Log(fmt.Sprintf("===> version: '%s' <===", handlers.Version()))

	// create the storage singleton
	err := dao.NewDatastore()
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
