package main

import (
	"fmt"
	"griffin-dao/api"
	"griffin-dao/gcrud"
	"net/http"
	"os"
	"time"
)

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

	gbe.AddEmployer()

	griffinPay := &http.Server{
		Addr:           ":" + os.Getenv("PORT"),
		Handler:        gbe.Conn,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	griffinPay.ListenAndServe()
	// gcrud.CreateCryptoCurrency(1, gcrud.BINANCE_CODE, "MATIC", context.Background(), griffin.Conn)
	// gcrud.CreateCryptoSource(gcrud.BINANCE, gcrud.BINANCE_CODE, context.Background(), griffin.Conn)
	// gcrud.LinkCryptoCurrencywSource("MATIC", gcrud.BINANCE_CODE, context.Background(), griffin.Conn)

	// This is the database update portion -
	//gcrud.CreateCryptoCurrency(2, gcrud.BINANCE_CODE, "ETHUSDT", context.Background(), griffin.Conn)
	//gcrud.LinkCryptoCurrencywSource("ETHUSDT", gcrud.BINANCE_CODE, context.Background(), griffin.Conn)

}
