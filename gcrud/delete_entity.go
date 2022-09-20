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

func DeleteEmployeewEmployer(employerGid string, ctx context.Context, client *ent.Client) {
	fmt.Printf("Deleting employee with employer gid %s", employerGid)
	_, err := client.EMPLOYEE.
		Delete().
		Where(employee.EmployerGid(employerGid)).
		Exec(ctx)
	if err != nil {
		fmt.Println("employee delete with employerGid failed: ", err)
		return
	}
}
