package start

import (
	"context"
	"fmt"
	"griffin-dao/ent"
	"griffin-dao/ent/crypto_currency"
	"griffin-dao/ent/crypto_prc_source"
)

func UpdateCryptoCurrencywSource(curr string, exchCode int, ctx context.Context, client *ent.Client) {
	obj, err := client.CRYPTO_CURRENCY.
		Query().
		Where(crypto_currency.Ticker(curr)).
		Only(ctx)
	if err != nil {
		fmt.Println("object finding", err)
		return
	}
	source, err := client.CRYPTO_PRC_SOURCE.
		Query().
		Where(crypto_prc_source.ID(exchCode)).
		Only(ctx)
	if err != nil {
		fmt.Println("source finding", err)
		return
	}
	source.Update().AddPriceOf(obj)
}
