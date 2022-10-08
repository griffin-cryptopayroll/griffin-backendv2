package gcrud

import (
	"context"
	"griffin-dao/ent"
	"griffin-dao/ent/employ_type"
	"griffin-dao/ent/employee"
	"griffin-dao/ent/employer_user_info"
	"griffin-dao/service"
)

func QueryEmployer(employerGid string, ctx context.Context, client *ent.Client) *ent.EMPLOYER_USER_INFO {
	service.PrintYellowStatus("Query employer")
	obj, err := client.EMPLOYER_USER_INFO.
		Query().
		Where(employer_user_info.Gid(employerGid)).
		Only(ctx)
	if err != nil {
		service.PrintRedError("employer query with employerGid failed: ", err)
		return nil
	}
	return obj
}

func QueryEmployeewEmployerGid(employerGid string, ctx context.Context, client *ent.Client) []*ent.EMPLOYEE {
	service.PrintYellowStatus("Query all employee with their employer")
	obj, err := client.EMPLOYEE.
		Query().
		Where(employee.EmployerGid(employerGid)).
		All(ctx)
	if err != nil {
		service.PrintRedError("employee query with their employerGid failed: ", err)
		return nil
	}
	return obj
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
	if err != nil {
		service.PrintRedError("employee query with their employerGid and type failed: ", err)
	}
	return obj
}

func QueryEmployee(employeeGid, employerGid string, ctx context.Context, client *ent.Client) *ent.EMPLOYEE {
	service.PrintYellowStatus("Query single employee with their employer")
	obj, err := client.EMPLOYEE.
		Query().
		Where(
			employee.Gid(employeeGid),
			employee.EmployerGid(employerGid),
		).
		Only(ctx)
	if err != nil {
		service.PrintRedError("employee query with their employeeGid failed: ", err)
		return nil
	}
	return obj
}

func QueryEmployType(contractMonth int, ctx context.Context, client *ent.Client) *ent.EMPLOY_TYPE {
	service.PrintYellowStatus("Query employee type")
	obj, err := client.EMPLOY_TYPE.
		Query().
		Where(employ_type.ContractPeriod(contractMonth)).
		Only(ctx)
	if recover() != nil || err != nil {
		service.PrintRedError("employ type query with their contract period failed: ", err)
		return nil
	}
	return obj
}
