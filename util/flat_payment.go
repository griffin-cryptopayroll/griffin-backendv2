package util

import (
	"griffin-dao/ent"
	"time"
)

type Payment struct {
	Name        string    `json:"name"`
	EmployeeGid string    `json:"gid"`
	Position    string    `json:"position"`
	Currency    string    `json:"currency"`
	Payroll     float64   `json:"payroll"`
	ScheduleDt  time.Time `json:"payday_schedule"`
	ExecutedDt  time.Time `json:"payday_execute"`
}

func FlattenPayment(arr []*ent.PAYMENT) []Payment {
	res := []Payment{}
	for _, v := range arr {
		p := Payment{
			Name:        v.Edges.PaymentFromEmployee.Name,
			EmployeeGid: v.Edges.PaymentFromEmployee.Gid,
			Position:    v.Edges.PaymentFromEmployee.Position,
			Currency:    v.Edges.PaymentFromCurrency.Ticker,
			Payroll:     v.PaymentAmount,
			ScheduleDt:  v.PaymentScheduled,
			ExecutedDt:  v.PaymentExecuted,
		}
		res = append(res, p)
	}
	return res
}
