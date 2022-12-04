package dbstartup

import (
	"context"
	"griffin-dao/gcrud"
)

var ctx = context.Background()

func sourceStartUp(db gcrud.GriffinWeb2Conn) {
	gcrud.CreateCryptoSource(gcrud.BINANCE, gcrud.BINANCE_CODE, ctx, db.Conn)
	gcrud.CreateCryptoSource(gcrud.FTX, gcrud.FTX_CODE, ctx, db.Conn)
	gcrud.CreateCryptoSource(gcrud.UPBIT, gcrud.UPBIT_CODE, ctx, db.Conn)
	gcrud.CreateCryptoSource(gcrud.MEXC, gcrud.MEXC_CODE, ctx, db.Conn)
}

func currencyStartUp(db gcrud.GriffinWeb2Conn) {
	gcrud.CreateCryptoCurrency(gcrud.BINANCE_CODE, "MATICUSDT", ctx, db.Conn)
	gcrud.CreateCryptoCurrency(gcrud.BINANCE_CODE, "ETHUSDT", ctx, db.Conn)
	gcrud.CreateCryptoCurrency(gcrud.BINANCE_CODE, "USDCUSDT", ctx, db.Conn)
}
