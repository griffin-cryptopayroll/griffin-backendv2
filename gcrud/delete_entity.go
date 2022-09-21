package gcrud

import (
	"context"
	"fmt"
	"griffin-dao/ent"
	"griffin-dao/ent/employee"
	"griffin-dao/ent/employer_user_info"
)

func DeleteEmployer(employerGid string, ctx context.Context, client *ent.Client) {
	fmt.Println("Deleting employer")
	_, err := client.EMPLOYER_USER_INFO.
		Delete().
		Where(employer_user_info.Gid(employerGid)).
		Exec(ctx)
	if err != nil {
		fmt.Println("employer delete with employerGid failed: ", err)
		return
	}
}

func DeleteEmployeewEmployerAll(employerGid string, ctx context.Context, client *ent.Client) {
	fmt.Printf("Deleting all employee with employer gid %s\n", employerGid)
	_, err := client.EMPLOYEE.
		Delete().
		Where(employee.EmployerGid(employerGid)).
		Exec(ctx)
	if err != nil {
		fmt.Println("employee delete with employerGid failed: ", err)
		return
	}
}

func DeleteEmployeewEmployerInd(employerGid, employeeGid string, ctx context.Context, client *ent.Client) {
	fmt.Printf("Deleting employee %s with employer gid %s\n", employeeGid, employerGid)
	_, err := client.EMPLOYEE.
		Delete().
		Where(
			employee.EmployerGid(employerGid),
			employee.Gid(employeeGid),
		).
		Exec(ctx)
	if err != nil {
		fmt.Println("employee delete with employerGid failed: ", err)
	}
}
