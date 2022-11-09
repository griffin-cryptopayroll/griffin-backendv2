package dbstartup

import (
	"context"
	"griffin-dao/gcrud"
)

func sourceStartUp(db gcrud.GriffinWeb2Conn) {
	gcrud.CreateCryptoSource(gcrud.BINANCE, gcrud.BINANCE_CODE, context.Background(), db.Conn)
	gcrud.CreateCryptoSource(gcrud.FTX, gcrud.FTX_CODE, context.Background(), db.Conn)
	gcrud.CreateCryptoSource(gcrud.UPBIT, gcrud.UPBIT_CODE, context.Background(), db.Conn)
	gcrud.CreateCryptoSource(gcrud.MEXC, gcrud.MEXC_CODE, context.Background(), db.Conn)
}

func currencyStartUp(db gcrud.GriffinWeb2Conn) {
	gcrud.CreateCryptoCurrency(gcrud.BINANCE_CODE, "MATICUSDT", context.Background(), db.Conn)
	gcrud.CreateCryptoCurrency(gcrud.BINANCE_CODE, "ETHUSDT", context.Background(), db.Conn)
	gcrud.CreateCryptoCurrency(gcrud.BINANCE_CODE, "USDCUSDT", context.Background(), db.Conn)
}
