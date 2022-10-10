package gcrud

import (
	"context"
	"fmt"
	"griffin-dao/ent"
	"griffin-dao/ent/employ_type"
	"griffin-dao/ent/employee"
	"griffin-dao/ent/employer_user_info"
	"griffin-dao/service"
)

func DeleteEmployer(employerGid string, ctx context.Context, client *ent.Client) {
	fmt.Println("Deleting employer")
	delNum, err := client.EMPLOYER_USER_INFO.
		Delete().
		Where(employer_user_info.Gid(employerGid)).
		Exec(ctx)
	if err != nil {
		service.PrintRedError("employer delete with employerGid failed: ", err)
		return
	}
	service.PrintGreenStatus("employer deleted: ", delNum)
}

func DeleteEmployeewEmployerAll(employerGid string, ctx context.Context, client *ent.Client) {
	fmt.Printf("Deleting all employee with employer gid %s\n", employerGid)
	delNum, err := client.EMPLOYEE.
		Delete().
		Where(employee.EmployerGid(employerGid)).
		Exec(ctx)
	if err != nil {
		service.PrintRedError("employee delete with employerGid failed: ", err)
		return
	}
	service.PrintGreenStatus("employee deleted: ", delNum)
}

func DeleteEmployeewEmployerInd(employerGid, employeeGid string, ctx context.Context, client *ent.Client) {
	fmt.Printf("Deleting employee %s with employer gid %s\n", employeeGid, employerGid)
	delNum, err := client.EMPLOYEE.
		Delete().
		Where(
			employee.EmployerGid(employerGid),
			employee.Gid(employeeGid),
		).
		Exec(ctx)
	if err != nil {
		service.PrintRedError("employer delete with employerGid failed: ", err)
		return
	}
	service.PrintGreenStatus("employer deleted: ", delNum)
}

func DeleteEmployType(contractMonth int, ctx context.Context, client *ent.Client) {
	fmt.Printf("Deleting employ type with contract month of %v", contractMonth)
	delNum, err := client.EMPLOY_TYPE.
		Delete().
		Where(
			employ_type.ContractPeriod(contractMonth),
		).
		Exec(ctx)
	if recover() != nil || err != nil {
		service.PrintRedError("employee type delete with contract month failed: ", err)
		return
	}
	service.PrintGreenStatus("employ type deleted: ", delNum)
}
