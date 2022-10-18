package gcrud

import (
	"context"
	"errors"
	"griffin-dao/ent"
	"griffin-dao/ent/crypto_currency"
	"griffin-dao/ent/employ_type"
	"griffin-dao/ent/employee"
	"griffin-dao/ent/employer_user_info"
	"griffin-dao/service"
)

func QueryLoginPassword(email string, ctx context.Context, client *ent.Client) (*ent.EMPLOYER_USER_INFO, error) {
	service.PrintYellowStatus("Query login info")
	obj, err := client.EMPLOYER_USER_INFO.
		Query().
		Where(employer_user_info.CorpEmail(email)).
		Only(ctx)
	if recover() != nil || err != nil {
		service.PrintRedError("login query failed")
		return nil, errors.New(DATABASE_SELECT_FAIL)
	}
	return obj, nil
}

func QueryCurrency(curr string, ctx context.Context, client *ent.Client) (*ent.CRYPTO_CURRENCY, error) {
	service.PrintYellowStatus("Query currency with ticker")
	obj, err := client.CRYPTO_CURRENCY.
		Query().
		Where(crypto_currency.Ticker(curr)).
		Only(ctx)
	if recover() != nil || err != nil {
		service.PrintRedError("crypto currency type query failed")
		return nil, errors.New(DATABASE_SELECT_FAIL)
	}
	return obj, nil
}

func QueryEmployerwEmployerGid(employerGid string, ctx context.Context, client *ent.Client) (*ent.EMPLOYER_USER_INFO, error) {
	service.PrintYellowStatus("Query employer with employerGid")
	obj, err := client.EMPLOYER_USER_INFO.
		Query().
		Where(employer_user_info.Gid(employerGid)).
		Only(ctx)
	if recover() != nil || err != nil {
		service.PrintRedError("employer query with employerGid failed: ", err)
		return nil, errors.New(DATABASE_SELECT_FAIL)
	}
	return obj, nil
}

func QueryEmployeewEmployerGid(employerGid string, ctx context.Context, client *ent.Client) ([]*ent.EMPLOYEE, error) {
	service.PrintYellowStatus("Query all employee with their employer")
	obj, err := client.EMPLOYEE.
		Query().
		Where(employee.EmployerGid(employerGid)).
		All(ctx)
	if recover() != nil || err != nil {
		service.PrintRedError("employee query with their employerGid failed: ", err)
		return nil, errors.New(DATABASE_SELECT_FAIL)
	}
	return obj, nil
}

func QueryEmployeewEmployerGidType(employerGid string, employType int, ctx context.Context, client *ent.Client) []*ent.EMPLOYEE {
	service.PrintYellowStatus("Query all employee with their employer + suggested type")
	obj, err := client.EMPLOYEE.
		Query().
		Where(
			employee.EmployerGid(employerGid),
			employee.Employ(employType),
		).
		All(ctx)
	if recover() != nil || err != nil {
		service.PrintRedError("employee query with their employerGid and type failed: ", err)
	}
	return obj
}

func QueryEmployee(employeeGid, employerGid string, ctx context.Context, client *ent.Client) (*ent.EMPLOYEE, error) {
	service.PrintYellowStatus("Query single employee with their employer")
	obj, err := client.EMPLOYEE.
		Query().
		Where(
			employee.Gid(employeeGid),
			employee.EmployerGid(employerGid),
		).
		Only(ctx)
	if recover() != nil || err != nil {
		service.PrintRedError("employee query with their employeeGid failed: ", err)
		return nil, errors.New(DATABASE_SELECT_FAIL)
	}
	return obj, nil
}

func QueryEmployType(isPerma, payFreq string, ctx context.Context, client *ent.Client) (*ent.EMPLOY_TYPE, error) {
	service.PrintYellowStatus("Query employee type by isPermanent or not")
	obj, err := client.EMPLOY_TYPE.
		Query().
		Where(
			employ_type.IsPermanent(isPerma),
			employ_type.PayFreq(payFreq),
		).
		Only(ctx)
	if recover() != nil || err != nil {
		service.PrintRedError("employ type query with their contract period failed: ", err)
		return nil, errors.New(DATABASE_SELECT_FAIL)
	}
	return obj, nil
}
