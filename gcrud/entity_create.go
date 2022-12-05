package gcrud

import (
	"context"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"griffin-dao/ent"
	"griffin-dao/ent/crypto_currency"
	"griffin-dao/ent/employ_type"
	"griffin-dao/ent/employer"
	"griffin-dao/service"
	"time"
)

func CreateCryptoCurrency(exchCode int, ticker string, ctx context.Context, client *ent.Client) error {
	// id, ticker, source
	obj, err := client.CRYPTO_CURRENCY.
		Create().
		SetTicker(ticker).
		SetSourceID(exchCode).
		Save(ctx)
	if recover() != nil || err != nil {
		service.PrintRedError(err)
		return errors.New(DATABASE_CREATE_FAIL)
	}
	service.PrintGreenStatus("Crypto_Currency created", obj)
	return nil
}

func CreateCryptoSource(exch string, exchCode int, ctx context.Context, client *ent.Client) error {
	obj, err := client.CRYPTO_PRC_SOURCE.
		Create().
		SetID(exchCode).
		SetName(exch).
		Save(ctx)
	if recover() != nil || err != nil {
		service.PrintRedError(err)
		return errors.New(DATABASE_CREATE_FAIL)
	}
	service.PrintGreenStatus("Crypto_Source created", obj)
	return nil
}

func CreateEmployType(permaBool, payFreq string, ctx context.Context, client *ent.Client) error {
	obj, err := client.EMPLOY_TYPE.
		Create().
		SetIsPermanent(permaBool).
		SetPayFreq(payFreq).
		Save(ctx)
	if recover() != nil || err != nil {
		service.PrintRedError(err)
		return errors.New(DATABASE_CREATE_FAIL)
	}
	service.PrintGreenStatus("Employ_Type created", obj)
	return nil
}

func CreateEmployee(entity ent.EMPLOYEE, employerGid, currency, employType, payFreq string, ctx context.Context, client *ent.Client) error {
	gidNew := uuid.New()
	obj, err := client.EMPLOYEE.
		Create().
		SetGid(gidNew.String()).
		// Personal information
		SetName(entity.Name).SetPosition(entity.Position).SetEmail(entity.Email).
		// Personal work information
		SetWorkStart(entity.WorkStart).SetWorkEnds(entity.WorkEnds).
		// Compensation, and receiving information
		SetWallet(entity.Wallet).SetPayroll(entity.Payroll).SetPayday(entity.Payday).
		// Datapoint edge settings
		// 1) Employee's employ type `employee_from_employ_type`
		SetEmployeeFromEmployType(
			client.EMPLOY_TYPE.
				Query().
				Where(
					employ_type.IsPermanent(employType),
					employ_type.PayFreq(payFreq),
				).
				OnlyX(ctx),
		).
		// 2) Employee's Employer `employee_from_employer`
		SetEmployeeFromEmployer(
			client.EMPLOYER.
				Query().
				Where(
					employer.Gid(employerGid),
				).
				OnlyX(ctx),
		).
		// 3) Employee's payment currency
		SetEmployeeFromCurrency(
			client.CRYPTO_CURRENCY.
				Query().
				Where(
					crypto_currency.Ticker(currency),
				).
				OnlyX(ctx),
		).
		// Database record purpose
		SetCreatedAt(time.Now()).SetUpdatedAt(time.Now()).
		SetCreatedBy(entity.CreatedBy).SetUpdatedBy(entity.UpdatedBy).
		Save(ctx)

	if recover() != nil || err != nil {
		msg := fmt.Sprintf(
			"Cannot create Employee with %v. Employer Gid info: %s. Type: %s:%s",
			entity, employerGid, employType, payFreq,
		)
		service.PrintRedError(err)
		service.PrintRedError("Additional Info", msg)
		return errors.New(DATABASE_CREATE_FAIL)
	}
	service.PrintGreenStatus("Employee created", obj)
	return nil
}

func CreateEmployer(entity ent.EMPLOYER, ctx context.Context, client *ent.Client) error {
	gidNew := uuid.New()
	obj, err := client.EMPLOYER.
		Create().
		SetUsername(entity.Username).
		SetPassword(entity.Password).
		SetGid(gidNew.String()).
		SetCorpName(entity.CorpName).
		SetCorpEmail(entity.CorpEmail).
		SetWallet(entity.Wallet).
		SetCreatedAt(time.Now()).
		SetCreatedBy(entity.CreatedBy).
		SetUpdatedAt(time.Now()).
		SetUpdatedBy(entity.UpdatedBy).
		Save(ctx)
	if recover() != nil || err != nil {
		service.PrintRedError(err)
		return errors.New(DATABASE_CREATE_FAIL)
	}
	service.PrintGreenStatus("Employer_User_Info created", obj)
	return nil
}

func CreatePaymentHistory(entity ent.EMPLOYEE, ctx context.Context, client *ent.Client) error {
	service.PrintYellowStatus("Creating Payment Log for employee", entity.Gid)
	obj, err := client.PAYMENT_HISTORY.
		Create().
		// Payment information
		SetCreatedAt(time.Now()).SetAmount(entity.Payroll).
		// Datapoint edge settings
		// 1) PaymentHistory of which Employee?
		SetPaymentHistoryFromEmployeeID(entity.EmployerID).
		// 2) PaymentHistory of which Employer?
		SetPaymentHistoryFromEmployerID(entity.ID).
		// 3) PaymentHistory of what currency?
		SetPaymentHistoryFromCurrencyIDID(entity.CryptoCurrencyID).
		Save(ctx)
	if recover() != nil || err != nil {
		service.PrintRedError(err)
		return errors.New(DATABASE_CREATE_FAIL)
	}
	service.PrintGreenStatus("Payment_History created", obj)
	return nil
}

func CreateTrLog(entity ent.Tr_log, ctx context.Context, client *ent.Client) error {
	obj, err := client.Tr_log.
		Create().
		SetTrType(entity.TrType).
		SetCreatedAt(time.Now()).
		Save(ctx)
	if recover() != nil || err != nil {
		service.PrintRedError(err)
		return errors.New(DATABASE_CREATE_FAIL)
	}
	service.PrintGreenStatus("Trlog created", obj)
	return nil
}
