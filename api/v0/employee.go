package v0

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

// GenerateEmployee
// @Summary Post employee to the database.
// @Description Worker's information needed.
// @Accept  json
// @Produce  json
// @Param name query string true "Full name, since crypto lovers don't use their original name"
// @Param employ_type query string true "permanent or freelance"
// @Param pay_freq query string true "D, W. M(Not yet implemented)"
// @Param position query string false "Position ex: Backend engineer, Frontend engineer"
// @Param wallet query string true "Employee's information. His or her payment wallet address"
// @Param payroll query float32 true "Payroll amount in float"
// @Param currency query string true "ETHUSDT, MATICUSDT or USDCUSDT. Capitalization required"
// @Param email query string true "Employee's information. Corp or organization's em"
// @Param payday query time.Time true "Employee's information. Payday information"
// @Param employer_gid query string true "Employee's information. Corp Gid or Organization Gid"
// @Param work_start query string true "Employee's information. When does he or she starts work. In YYYYMMDD"
// @Param work_end query string false "Employee's information. When does he or she ends work. Required if freelance. In YYYYMMDD"
// @Router /employee [post]
// @Success 200 {object} CommonResponse
// @Failure 400 {object} CommonResponse
// @Failure 500 {object} CommonResponse
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
			"status":  false,
			"message": REQUEST_WRONG_TYPE + " " + argsQuery[EMPLOYEE_PAYROLL],
		})
	}
	err = common.ValidateEmployType(argsQuery[EMPLOYEE_TYPE])
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": REQUEST_WRONG_TYPE + " " + argsQuery[EMPLOYEE_TYPE],
		})
	}
	err = common.ValidatePayFreq(argsQuery[EMP_PAY_FREQ])
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": REQUEST_WRONG_TYPE + " " + argsQuery[EMP_PAY_FREQ],
		})
	}
	payday, err := time.Parse("20060102", argsQuery[EMPLOYEE_PAYDAY])
	if args[EMPLOYEE_PAYDAY] && err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
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

	obj, err := gcrud.CreateEmployee(
		freshMeat,
		argsQuery[EMPLOYEE_WORKFOR],
		argsQuery[EMPLOYEE_CURRENCY],
		argsQuery[EMPLOYEE_TYPE],
		argsQuery[EMP_PAY_FREQ],
		ctx, db.Conn,
	)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": DATABASE_CREATE_FAIL,
		})
		return
	}
	switch argsQuery[EMPLOYEE_TYPE] {
	case "permanent":
		// Generate 1 year
		ws, err := time.Parse("20060102", argsQuery[EMPLOYEE_WORKSTART])
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": REQUEST_WRONG_TYPE,
			})
			return
		}
		err = gcrud.CreatePermanent(obj, ws, argsQuery[EMP_PAY_FREQ], ctx, db.Conn)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status":  false,
				"message": "Permanent worker payment schedule creation failed.",
			})
			return
		}
	case "freelance":
		// Generate from WorkStart ~ WorkEnds
		ws, err1 := time.Parse("20060102", argsQuery[EMPLOYEE_WORKSTART])
		we, err2 := time.Parse("20060102", argsQuery[EMPLOYEE_WORKEND])
		if err1 != nil || err2 != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": REQUEST_WRONG_TYPE,
			})
			return
		}
		err = gcrud.CreateFreelance(obj, ws, we, argsQuery[EMP_PAY_FREQ], ctx, db.Conn)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status":  false,
				"message": "Freelance worker payment schedule creation failed.",
			})
			return
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": DATABASE_CREATE_SUCCESS,
	})
}

// RemoveEmployee
// @Summary Delete employee from the database.
// @Description Worker's information needed. His/Her Griffin ID (GID) and employer Griffin ID.
// @Accept  json
// @Produce  json
// @Param gid query string true "Employee's griffin id (in uuid form)"
// @Param employer_gid query string true "Employee's information. Corp Gid or Organization Gid"
// @Router /employee [delete]
// @Success 200 {object} CommonResponse
// @Failure 400 {object} CommonResponse
// @Failure 500 {object} CommonResponse
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
			"status":  false,
			"message": DATABASE_DELETE_FAIL,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": DATABASE_DELETE_SUCCESS,
	})
}

// EmployeeSingle
// @Summary Query employee from the database.
// @Description Worker's information needed. Worker is singled out with their griffin id and his employer id.
// @Accept  json
// @Produce  json
// @Param gid query string true "Employee's griffin id (in uuid form)"
// @Param employer_gid query string true "Employee's information. Corp Gid or Organization Gid"
// @Router /employee/single [get]
// @Success 200 {object} ent.EMPLOYEE
// @Failure 400 {object} CommonResponse
// @Failure 500 {object} CommonResponse
func EmployeeSingle(c *gin.Context, db gcrud.GriffinWeb2Conn) {
	var ctx = context.Background()
	args := map[string]bool{
		EMPLOYEE_GID:     true,
		EMPLOYEE_WORKFOR: true,
	}
	argsQuery, err := handleOptionalQueryParam(c, args)
	if err != nil {
		return
	}

	result, err := gcrud.QueryEmployee(argsQuery[EMPLOYEE_GID], argsQuery[EMPLOYEE_WORKFOR], ctx, db.Conn)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": DATABASE_SELECT_FAIL,
		})
		return
	}
	c.JSON(http.StatusOK, result)
}

// EmployeeMulti
// @Summary Query employee from the database.
// @Description Worker's information needed.
// @Accept  json
// @Produce  json
// @Param employer_gid query string true "Employee's information. Corp Gid or Organization Gid"
// @Router /employee/multi [get]
// @Success 200 {object} []ent.EMPLOYEE
// @Failure 400 {object} CommonResponse
// @Failure 500 {object} CommonResponse
func EmployeeMulti(c *gin.Context, db gcrud.GriffinWeb2Conn) {
	var ctx = context.Background()
	args := map[string]bool{
		EMPLOYEE_WORKFOR: true,
	}
	argsQuery, err := handleOptionalQueryParam(c, args)
	if err != nil {
		return
	}

	results, err := gcrud.QueryEmployeewEmployerGid(argsQuery[EMPLOYEE_WORKFOR], ctx, db.Conn)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": DATABASE_SELECT_FAIL,
		})
	}
	c.JSON(http.StatusOK, results)
}
