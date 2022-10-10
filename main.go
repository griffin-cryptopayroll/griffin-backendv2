package main

import (
	"context"
	"fmt"
	"griffin-dao/api/v1"
	"griffin-dao/gcrud"
	"net/http"
	"os"
	"time"
)

func currencyStartUp(db gcrud.GriffinWeb2Conn) {
	gcrud.CreateCryptoSource(gcrud.BINANCE, gcrud.BINANCE_CODE, context.Background(), db.Conn)
	gcrud.CreateCryptoSource(gcrud.FTX, gcrud.FTX_CODE, context.Background(), db.Conn)
	gcrud.CreateCryptoSource(gcrud.UPBIT, gcrud.UPBIT_CODE, context.Background(), db.Conn)
	gcrud.CreateCryptoSource(gcrud.MEXC, gcrud.MEXC_CODE, context.Background(), db.Conn)

	gcrud.CreateCryptoCurrency(gcrud.BINANCE_CODE, "MATICUSDT", context.Background(), db.Conn)
	gcrud.CreateCryptoCurrency(gcrud.BINANCE_CODE, "ETHUSDT", context.Background(), db.Conn)
	gcrud.CreateCryptoCurrency(gcrud.BINANCE_CODE, "USCDUSDT", context.Background(), db.Conn)
}

// @title           Griffin Web Server API Documentation
// @version         Document 1.0
// @description     Griffin webserver that serves, employee .

// @contact.name   Sang Il Bae
// @contact.email  baesangil0906@gmail.com

// @host      localhost:8080
// @BasePath  /
func main() {
	var griffin gcrud.GriffinWeb2Conn
	var gbe v1.GriffinWS
	if griffinConn, err := gcrud.NewDao(); err == nil {
		griffin = griffinConn.Conn()
	}
	fmt.Println(griffin.Conn.CRYPTO_PRC_SOURCE)

	gbe = gbe.
		StartService().
		PingTest().
		Version()

	// CRUD Op.
	gbe.
		// Employ Type CRUD
		AddEmployType().DeleteEmployType().GetEmployType().
		// Employer CRUD
		AddEmployer().DeleteEmployer().GetEmployer().
		// Employee CRUD
		AddEmployee().DeleteEmployee().GetEmployeeSingle().GetEmployeeMulti().
		// Price points and payment currency
		GetPrice()

	griffinPay := &http.Server{
		Addr:           ":" + os.Getenv("PORT"),
		Handler:        gbe.Conn,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	griffinPay.ListenAndServe()
}
