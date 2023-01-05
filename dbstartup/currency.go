package dbstartup

import (
	"context"
	"errors"
	"fmt"
	"griffin-dao/gcrud"
	"griffin-dao/service"
)

var ctx = context.Background()

func sourceStartUp(db gcrud.GriffinWeb2Conn) error {
	err1 := gcrud.CreateCryptoSource(gcrud.BINANCE, gcrud.BINANCE_CODE, ctx, db.Conn)
	err2 := gcrud.CreateCryptoSource(gcrud.FTX, gcrud.FTX_CODE, ctx, db.Conn)
	err3 := gcrud.CreateCryptoSource(gcrud.UPBIT, gcrud.UPBIT_CODE, ctx, db.Conn)
	err4 := gcrud.CreateCryptoSource(gcrud.MEXC, gcrud.MEXC_CODE, ctx, db.Conn)
	if err1 != nil || err2 != nil || err3 != nil || err4 != nil {
		service.PrintRedError(
			fmt.Sprintf(
				"%s | %s | %s | %s",
				err1.Error(),
				err2.Error(),
				err3.Error(),
				err4.Error(),
			),
		)
		return errors.New("startup::create crypto_source error")
	}
	return nil
}

func currencyStartUp(db gcrud.GriffinWeb2Conn) error {
	err1 := gcrud.CreateCryptoCurrency(gcrud.BINANCE_CODE, "MATIC", ctx, db.Conn)
	err2 := gcrud.CreateCryptoCurrency(gcrud.BINANCE_CODE, "ETH", ctx, db.Conn)
	err3 := gcrud.CreateCryptoCurrency(gcrud.BINANCE_CODE, "USDC", ctx, db.Conn)
	if err1 != nil || err2 != nil || err3 != nil {
		service.PrintRedError(
			fmt.Sprintf(
				"%s | %s | %s | %s",
				err1.Error(),
				err2.Error(),
				err3.Error(),
			),
		)
		return errors.New("startup::create currency error")
	}
	return nil
}
