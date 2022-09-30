package gcrud

import (
	"context"
	"griffin-dao/ent"
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
