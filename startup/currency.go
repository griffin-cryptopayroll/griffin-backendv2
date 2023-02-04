package startup

import (
	"context"
	"errors"
	"fmt"
	"griffin-dao/dao"
	"griffin-dao/util"
)

var ctx = context.Background()

func sourceStartUp(db dao.GriffinWeb2Conn) error {
	err1 := dao.CreateCryptoSource(dao.BINANCE, dao.BINANCE_CODE, ctx, db.Conn)
	err2 := dao.CreateCryptoSource(dao.FTX, dao.FTX_CODE, ctx, db.Conn)
	err3 := dao.CreateCryptoSource(dao.UPBIT, dao.UPBIT_CODE, ctx, db.Conn)
	err4 := dao.CreateCryptoSource(dao.MEXC, dao.MEXC_CODE, ctx, db.Conn)
	if err1 != nil || err2 != nil || err3 != nil || err4 != nil {
		util.PrintRedError(
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

func currencyStartUp(db dao.GriffinWeb2Conn) error {
	err1 := dao.CreateCryptoCurrency(dao.BINANCE_CODE, "MATIC", ctx, db.Conn)
	err2 := dao.CreateCryptoCurrency(dao.BINANCE_CODE, "ETH", ctx, db.Conn)
	err3 := dao.CreateCryptoCurrency(dao.BINANCE_CODE, "USDC", ctx, db.Conn)
	if err1 != nil || err2 != nil || err3 != nil {
		util.PrintRedError(
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
