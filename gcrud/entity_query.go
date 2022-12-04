package gcrud

import (
	"context"
	"errors"
	"griffin-dao/ent"
	"griffin-dao/ent/crypto_currency"
	"griffin-dao/ent/employ_type"
	"griffin-dao/ent/employee"
	"griffin-dao/ent/employer"
	"griffin-dao/ent/tr_log"
	"griffin-dao/service"
)

func LoginHandler(email string, ctx context.Context, client *ent.Client) (*ent.EMPLOYER, error) {
	service.PrintYellowStatus("Query login info")
	obj, err := client.EMPLOYER.
		Query().
		Where(employer.CorpEmail(email)).
		Only(ctx)
	if recover() != nil || err != nil {
		service.PrintRedError("login query failed - no email")
		return nil, errors.New(DATABASE_SELECT_FAIL)
	}
	return obj, nil
}

func QueryEmployerwEmployerGid(employerGid string, ctx context.Context, client *ent.Client) (*ent.EMPLOYER, error) {
	service.PrintYellowStatus("Query employer with employerGid")
	obj, err := client.EMPLOYER.
		Query().
		Where(employer.Gid(employerGid)).
		Only(ctx)
	if recover() != nil || err != nil {
		service.PrintRedError("employer query with employerGid failed: ", err)
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

func QueryEmployeewEmployerGid(employerGid string, ctx context.Context, client *ent.Client) ([]*ent.EMPLOYEE, error) {
	service.PrintYellowStatus("Query all employee with their employerGid")
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
		service.PrintRedError("employee query with their employerGid failed: ", err)
		return nil, errors.New(DATABASE_SELECT_FAIL)
	}
	return obj, nil
}

func QueryEmployee(employeeGid, employerGid string, ctx context.Context, client *ent.Client) (*ent.EMPLOYEE, error) {
	service.PrintYellowStatus("Query single employee with their employerGid && employeeGid")
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
		service.PrintRedError("employee query with their employeeGid && employerGid failed: ", err)
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

func QueryTrLog(trType string, ctx context.Context, client *ent.Client) ([]*ent.Tr_log, error) {
	service.PrintYellowStatus("Query tr log by its tr_type")
	obj, err := client.Tr_log.
		Query().
		Where(
			tr_log.TrType(trType),
		).
		All(ctx)
	if recover() != nil || err != nil {
		service.PrintRedError("tr_log query with their tr_type failed: ", err)
		return nil, errors.New(DATABASE_SELECT_FAIL)
	}
	return obj, nil
}
