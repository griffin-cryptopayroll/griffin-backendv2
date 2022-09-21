package main

import (
	"context"
	"fmt"
	"griffin-dao/api"
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

func main() {
	var griffin gcrud.GriffinWeb2Conn
	var gbe api.GriffinWS
	if griffinConn, err := gcrud.NewDao(); err == nil {
		griffin = griffinConn.Conn()
	}
	fmt.Println(griffin.Conn.CRYPTO_PRC_SOURCE)

	gbe = gbe.
		StartService().
		PingTest().
		Version()

	gbe.
		AddEmployer().DeleteEmployer().
		AddEmployee().DeleteEmployee()

	griffinPay := &http.Server{
		Addr:           ":" + os.Getenv("PORT"),
		Handler:        gbe.Conn,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	griffinPay.ListenAndServe()
	// gcrud.CreateCryptoCurrency(gcrud.BINANCE_CODE, "MATIC", context.Background(), griffin.Conn)
	// gcrud.CreateCryptoSource(gcrud.BINANCE, gcrud.BINANCE_CODE, context.Background(), griffin.Conn)
	// gcrud.LinkCryptoCurrencywSource("MATIC", gcrud.BINANCE_CODE, context.Background(), griffin.Conn)

	// This is the database update portion -
	//gcrud.CreateCryptoCurrency(gcrud.BINANCE_CODE, "ETHUSDT", context.Background(), griffin.Conn)
	//gcrud.LinkCryptoCurrencywSource("ETHUSDT", gcrud.BINANCE_CODE, context.Background(), griffin.Conn)

}
