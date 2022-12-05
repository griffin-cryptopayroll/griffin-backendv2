package gcrud

import (
	"context"
	"errors"
	"fmt"
	"griffin-dao/ent"
	"griffin-dao/ent/employee"
	"griffin-dao/service"
	"math"
	"os"
	"strings"
	"time"
)

func UpdateNextEmployeePayday(employeeGid string, ctx context.Context, client ent.Client) error {
	service.PrintYellowStatus(
		fmt.Sprintf("Payment Completed. Updating employee Payday with employt type info"))
	obj, err := client.EMPLOYEE.
		Query().
		Where(
			employee.Gid(employeeGid),
		).
		WithEmployeeFromEmployType().
		Only(ctx)
	if recover() != nil || err != nil {
		service.PrintRedError(err)
		return errors.New(DATABASE_SELECT_FAIL)
	}
	newPayDay := time.Time{}
	payday := obj.Payday
	interval := strings.ToUpper(obj.Edges.EmployeeFromEmployType.PayFreq)
	switch interval {
	case "D":
		newPayDay = payday.Add(time.Hour * 24)
	case "W":
		newPayDay = payday.Add(time.Hour * 24 * 7)
	case "M":
		// extract day from payday
		payYear := payday.Year()
		payMonth := int(payday.Month())
		payDay := payday.Day()
		switch {
		case monthContain(payMonth, []int{1, 3, 5, 7, 8, 10, 12}):
			// Month Contains 31 dates
			MONTHDATE := 31

			newYear := payYear
			newday := int(math.Min(float64(payDay), float64(MONTHDATE)))
			newMonth := (payMonth + 1) % 13
			if newMonth == 0 {
				newMonth = 1
				newYear += 1
			}
			newPayDayStr := fmt.Sprintf("%v-%v-%v", newYear, newMonth, newday)
			newPayDay, err = time.Parse("2006-01-02", newPayDayStr)
			if err != nil {
				return err
			}
		case monthContain(payMonth, []int{4, 6, 9, 11}):
			// Month Contains 30 dates
			MONTHDATE := 30

			newYear := payYear
			newday := int(math.Min(float64(payDay), float64(MONTHDATE)))
			newMonth := (payMonth + 1) % 13
			if newMonth == 0 {
				newMonth = 1
				newYear += 1
			}
			newPayDayStr := fmt.Sprintf("%v-%v-%v", newYear, newMonth, newday)
			newPayDay, err = time.Parse("2006-01-02", newPayDayStr)
			if err != nil {
				return err
			}
		case monthContain(payMonth, []int{2}):
			// Month Contains 28 dates
			MONTHDATE := 28

			newYear := payYear
			newday := int(math.Min(float64(payDay), float64(MONTHDATE)))
			newMonth := (payMonth + 1) % 13
			if newMonth == 0 {
				newMonth = 1
				newYear += 1
			}
			newPayDayStr := fmt.Sprintf("%v-%v-%v", newYear, newMonth, newday)
			newPayDay, err = time.Parse("2006-01-02", newPayDayStr)
			if err != nil {
				return err
			}
		}
	}

	err = client.EMPLOYEE.
		Update().
		SetPayday(newPayDay).
		SetUpdatedAt(time.Now()).
		SetUpdatedBy(os.Getenv("UPDATER")).
		Exec(ctx)
	return err
}

func monthContain[T comparable](eval T, evalFrom []T) bool {
	for _, v := range evalFrom {
		if v == eval {
			return true
		}
	}
	return false
}
