package v1

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"griffin-dao/api/common"
	"griffin-dao/gcrud"

	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func addEmployee(c *gin.Context, db gcrud.GriffinWeb2Conn) {
	args := map[string]bool{
		EMPLOYEE_NAME:      true,
		EMPLOYEE_TYPE:      true,
		EMP_PAY_FREQ:       true,
		EMPLOYEE_POSITION:  false,
		EMPLOYEE_WALLET:    true,
		EMPLOYEE_PAYROLL:   true,
		EMPLOYEE_CURRENCY:  true,
		EMPLOYEE_EMAIL:     true,
		EMPLOYEE_PAYDAY:    true,
		EMPLOYEE_WORKFOR:   true,
		EMPLOYEE_WORKSTART: true,
		EMPLOYEE_WORKEND:   false,
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
	err = common.ValidateEmployType(argsQuery[EMPLOYEE_TYPE])
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": REQUEST_WRONG_TYPE + " " + argsQuery[EMPLOYEE_TYPE],
		})
	}
	err = common.ValidatePayFreq(argsQuery[EMP_PAY_FREQ])
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": REQUEST_WRONG_TYPE + " " + argsQuery[EMP_PAY_FREQ],
		})
	}
	payday, err := time.Parse("20060102", argsQuery[EMPLOYEE_PAYDAY])
	if args[EMPLOYEE_PAYDAY] && err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": REQUEST_WRONG_TYPE + " " + argsQuery[EMPLOYEE_PAYDAY],
		})
	}

	// Currency query, EmployType query
	currency, err := gcrud.QueryCurrency(argsQuery[EMPLOYEE_CURRENCY], context.Background(), db.Conn)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
	}

	gid := uuid.New()
	freshMeet := gcrud.EmployeeJson{
		GriffinID:         gid.String(),
		EmployerGriffinID: argsQuery[EMPLOYEE_WORKFOR],
		Name:              argsQuery[EMPLOYEE_NAME],
		Position:          argsQuery[EMPLOYEE_POSITION],
		Wallet:            argsQuery[EMPLOYEE_WALLET],
		Payroll:           payroll,
		Currency:          currency.ID,
		PayDay:            payday,
		EmployType:        argsQuery[EMPLOYEE_TYPE],
		Email:             argsQuery[EMPLOYEE_EMAIL],
		WorkStart:         argsQuery[EMPLOYEE_WORKSTART],
		WorkEnd:           argsQuery[EMPLOYEE_WORKEND],
		CreatedAt:         time.Now(),
		CreatedBy:         os.Getenv("UPDATER"),
		UpdatedAt:         time.Now(),
		UpdatedBy:         os.Getenv("UPDATER"),
	}
	err = gcrud.CreateEmployee(
		freshMeet,
		argsQuery[EMPLOYEE_TYPE], argsQuery[EMP_PAY_FREQ],
		context.Background(), db.Conn,
	)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": DATABASE_CREATE_FAIL,
		})
		return
	}
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
	err = gcrud.DeleteEmployeewEmployerInd(argsQuery[EMPLOYEE_WORKFOR], argsQuery[EMPLOYEE_GID], context.Background(), db.Conn)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": DATABASE_DELETE_FAIL,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": DATABASE_DELETE_SUCCESS,
	})
}

func getEmployeeSingle(c *gin.Context, db gcrud.GriffinWeb2Conn) {
	args := map[string]bool{
		EMPLOYEE_GID:     true,
		EMPLOYEE_WORKFOR: true,
	}
	argsQuery, err := handleOptionalQueryParam(c, args)
	if err != nil {
		return
	}

	result, err := gcrud.QueryEmployee(argsQuery[EMPLOYEE_GID], argsQuery[EMPLOYEE_WORKFOR], context.Background(), db.Conn)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": DATABASE_SELECT_FAIL,
		})
	}

	meet := gcrud.EmployeeJson{
		GriffinID:  result.Gid,
		Name:       result.Name,
		Position:   result.Position,
		Wallet:     result.Wallet,
		Payroll:    result.Payroll,
		Currency:   result.Currency,
		PayDay:     result.Payday,
		EmployType: result.Edges.EmployeeTypeFrom.IsPermanent,
		Email:      result.Email,
	}
	fmt.Println(meet)
	c.JSON(http.StatusOK, meet)
}

func getEmployeeMulti(c *gin.Context, db gcrud.GriffinWeb2Conn) {
	args := map[string]bool{
		EMPLOYEE_WORKFOR: true,
	}
	argsQuery, err := handleOptionalQueryParam(c, args)
	if err != nil {
		return
	}

	results, err := gcrud.QueryEmployeewEmployerGid(argsQuery[EMPLOYEE_WORKFOR], context.Background(), db.Conn)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": DATABASE_SELECT_FAIL,
		})
	}

	var meets []gcrud.EmployeeJson
	for _, result := range results {
		meet := gcrud.EmployeeJson{
			GriffinID:  result.Gid,
			Name:       result.Name,
			Position:   result.Position,
			Wallet:     result.Wallet,
			Payroll:    result.Payroll,
			Currency:   result.Currency,
			PayDay:     result.Payday,
			EmployType: result.Edges.EmployeeTypeFrom.IsPermanent,
			Email:      result.Email,
		}
		meets = append(meets, meet)
	}
	c.JSON(http.StatusOK, meets)
}
