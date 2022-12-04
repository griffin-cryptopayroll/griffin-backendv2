package main

import (
	"griffin-dao/api/v0"
	"griffin-dao/dbstartup"
	"log"
	"net/http"
	"os"
	"time"
)

// @title           Griffin Web Server API Documentation
// @version         Document 1.0
// @description     Griffin webserver that serves, employee .

// @contact.name   Sang Il Bae
// @contact.email  baesangil0906@gmail.com

// @host      localhost:8080/api
// @BasePath  /api/v0
func main() {
	griffinServer := v0.WebServerStartUp()
	dbstartup.ExecStartUp(griffinServer.Database)

	gbe := griffinServer.
		InitializeApiCommon(). // ping, version, login
		InitializeApiV0()

	griffinPay := &http.Server{
		Addr:           ":" + os.Getenv("PORT"),
		Handler:        gbe.Conn,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Fatal(griffinPay.ListenAndServe())
}
