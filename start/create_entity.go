package start

import (
	"context"
	"log"
	"time"

	"griffin-dao/ent"
)

func CreateCryptoCurrency(id, exchCode int, ticker string, ctx context.Context, client *ent.Client) {
	// id, ticker, source
	obj, err := client.CRYPTO_CURRENCY.
		Create().
		SetID(id).
		SetTicker(ticker).
		SetSource(exchCode).
		Save(ctx)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Crypto_Currency created", obj)
}

func CreateCryptoSource(exch string, exchCode int, ctx context.Context, client *ent.Client) {
	obj, err := client.CRYPTO_PRC_SOURCE.
		Create().
		SetID(BINANCE_CODE).
		SetName(BINANCE).
		Save(ctx)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Crypto_Source created", obj)
}

func CreateEmployType(id int, permaBool string, contractStart time.Time, contractMonth int, ctx context.Context, client *ent.Client) {
	obj, err := client.EMPLOY_TYPE.
		Create().
		SetID(id).
		SetIsPermanent(permaBool).
		SetContractStart(contractStart).
		SetContractPeriod(contractMonth).
		Save(ctx)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Employ_Type created", obj)
}

func CreateEmployee() {

}

func CreateEmployerUserInfo() {

}

func CreatePaymentHistory(id int, employeeGriffinId string, createdAt time.Time, createdBy string, ctx context.Context, client *ent.Client) {
	obj, err := client.PAYMENT_HISTORY.
		Create().
		SetID(id).
		SetEmployeeGid(employeeGriffinId).
		SetCreatedAt(createdAt).
		SetCreatedBy(createdBy).
		Save(ctx)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Payment_History created", obj)
}
