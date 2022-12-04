package v1

import (
	"context"
	"github.com/google/uuid"
	"griffin-dao/api/common"
	"griffin-dao/ent"
	"griffin-dao/gcrud"

	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func GenerateEmployee(c *gin.Context, db gcrud.GriffinWeb2Conn) {
	var ctx = context.Background()
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

	gid := uuid.New()
	freshMeat := ent.EMPLOYEE{
		Gid:       gid.String(),
		Name:      argsQuery[EMPLOYEE_NAME],
		Position:  argsQuery[EMPLOYEE_POSITION],
		Wallet:    argsQuery[EMPLOYEE_WALLET],
		Payroll:   payroll,
		Payday:    payday,
		Email:     argsQuery[EMPLOYEE_EMAIL],
		WorkStart: argsQuery[EMPLOYEE_WORKSTART],
		WorkEnds:  argsQuery[EMPLOYEE_WORKEND],
		CreatedBy: os.Getenv("UPDATER"),
		UpdatedBy: os.Getenv("UPDATER"),
	}

	err = gcrud.CreateEmployee(
		freshMeat,
		argsQuery[EMPLOYEE_WORKFOR],
		argsQuery[EMPLOYEE_CURRENCY],
		argsQuery[EMPLOYEE_TYPE],
		argsQuery[EMP_PAY_FREQ],
		ctx, db.Conn,
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

func RemoveEmployee(c *gin.Context, db gcrud.GriffinWeb2Conn) {
	var ctx = context.Background()
	args := map[string]bool{
		EMPLOYEE_GID:     true,
		EMPLOYEE_WORKFOR: true,
	}
	argsQuery, err := handleOptionalQueryParam(c, args)
	if err != nil {
		return
	}
	err = gcrud.DeleteEmployeewEmployerInd(
		argsQuery[EMPLOYEE_WORKFOR],
		argsQuery[EMPLOYEE_GID],
		ctx,
		db.Conn,
	)
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

func EmployeeSingle(c *gin.Context, db gcrud.GriffinWeb2Conn) {
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
	c.JSON(http.StatusOK, result)
}

func EmployeeMulti(c *gin.Context, db gcrud.GriffinWeb2Conn) {
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
	c.JSON(http.StatusOK, results)
}
