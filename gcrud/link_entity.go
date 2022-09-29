package gcrud

import (
	"context"
	"fmt"
	"griffin-dao/ent"
	"griffin-dao/ent/crypto_currency"
	"griffin-dao/ent/crypto_prc_source"
	"griffin-dao/ent/employ_type"
	"griffin-dao/ent/employee"
	"griffin-dao/ent/employer_user_info"
	"griffin-dao/ent/payment_history"
)

func LinkCryptoCurrencywSource(curr string, exchCode int, ctx context.Context, client *ent.Client) {
	obj, err := client.CRYPTO_CURRENCY.
		Query().
		Where(crypto_currency.Ticker(curr)).
		Only(ctx)
	if err != nil {
		fmt.Println("object finding", err)
		return
	}
	source, err := client.CRYPTO_PRC_SOURCE.
		Query().
		Where(crypto_prc_source.ID(exchCode)).
		Only(ctx)
	if err != nil {
		fmt.Println("source finding", err)
		return
	}
	source.Update().AddPriceOf(obj)
}

func LinkCryptoCurrencywEmployee(employCurr int, curr string, ctx context.Context, client *ent.Client) {
	obj, err := client.CRYPTO_CURRENCY.
		Query().
		Where(crypto_currency.Ticker(curr)).
		Only(ctx)
	if err != nil {
		fmt.Println("object finding", err)
		return
	}
	emp, err := client.EMPLOYEE.
		Query().
		Where(employee.Currency(employCurr)).
		Only(ctx)
	if err != nil {
		fmt.Println("employee finding w currency", err)
		return
	}
	obj.Update().AddEmployeePaid(emp).Save(ctx)
}

func LinkEmployTypewEmployee(isPerma, employee_gid string, ctx context.Context, client *ent.Client) {
	employT, err := client.EMPLOY_TYPE.
		Query().
		Where(employ_type.IsPermanent(isPerma)).
		Only(ctx)
	if err != nil {
		fmt.Println("object finding", err)
		return
	}
	emp, err := client.EMPLOYEE.
		Query().
		Where(employee.Gid(employee_gid)).
		Only(ctx)
	if err != nil {
		fmt.Println("employee finding w gid", err)
		return
	}
	employT.Update().AddEmployeeTypeTo(emp).Save(ctx)
}

func LinkEmployeewEmployer(employerGid, employeeGid string, ctx context.Context, client *ent.Client) {
	employerInd, err := client.EMPLOYER_USER_INFO.
		Query().
		Where(employer_user_info.Gid(employerGid)).
		Only(ctx)
	if err != nil {
		fmt.Println("employer finding w gid", err)
		return
	}
	employeeInd, err := client.EMPLOYEE.
		Query().
		Where(employee.Gid(employeeGid)).
		Only(ctx)
	if err != nil {
		fmt.Println("employee finding w gid", err)
		return
	}
	employerInd.Update().AddWorkUnder(employeeInd).Save(ctx)
}

func LinkPaymentHistorywEmployee(employeeGid string, ctx context.Context, client *ent.Client) {
	payHistory, err := client.PAYMENT_HISTORY.
		Query().
		Where(payment_history.EmployeeGid(employeeGid)).
		Only(ctx)
	if err != nil {
		fmt.Println("payment history w gid", err)
		return
	}
	employeeInd, err := client.EMPLOYEE.
		Query().
		Where(employee.Gid(employeeGid)).
		Only(ctx)
	if err != nil {
		fmt.Println("employee finding w gid", err)
		return
	}
	employeeInd.Update().AddPaymentHistory(payHistory).Save(ctx)
}
