package dao

import (
	"context"
	"errors"
	"fmt"
	"griffin-dao/ent"
	"griffin-dao/ent/crypto_currency"
	"griffin-dao/ent/employ_type"
	"griffin-dao/ent/employee"
	"griffin-dao/ent/employer"
	"griffin-dao/ent/payment"
	"griffin-dao/ent/tr_log"
	"griffin-dao/util"
	"strconv"
	"strings"
	"time"
)

func LoginHandler(email string, ctx context.Context, client *ent.Client) (*ent.EMPLOYER, error) {
	util.PrintYellowStatus("Query login info")
	obj, err := client.EMPLOYER.
		Query().
		Where(employer.CorpEmail(email)).
		Only(ctx)
	if recover() != nil || err != nil {
		util.PrintRedError("login query failed - no email")
		return nil, errors.New(DATABASE_SELECT_FAIL)
	}
	return obj, nil
}

func QueryEmployerwEmployerGid(employerGid string, ctx context.Context, client *ent.Client) (*ent.EMPLOYER, error) {
	util.PrintYellowStatus("Query employer with employerGid")
	obj, err := client.EMPLOYER.
		Query().
		Where(employer.Gid(employerGid)).
		Only(ctx)
	if recover() != nil || err != nil {
		util.PrintRedError("employer query with employerGid failed: ", err)
		return nil, errors.New(DATABASE_SELECT_FAIL)
	}
	return obj, nil
}

func QueryCurrency(curr string, ctx context.Context, client *ent.Client) (*ent.CRYPTO_CURRENCY, error) {
	util.PrintYellowStatus("Query currency with ticker")
	obj, err := client.CRYPTO_CURRENCY.
		Query().
		Where(crypto_currency.Ticker(curr)).
		Only(ctx)
	if recover() != nil || err != nil {
		util.PrintRedError("crypto currency type query failed")
		return nil, errors.New(DATABASE_SELECT_FAIL)
	}
	return obj, nil
}

func QueryEmployeewEmployerGid(employerGid string, ctx context.Context, client *ent.Client) ([]*ent.EMPLOYEE, error) {
	util.PrintYellowStatus("Query all employee with their employerGid")
	obj, err := client.EMPLOYEE.
		Query().
		Where(
			employee.HasEmployeeFromEmployerWith(
				employer.Gid(employerGid),
			),
		).
		WithEmployeeFromCurrency().
		WithEmployeeFromEmployType().
		All(ctx)
	if recover() != nil || err != nil {
		util.PrintRedError("employee query with their employerGid failed: ", err)
		return nil, errors.New(DATABASE_SELECT_FAIL)
	}
	return obj, nil
}

func QueryEmployee(employeeGid, employerGid string, ctx context.Context, client *ent.Client) (*ent.EMPLOYEE, error) {
	util.PrintYellowStatus("Query single employee with their employerGid && employeeGid")
	obj, err := client.EMPLOYEE.
		Query().
		Where(
			employee.Gid(employeeGid),
			employee.HasEmployeeFromEmployerWith(
				employer.Gid(employerGid),
			),
		).
		WithEmployeeFromCurrency().
		WithEmployeeFromEmployType().
		Only(ctx)
	if recover() != nil || err != nil {
		util.PrintRedError("employee query with their employeeGid && employerGid failed: ", err)
		return nil, errors.New(DATABASE_SELECT_FAIL)
	}
	return obj, nil
}

func QueryEmployType(isPerma, payFreq string, ctx context.Context, client *ent.Client) (*ent.EMPLOY_TYPE, error) {
	util.PrintYellowStatus("Query employee type by isPermanent or not")
	obj, err := client.EMPLOY_TYPE.
		Query().
		Where(
			employ_type.IsPermanent(isPerma),
			employ_type.PayFreq(payFreq),
		).
		Only(ctx)
	if recover() != nil || err != nil {
		util.PrintRedError("employ type query with their contract period failed: ", err)
		return nil, errors.New(DATABASE_SELECT_FAIL)
	}
	return obj, nil
}

func QueryTrLog(trType string, ctx context.Context, client *ent.Client) ([]*ent.Tr_log, error) {
	util.PrintYellowStatus("Query tr log by its tr_type")
	obj, err := client.Tr_log.
		Query().
		Where(
			tr_log.TrType(trType),
		).
		All(ctx)
	if recover() != nil || err != nil {
		util.PrintRedError("tr_log query with their tr_type failed: ", err)
		return nil, errors.New(DATABASE_SELECT_FAIL)
	}
	return obj, nil
}

func QueryPaymentEmployee(employeeGid string, ctx context.Context, client *ent.Client) ([]*ent.PAYMENT, error) {
	obj, err := client.PAYMENT.
		Query().
		Where(payment.HasPaymentFromEmployeeWith(
			employee.Gid(employeeGid),
		)).
		All(ctx)
	if recover() != nil || err != nil {
		util.PrintRedError(err)
		return nil, errors.New(DATABASE_SELECT_FAIL)
	}
	util.PrintGreenStatus(fmt.Sprintf("Payment status for employee %s", employeeGid))
	return obj, nil
}

func QueryPaymentEmployer(employerGid string, ctx context.Context, client *ent.Client) ([]*ent.PAYMENT, error) {
	obj, err := client.PAYMENT.
		Query().
		Where(payment.HasPaymentFromEmployerWith(
			employer.Gid(employerGid),
		)).
		WithPaymentFromEmployee().
		WithPaymentFromCurrency().
		All(ctx)
	if recover() != nil || err != nil {
		util.PrintRedError(err)
		return nil, errors.New(DATABASE_SELECT_FAIL)
	}
	util.PrintGreenStatus(fmt.Sprintf("Payment status for employer %s", employerGid))
	return obj, nil
}

func PaymentScheduled(pay []*ent.PAYMENT) []*ent.PAYMENT {
	var scheduled []*ent.PAYMENT
	for _, p := range pay {
		if p.PaymentExecuted.IsZero() && !p.PaymentScheduled.IsZero() {
			// Scheduled, but not executed
			scheduled = append(scheduled, p)
		}
	}
	return scheduled
}

func PaymentExecuted(pay []*ent.PAYMENT) []*ent.PAYMENT {
	var executed []*ent.PAYMENT
	for _, p := range pay {
		if !p.PaymentExecuted.IsZero() && !p.PaymentScheduled.IsZero() {
			// Scheduled and executed
			executed = append(executed, p)
		}
	}
	return executed
}

func PaymentOneOff(pay []*ent.PAYMENT) []*ent.PAYMENT {
	var oneOff []*ent.PAYMENT
	for _, p := range pay {
		if !p.PaymentExecuted.IsZero() && p.PaymentScheduled.IsZero() {
			// Not scheduled, but executed
			oneOff = append(oneOff, p)
		}
	}
	return oneOff
}

func PaymentFuture(pay []*ent.PAYMENT, interval string) ([]*ent.PAYMENT, error) {
	var future []*ent.PAYMENT
	futureDays, err := parseInterval(interval)
	if err != nil {
		return nil, err
	}
	currentT := time.Now()
	futureT := time.Now().Add(time.Duration(futureDays) * time.Hour * 24)
	for _, p := range pay {
		if p.PaymentExecuted.IsZero() && !p.PaymentScheduled.IsZero() &&
			p.PaymentScheduled.Before(futureT) && p.PaymentScheduled.After(currentT) {
			future = append(future, p)
		}
	}
	return future, nil
}

func PaymentPast(pay []*ent.PAYMENT, interval string) ([]*ent.PAYMENT, error) {
	var past []*ent.PAYMENT
	pastDays, err := parseInterval(interval)
	if err != nil {
		return nil, err
	}
	currentT := time.Now()
	pastT := time.Now().Add(time.Duration(-1*pastDays) * time.Hour * 24)
	for _, p := range pay {
		if !p.PaymentExecuted.IsZero() && !p.PaymentScheduled.IsZero() &&
			p.PaymentExecuted.After(pastT) && p.PaymentScheduled.Before(currentT) {
			past = append(past, p)
		}
	}
	return past, nil
}

func PaymentMissed(pay []*ent.PAYMENT) []*ent.PAYMENT {
	var missed []*ent.PAYMENT
	std := time.Now()
	for _, p := range pay {
		if p.PaymentExecuted.IsZero() && !p.PaymentScheduled.IsZero() && p.PaymentScheduled.Before(std) {
			// Not executed payment before `time.Now()`
			missed = append(missed, p)
		}
	}
	return missed
}

func parseInterval(interval string) (int, error) {
	// returns period in date or error
	if len(interval) != 2 {
		msg := fmt.Sprintf("%s is not supported interval", interval)
		return 0, errors.New(msg)
	}
	supportedUnit := map[string]bool{
		"d": true,
		"m": true,
		"y": true,
	}
	approxUnit := map[string]int{
		"d": 1,
		"m": 30,
		"y": 365,
	}

	itvNum, err := strconv.Atoi(string(interval[0]))
	itvUnit := strings.ToLower(string(interval[1]))
	if err != nil {
		msg := fmt.Sprintf("%v in %s is not a number", itvNum, interval)
		return 0, errors.New(msg)
	}
	if !supportedUnit[itvUnit] {
		msg := fmt.Sprintf("%s has no supported unit(must be d, m, y)", interval)
		return 0, errors.New(msg)
	}
	return itvNum * approxUnit[itvUnit], nil
}
