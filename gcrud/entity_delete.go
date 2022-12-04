package gcrud

import (
	"context"
	"errors"
	"fmt"
	"griffin-dao/ent"
	"griffin-dao/ent/employee"
	"griffin-dao/ent/employer"
	"griffin-dao/service"
)

func DeleteEmployer(employerGid string, ctx context.Context, client *ent.Client) error {
	service.PrintYellowStatus(fmt.Sprintf("Deleting employer with employer gid %s\n", employerGid))
	delNum, err := client.EMPLOYER.
		Delete().
		Where(employer.Gid(employerGid)).
		Exec(ctx)
	if recover() != nil || err != nil {
		service.PrintRedError("Delete request::employer delete with employerGid failed: ", err)
		return errors.New(DATABASE_DELETE_FAIL)
	}
	service.PrintGreenStatus("employer deleted: ", delNum)
	return nil
}

func DeleteEmployeewEmployerAll(employerGid string, ctx context.Context, client *ent.Client) error {
	service.PrintYellowStatus(fmt.Sprintf("Deleting all employee with employer gid %s\n", employerGid))
	delNum, err := client.EMPLOYEE.
		Delete().
		Where(
			employee.HasEmployeeFromEmployerWith(
				employer.GidEQ(employerGid),
			),
		).
		Exec(ctx)
	if recover() != nil || err != nil {
		service.PrintRedError("Delete request::employee with employerGid failed: ", err)
		return errors.New(DATABASE_DELETE_FAIL)
	}
	err = DeleteEmployer(employerGid, ctx, client)
	if err != nil {
		return err
	}
	service.PrintGreenStatus("employee deleted: ", delNum)
	return nil
}

func DeleteEmployeewEmployerInd(employerGid, employeeGid string, ctx context.Context, client *ent.Client) error {
	service.PrintYellowStatus(fmt.Sprintf("Deleting employee %s with employer gid %s\n", employeeGid, employerGid))
	delNum, err := client.EMPLOYEE.
		Delete().
		Where(
			employee.HasEmployeeFromEmployerWith(
				employer.Gid(employerGid),
			),
			employee.Gid(employeeGid),
		).
		Exec(ctx)
	if err != nil {
		service.PrintRedError("employer delete with employerGid failed: ", err)
		return errors.New(DATABASE_DELETE_FAIL)
	}
	service.PrintGreenStatus("employer deleted: ", delNum)
	return nil
}
