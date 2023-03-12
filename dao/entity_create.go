package dao

import (
	"context"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"griffin-dao/ent"
	"griffin-dao/ent/crypto_currency"
	"griffin-dao/ent/employ_type"
	"griffin-dao/ent/employer"
	"griffin-dao/util"
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
		util.PrintRedError(err)
		return errors.New(DATABASE_CREATE_FAIL)
	}
	util.PrintGreenStatus("Crypto_Currency created", obj)
	return nil
}

func CreateCryptoSource(exch string, exchCode int, ctx context.Context, client *ent.Client) error {
	obj, err := client.CRYPTO_PRC_SOURCE.
		Create().
		SetID(exchCode).
		SetName(exch).
		Save(ctx)
	if recover() != nil || err != nil {
		util.PrintRedError(err)
		return errors.New(DATABASE_CREATE_FAIL)
	}
	util.PrintGreenStatus("Crypto_Source created", obj)
	return nil
}

func CreateEmployType(permaBool, payFreq string, ctx context.Context, client *ent.Client) error {
	obj, err := client.EMPLOY_TYPE.
		Create().
		SetIsPermanent(permaBool).
		SetPayFreq(payFreq).
		Save(ctx)
	if recover() != nil || err != nil {
		util.PrintRedError(err)
		return errors.New(DATABASE_CREATE_FAIL)
	}
	util.PrintGreenStatus("Employ_Type created", obj)
	return nil
}

func CreateEmployee(entity ent.EMPLOYEE, employerGid, currency, employType, payFreq string, ctx context.Context, client *ent.Client) (*ent.EMPLOYEE, error) {
	gidNew := uuid.New()
	obj, err := client.EMPLOYEE.
		Create().
		SetGid(gidNew.String()).
		// Personal information
		SetName(entity.Name).SetPosition(entity.Position).SetEmail(entity.Email).
		// Personal work information
		SetWorkStart(entity.WorkStart).SetWorkEnds(entity.WorkEnds).
		// Compensation, and receiving information
		SetWalletAztec(entity.WalletAztec).SetWallet(entity.Wallet).SetPayroll(entity.Payroll).SetPayday(entity.Payday).
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
		util.PrintRedError(err)
		util.PrintRedError("Additional Info", msg)
		return nil, errors.New(DATABASE_CREATE_FAIL)
	}
	util.PrintGreenStatus("Employee created", obj)
	return obj, nil
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
		SetWalletAztec(entity.WalletAztec).
		SetCreatedAt(time.Now()).
		SetCreatedBy(entity.CreatedBy).
		SetUpdatedAt(time.Now()).
		SetUpdatedBy(entity.UpdatedBy).
		Save(ctx)
	if recover() != nil || err != nil {
		util.PrintRedError(err)
		return errors.New(DATABASE_CREATE_FAIL)
	}
	util.PrintGreenStatus("Employer_User_Info created", obj)
	return nil
}

func CreatePaymentHistory(entity ent.EMPLOYEE, ctx context.Context, client *ent.Client) error {
	util.PrintYellowStatus("Creating Payment Log for employee", entity.Gid)
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
		util.PrintRedError(err)
		return errors.New(DATABASE_CREATE_FAIL)
	}
	util.PrintGreenStatus("Payment_History created", obj)
	return nil
}

func CreateTrLog(entity ent.Tr_log, ctx context.Context, client *ent.Client) error {
	util.PrintYellowStatus("Creating TR LOG")
	obj, err := client.Tr_log.
		Create().
		SetTrType(entity.TrType).
		SetCreatedAt(time.Now()).
		Save(ctx)
	if recover() != nil || err != nil {
		util.PrintRedError(err)
		return errors.New(DATABASE_CREATE_FAIL)
	}
	util.PrintGreenStatus("Trlog created", obj)
	return nil
}

func CreatePaymentScheduled(entity *ent.EMPLOYEE, scheduled time.Time, ctx context.Context, client *ent.Client) error {
	util.PrintYellowStatus("Creating Scheduled Payment")
	obj, err := client.PAYMENT.
		Create().
		SetEmployeeID(entity.ID).
		SetEmployerID(entity.EmployerID).
		SetPaymentScheduled(scheduled).
		SetPaymentAmount(entity.Payroll).
		SetCryptoCurrencyID(entity.CryptoCurrencyID).
		Save(ctx)
	if recover() != nil || err != nil {
		util.PrintRedError(err)
		return errors.New(DATABASE_CREATE_FAIL)
	}
	util.PrintGreenStatus("Scheduled Payment created", obj)
	return nil
}

func CreatePermanent(entity *ent.EMPLOYEE, workStart time.Time, interval string, ctx context.Context, client *ent.Client) error {
	util.PrintYellowStatus("Creating Scheduled Payment for permanent worker", entity.Gid)
	switch interval {
	case "D":
		// generate 365 days
		for i := 0; i < 365; i++ {
			schd := workStart.Add(time.Hour * 24 * time.Duration(i))
			err := CreatePaymentScheduled(entity, schd, ctx, client)
			if err != nil {
				return err
			}
		}
		util.PrintGreenStatus("Permanent worker scheduled payment created")
	case "W":
		// generate 52 weeks
		for i := 0; i < 52; i++ {
			schd := workStart.Add(time.Hour * 24 * time.Duration(7) * time.Duration(i))
			err := CreatePaymentScheduled(entity, schd, ctx, client)
			if err != nil {
				return err
			}
		}
	case "M":
		for i := 0; i < 12; i++ {
			schd := dateAddMonth(workStart, i)
			err := CreatePaymentScheduled(entity, schd, ctx, client)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func dateAddMonth(date time.Time, howMany int) time.Time {
	var (
		y int
		m int
		d int
	)
	if date.Month() == 12 {
		y = date.Year() + 1*howMany
		m = 1
		d = date.Day()
	} else {
		y = date.Year()
		m = int(date.Month()) + 1*howMany
		d = date.Day()
	}

	return time.Date(y, time.Month(m), d, 0, 0, 0, 0, date.Location())
}

func CreateFreelance(entity *ent.EMPLOYEE, workStart, workEnd time.Time, interval string, ctx context.Context, client *ent.Client) error {
	util.PrintYellowStatus("Creating Scheduled Payment for freelance worker", entity.Gid)
	switch interval {
	case "D":
		for i := 0; i < 365; i++ {
			schd := workStart.Add(time.Hour * 24 * time.Duration(i))
			if schd.After(workEnd) {
				break
			}
			err := CreatePaymentScheduled(entity, schd, ctx, client)
			if err != nil {
				return err
			}
		}
	case "W":
		for i := 0; i < 52; i++ {
			schd := workStart.Add(time.Hour * 24 * time.Duration(7) * time.Duration(i))
			if schd.After(workEnd) {
				break
			}
			err := CreatePaymentScheduled(entity, schd, ctx, client)
			if err != nil {
				return err
			}
		}
	case "M":
		for i := 0; i < 12; i++ {
			schd := dateAddMonth(workStart, i)
			if schd.After(workEnd) {
				break
			}
			err := CreatePaymentScheduled(entity, schd, ctx, client)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func CreatePaymentOneOff(entity *ent.EMPLOYEE, oneOff time.Time, ctx context.Context, client *ent.Client) error {
	obj, err := client.PAYMENT.
		Create().
		SetEmployeeID(entity.ID).
		SetEmployerID(entity.EmployerID).
		SetPaymentExecuted(oneOff).
		SetPaymentAmount(entity.Payroll).
		SetCryptoCurrencyID(entity.CryptoCurrencyID).
		Save(ctx)
	if recover() != nil || err != nil {
		util.PrintRedError(err)
		return errors.New(DATABASE_CREATE_FAIL)
	}
	util.PrintGreenStatus("OneOff Payment created", obj)
	return nil
}
