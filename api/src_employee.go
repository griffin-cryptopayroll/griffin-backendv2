package api

import (
	"context"
	"github.com/gin-gonic/gin"
	"griffin-dao/gcrud"
	"net/http"
	"os"
	"strconv"
	"time"
)

func addEmployee(c *gin.Context, db gcrud.GriffinWeb2Conn) {
	args := map[string]bool{
		EMPLOYEE_LNAME:    true,
		EMPLOYEE_FNAME:    true,
		EMPLOYEE_POSITION: true,
		EMPLOYEE_WALLET:   true,
		EMPLOYEE_PAYROLL:  false,
		EMPLOYEE_CURRENCY: false,
		EMPLOYEE_EMAIL:    true,
		EMPLOYEE_PAYDAY:   true,
		EMPLOYEE_WORKFOR:  true,
	}
	argsQuery, err := handleOptionalQueryParam(c, args)
	if err != nil {
		return
	}

	payroll, err := strconv.ParseFloat(argsQuery[EMPLOYEE_PAYROLL], 64)
	if args[EMPLOYEE_PAYROLL] && err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": REQUEST_WRONG_TYPE + " " + argsQuery[EMPLOYEE_PAYROLL],
		})
	}
	currency, err := strconv.Atoi(argsQuery[EMPLOYEE_CURRENCY])
	if args[EMPLOYEE_CURRENCY] && err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": REQUEST_WRONG_TYPE + " " + argsQuery[EMPLOYEE_CURRENCY],
		})
	}
	employType, err := strconv.Atoi(argsQuery[EMPLOYEE_TYPE])
	if args[EMPLOYEE_TYPE] && err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": REQUEST_WRONG_TYPE + " " + argsQuery[EMPLOYEE_TYPE],
		})
	}
	payday, err := time.Parse("20060102", argsQuery[EMPLOYEE_PAYDAY])
	if args[EMPLOYEE_PAYDAY] && err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": REQUEST_WRONG_TYPE + " " + argsQuery[EMPLOYEE_PAYDAY],
		})
	}

	freshMeet := gcrud.EmployeeJson{
		GriffinID:         argsQuery[EMPLOYEE_GID],
		EmployerGriffinID: argsQuery[EMPLOYEE_WORKFOR],
		LastName:          argsQuery[EMPLOYEE_LNAME],
		FirstName:         argsQuery[EMPLOYEE_FNAME],
		Position:          argsQuery[EMPLOYEE_POSITION],
		Wallet:            argsQuery[EMPLOYEE_WALLET],
		Payroll:           payroll,
		Currency:          currency,
		PayDay:            payday,
		EmployType:        employType,
		Email:             argsQuery[EMPLOYEE_EMAIL],
		CreatedAt:         time.Now(),
		CreatedBy:         os.Getenv("UPDATER"),
		UpdatedAt:         time.Now(),
		UpdatedBy:         os.Getenv("UPDATER"),
	}
	gcrud.CreateEmployee(freshMeet, context.Background(), db.Conn)
	c.JSON(http.StatusOK, gin.H{
		"message": DATABASE_CREATE_SUCCESS,
	})
}

func delEmployee(c *gin.Context, db gcrud.GriffinWeb2Conn) {
	args := map[string]bool{
		EMPLOYEE_GID:     true,
		EMPLOYEE_WORKFOR: true,
	}
	argsQuery, err := handleOptionalQueryParam(c, args)
	if err != nil {
		return
	}
	gcrud.DeleteEmployeewEmployerInd(argsQuery[EMPLOYEE_WORKFOR], argsQuery[EMPLOYEE_GID], context.Background(), db.Conn)
	c.JSON(http.StatusOK, gin.H{
		"message": DATABASE_DELETE_SUCCESS,
	})
}
