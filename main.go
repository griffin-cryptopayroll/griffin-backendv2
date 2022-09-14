package main

import (
	"context"
	"fmt"
	"griffin-dao/start"
)

func main() {
	var griffin start.GriffinWeb2Conn
	if griffinConn, err := start.NewDao(); err == nil {
		griffin = griffinConn.Conn()
	}
	fmt.Println(griffin.Conn.CRYPTO_PRC_SOURCE)

	// start.CreateCryptoCurrency(1, start.BINANCE_CODE, "MATIC", context.Background(), griffin.Conn)
	// start.CreateCryptoSource(start.BINANCE, start.BINANCE_CODE, context.Background(), griffin.Conn)
	// start.UpdateCryptoCurrencywSource("MATIC", start.BINANCE_CODE, context.Background(), griffin.Conn)

	// This is the database update portion -
	start.CreateCryptoCurrency(2, start.BINANCE_CODE, "ETHUSDT", context.Background(), griffin.Conn)
	start.UpdateCryptoCurrencywSource("ETHUSDT", start.BINANCE_CODE, context.Background(), griffin.Conn)
}
