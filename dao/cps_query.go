package dao

import (
	"context"
	"errors"
	"griffin-dao/ent"
	"griffin-dao/ent/employee"
	"griffin-dao/ent/payment_history"
	"griffin-dao/util"
	"time"
)

// Subroutines for Composite Query

type Period struct {
	StandardDate time.Time
	Interval     time.Duration
}

func PastPayment(pastPeriod Period, ctx context.Context, client *ent.Client) ([]*ent.PAYMENT_HISTORY, error) {
	util.PrintYellowStatus("Query Past payment event")
	obj, err := client.PAYMENT_HISTORY.
		Query().
		// Payment history which `CreatedAt` is
		// `StandardDate - Interval` <= `CreatedAt` <= `StandardDate`
		Where(
			payment_history.CreatedAtLTE(pastPeriod.StandardDate),
			payment_history.CreatedAtGTE(pastPeriod.StandardDate.Add(pastPeriod.Interval*-1)),
		).
		// Add edges
		// 1) Payment history to whom? Which employee?
		WithPaymentHistoryFromEmployee().
		// 2) Payment in what currency?
		WithPaymentHistoryFromCurrencyID().
		All(ctx)
	if recover() != nil || err != nil {
		util.PrintRedError("payment_history history query failed")
		return nil, errors.New(DATABASE_SELECT_FAIL)
	}
	return obj, nil
}

func FuturePayment(futurePeriod Period, ctx context.Context, client *ent.Client) ([]*ent.EMPLOYEE, error) {
	util.PrintYellowStatus("Query Upcoming payment event")
	obj, err := client.EMPLOYEE.
		Query().
		// Payment history which `CreatedAt` is
		// `StandardDate` <= `CreatedAt` <= `StandardDate + Interval`
		Where(
			employee.PaydayLTE(futurePeriod.StandardDate.Add(futurePeriod.Interval)),
			employee.PaydayGTE(futurePeriod.StandardDate),
		).
		All(ctx)
	if recover() != nil || err != nil {
		util.PrintRedError("employee query for future payment failed")
		return nil, errors.New(DATABASE_SELECT_FAIL)
	}
	return obj, nil
}
