package start

import (
	"context"
	"fmt"
	"griffin-dao/ent"
	"griffin-dao/ent/crypto_prc_source"
	"log"
)

func CreateCryptoPrcSource(ctx context.Context, client *ent.Client) (*ent.CRYPTO_PRC_SOURCE, error) {
	cryptoPrcSourceT, err := client.CRYPTO_PRC_SOURCE.
		Create().
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed creating table CRYPTO_PRC_SOURCE: %w", err)
	}
	log.Printf("table %v was created\n", cryptoPrcSourceT)
	return cryptoPrcSourceT, nil
}

func CreateCryptoCurrency(ctx context.Context, client *ent.Client) (*ent.CRYPTO_CURRENCY, error) {
	cryptoCurrencyT, err := client.CRYPTO_CURRENCY.
		Create().
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed creating table CRYPTO_CURRENCY: %w", err)
	}
	log.Printf("table %v was created\n", cryptoCurrencyT)
	return cryptoCurrencyT, nil
}

func InsertCryptoSource(exc string, excCode int, ctx context.Context, client *ent.Client) {
	exchange, err := client.CRYPTO_PRC_SOURCE.
		Create().
		SetID(excCode).
		SetName(exc).
		Save(ctx)
	if err != nil {
		fmt.Errorf("failed creating entity: %w", err)
	}
	log.Printf("Entity %v was created\n", exchange)
}

func InsertCryptoCurrency(fromExc int, currency string, ctx context.Context, source *ent.CRYPTO_CURRENCY) {
	_, err := source.QueryCryptoPrcSource().
		Where(crypto_prc_source.ID(fromExc)).
		Only(ctx)
	if err != nil {
		fmt.Errorf("Exchange with code %s does not exist", fromExc)
	}

}
