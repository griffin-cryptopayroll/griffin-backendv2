package main

import (
	"flag"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"griffin-dao/api/server"
	"griffin-dao/startup"
	"griffin-dao/util"
	"log"
	"net/http"
	"os"
	"time"
)

func init() {
	option := flag.String("env", "local", "local state or dev state")
	flag.Parse()

	var fileName string
	switch {
	case *option == "local":
		util.PrintPurpleWarning("ENVIRONMENT: serve in local env")
		fileName = "./.env.local"
	case *option == "deploy":
		util.PrintPurpleWarning("ENVIRONMENT: serve in deploy env")
		fileName = "./.env.serve"
		gin.SetMode(gin.ReleaseMode)
	default:
		log.Panicln("state your correct dev state")
	}
	godotenv.Load(fileName)
}

// @title           Griffin Web Server API Documentation
// @version         Document 1.0
// @description     Griffin webserver that serves, employee .

// @contact.name   Sang Il Bae
// @contact.email  baesangil0906@gmail.com

// @host      localhost:10433
// @BasePath  /
func main() {
	griffinServer := server.WebServerStartUp()
	startup.ExecStartUp(griffinServer.Database)

	gbe := griffinServer.
		InitializeApiCommon(). // ping, version, login
		InitializeApiV0().
		InitializeLoginV1().
		SessionUsage().TokenUsage()

	griffinPay := &http.Server{
		Addr:           os.Getenv("HOSTNAME") + ":" + os.Getenv("PORT"),
		Handler:        gbe.Conn,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Fatal(griffinPay.ListenAndServe())
}
