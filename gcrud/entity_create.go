package gcrud

import (
	"errors"
	"griffin-dao/ent"
	"griffin-dao/ent/employ_type"
	"griffin-dao/ent/employer_user_info"
	"griffin-dao/service"
	"time"

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
	if recover() != nil || err != nil {
		service.PrintRedError(err)
		return
	}
	service.PrintGreenStatus("Crypto_Currency created", obj)
}

func CreateCryptoSource(exch string, exchCode int, ctx context.Context, client *ent.Client) {
	obj, err := client.CRYPTO_PRC_SOURCE.
		Create().
		SetID(exchCode).
		SetName(exch).
		Save(ctx)
	if recover() != nil || err != nil {
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

func CreateEmployee(entity EmployeeJson, employType, payFreq string, ctx context.Context, client *ent.Client) error {
	gid, err := client.EMPLOYER_USER_INFO.
		Query().
		Where(
			employer_user_info.Gid(entity.EmployerGriffinID),
		).
		Only(ctx)
	if err != nil || recover() != nil {
		service.PrintRedError("No such GID")
		return errors.New("no such GID")
	}
	emp, err := client.EMPLOY_TYPE.
		Query().
		Where(
			employ_type.IsPermanent(employType),
			employ_type.PayFreq(payFreq),
		).
		Only(ctx)
	if err != nil || recover() != nil {
		service.PrintRedError("No such employ_type")
		return errors.New("no such employ_type")
	}
	gidNew := uuid.New()
	obj, err := client.EMPLOYEE.
		Create().
		SetGid(gidNew.String()).
		SetEmployerID(gid.ID). // uuid
		SetName(entity.Name).
		SetPosition(entity.Position).
		SetWallet(entity.Wallet).
		SetPayroll(entity.Payroll).
		SetCurrency(entity.Currency).
		SetPayday(entity.PayDay).
		SetEmail(entity.Email).
		SetWorkStart(entity.WorkStart).
		SetWorkEnds(entity.WorkEnd).
		SetCreatedAt(entity.CreatedAt).
		SetCreatedBy(entity.CreatedBy).
		SetUpdatedAt(entity.UpdatedAt).
		SetUpdatedBy(entity.UpdatedBy).
		SetEmployeeTypeFrom(emp).
		Save(ctx)
	if recover() != nil || err != nil {
		service.PrintRedError(err)
		return errors.New(DATABASE_CREATE_FAIL)
	}
	service.PrintGreenStatus("Employee created", obj)
	return nil
}

func CreateEmployerUserInfo(entity EmployerJson, ctx context.Context, client *ent.Client) error {
	gidNew := uuid.New()
	obj, err := client.EMPLOYER_USER_INFO.
		Create().
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
	if recover() != nil || err != nil {
		service.PrintRedError(err)
		return errors.New(DATABASE_CREATE_FAIL)
	}
	service.PrintGreenStatus("Employer_User_Info created", obj)
	return nil
}

func CreatePaymentHistory(entity ent.PAYMENT_HISTORY, ctx context.Context, client *ent.Client) {
	obj, err := client.PAYMENT_HISTORY.
		Create().
		SetEmployeeGid(entity.EmployeeGid).
		SetCreatedAt(entity.CreatedAt).
		SetCreatedBy(entity.CreatedBy).
		Save(ctx)
	if recover() != nil || err != nil {
		service.PrintRedError(err)
		return
	}
	service.PrintGreenStatus("Payment_History created", obj)
}

func CreateTrLog(entity ent.Tr_log, ctx context.Context, client *ent.Client) {
	obj, err := client.Tr_log.
		Create().
		SetTrType(entity.TrType).
		SetCreatedAt(time.Now()).
		Save(ctx)
	if recover() != nil || err != nil {
		service.PrintRedError(err)
		return
	}
	service.PrintGreenStatus("Trlog created", obj)
}
