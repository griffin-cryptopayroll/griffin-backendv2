package gcrud

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

func CreateEmployee(entity EmployeeJson, ctx context.Context, client *ent.Client) {
	obj, err := client.EMPLOYEE.
		Create().
		SetID(entity.ID).
		SetGid(entity.GriffinID).
		SetEmployerGid(entity.EmployerGriffinID).
		SetLastName(entity.LastName).
		SetFirstName(entity.FirstName).
		SetPosition(entity.Position).
		SetWallet(entity.Wallet).
		SetPayroll(entity.Payroll).
		SetCurrency(entity.Currency).
		SetPayday(entity.PayDay).
		SetEmploy(entity.EmployType).
		SetEmail(entity.Email).
		SetCreatedAt(entity.CreatedAt).
		SetCreatedBy(entity.CreatedBy).
		Save(ctx)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Employee created", obj)
}

func CreateEmployerUserInfo(entity EmployerJson, ctx context.Context, client *ent.Client) {
	obj, err := client.EMPLOYER_USER_INFO.
		Create().
		SetID(entity.ID).
		SetUsername(entity.Username).
		SetPassword(entity.Password).
		SetGid(entity.GriffinID).
		SetCorpName(entity.CorpName).
		SetCorpEmail(entity.CorpEmail).
		SetWallet(entity.Wallet).
		SetCreatedAt(entity.CreatedAt).
		SetCreatedBy(entity.CreatedBy).
		SetUpdatedAt(entity.UpdatedAt).
		SetUpdatedBy(entity.UpdatedBy).
		Save(ctx)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Employer User Info created", obj)
}

func CreatePaymentHistory(entity PaymentHistory, ctx context.Context, client *ent.Client) {
	obj, err := client.PAYMENT_HISTORY.
		Create().
		SetID(entity.ID).
		SetEmployeeGid(entity.EmployeeGriffinID).
		SetCreatedAt(entity.CreatedAt).
		SetCreatedBy(entity.CreatedBy).
		Save(ctx)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Payment_History created", obj)
}
