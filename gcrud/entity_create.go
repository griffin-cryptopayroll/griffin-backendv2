package gcrud

import (
	"errors"
	"griffin-dao/ent"
	"griffin-dao/service"

	"context"

	"github.com/google/uuid"
)

func CreateCryptoCurrency(exchCode int, ticker string, ctx context.Context, client *ent.Client) {
	// id, ticker, source
	obj, err := client.CRYPTO_CURRENCY.
		Create().
		SetTicker(ticker).
		SetSource(exchCode).
		Save(ctx)
	if err != nil {
		service.PrintRedError(err)
		return
	}
	service.PrintGreenStatus("Crypto_Currency created", obj)
}

func CreateCryptoSource(exch string, exchCode int, ctx context.Context, client *ent.Client) {
	obj, err := client.CRYPTO_PRC_SOURCE.
		Create().
		SetID(BINANCE_CODE).
		SetName(BINANCE).
		Save(ctx)
	if err != nil {
		service.PrintRedError(err)
		return
	}
	service.PrintGreenStatus("Crypto_Source created", obj)
}

func CreateEmployType(permaBool, payFreq string, ctx context.Context, client *ent.Client) error {
	obj, err := client.EMPLOY_TYPE.
		Create().
		SetIsPermanent(permaBool).
		SetPayFreq(payFreq).
		Save(ctx)
	if recover() != nil || err != nil {
		service.PrintRedError(err)
		return err
	}
	service.PrintGreenStatus("Employ_Type created", obj)
	return nil
}

func CreateEmployee(entity EmployeeJson, ctx context.Context, client *ent.Client) error {
	gidNew := uuid.New()
	obj, err := client.EMPLOYEE.
		Create().
		SetID(entity.ID).
		SetGid(gidNew.String()).
		SetEmployerGid(entity.EmployerGriffinID). // uuid
		SetName(entity.Name).
		SetPosition(entity.Position).
		SetWallet(entity.Wallet).
		SetPayroll(entity.Payroll).
		SetCurrency(entity.Currency).
		SetPayday(entity.PayDay).
		SetEmploy(entity.EmployType).
		SetEmail(entity.Email).
		SetWorkStart(entity.WorkStart).
		SetWorkEnds(entity.WorkEnd).
		SetCreatedAt(entity.CreatedAt).
		SetCreatedBy(entity.CreatedBy).
		SetUpdatedAt(entity.UpdatedAt).
		SetUpdatedBy(entity.UpdatedBy).
		Save(ctx)
	if recover() != nil || err != nil {
		service.PrintRedError(err)
		return errors.New(DATABASE_CREATE_FAIL)
	}
	service.PrintGreenStatus("Employee created", obj)
	return nil
}

func CreateEmployerUserInfo(entity EmployerJson, ctx context.Context, client *ent.Client) {
	gidNew := uuid.New()
	obj, err := client.EMPLOYER_USER_INFO.
		Create().
		SetID(entity.ID).
		SetUsername(entity.Username).
		SetPassword(entity.Password).
		SetGid(gidNew.String()).
		SetCorpName(entity.CorpName).
		SetCorpEmail(entity.CorpEmail).
		SetWallet(entity.Wallet).
		SetCreatedAt(entity.CreatedAt).
		SetCreatedBy(entity.CreatedBy).
		SetUpdatedAt(entity.UpdatedAt).
		SetUpdatedBy(entity.UpdatedBy).
		Save(ctx)
	if err != nil {
		service.PrintRedError(err)
		return
	}
	service.PrintGreenStatus("Employer_User_Info created", obj)
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
		service.PrintRedError(err)
		return
	}
	service.PrintGreenStatus("Payment_History created", obj)
}
