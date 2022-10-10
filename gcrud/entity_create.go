package gcrud

import (
	"context"
	"griffin-dao/ent"
	"griffin-dao/service"
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

func CreateEmployType(permaBool string, contractMonth int, ctx context.Context, client *ent.Client) {
	obj, err := client.EMPLOY_TYPE.
		Create().
		SetIsPermanent(permaBool).
		SetContractPeriod(contractMonth).
		Save(ctx)
	if err != nil {
		service.PrintRedError(err)
		return
	}
	service.PrintGreenStatus("Employ_Type created", obj)
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
		SetWorkStart(entity.WorkStart).
		SetCreatedAt(entity.CreatedAt).
		SetCreatedBy(entity.CreatedBy).
		SetUpdatedAt(entity.UpdatedAt).
		SetUpdatedBy(entity.UpdatedBy).
		Save(ctx)
	if err != nil {
		service.PrintRedError(err)
		return
	}
	service.PrintGreenStatus("Employee created", obj)
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
		service.PrintRedError(err)
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
