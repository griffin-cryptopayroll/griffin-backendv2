package gcrud

import (
	"context"
	"fmt"
	"griffin-dao/ent"
	"griffin-dao/ent/employee"
	"griffin-dao/ent/employer_user_info"
)

func QueryEmployer(employerGid string, ctx context.Context, client *ent.Client) *ent.EMPLOYER_USER_INFO {
	fmt.Println("Query employer")
	obj, err := client.EMPLOYER_USER_INFO.
		Query().
		Where(employer_user_info.Gid(employerGid)).
		Only(ctx)
	if err != nil {
		fmt.Println("employer query with employerGid failed: ", err)
		return nil
	}
	return obj
}

func QueryEmployeewEmployerGid(employerGid string, ctx context.Context, client *ent.Client) []*ent.EMPLOYEE {
	fmt.Println("Query all employee with their employer")
	obj, err := client.EMPLOYEE.
		Query().
		Where(employee.EmployerGid(employerGid)).
		All(ctx)
	if err != nil {
		fmt.Println("employee query with their employerGid failed: ", err)
		return nil
	}
	return obj
}

func QueryEmployee(employeeGid string, ctx context.Context, client *ent.Client) *ent.EMPLOYEE {
	fmt.Println("Query single employee with their employer")
	obj, err := client.EMPLOYEE.
		Query().
		Where(employee.Gid(employeeGid)).
		Only(ctx)
	if err != nil {
		fmt.Println("employee query with their employeeGid failed: ", err)
		return nil
	}
	return obj
}
