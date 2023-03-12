package api_v0

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	api_base "griffin-dao/api/base"
	"griffin-dao/api/common"
	"griffin-dao/dao"
	"griffin-dao/ent"

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
// @Param pay_freq query string true "D, W. M"
// @Param position query string false "Position ex: Backend engineer, Frontend engineer"
// @Param wallet query string true "Employee's information. His or her payment wallet address"
// @Param wallet_aztec query string true "Employee's information. His or her payment wallet address"
// @Param payroll query float32 true "Payroll amount in float"
// @Param currency query string true "ETH, MATIC or USDC. Capitalization required"
// @Param email query string true "Employee's information. Corp or organization's em"
// @Param payday query time.Time true "Employee's information. Payday information"
// @Param employer_gid query string true "Employee's information. Corp Gid or Organization Gid"
// @Param work_start query string true "Employee's information. When does he or she starts work. In YYYYMMDD"
// @Param work_end query string false "Employee's information. When does he or she ends work. Required if freelance. In YYYYMMDD"
// @Router /api/v0/employee [post]
// @Success 200 {object} api_base.CommonResponse
// @Failure 400 {object} api_base.CommonResponse
// @Failure 500 {object} api_base.CommonResponse
func GenerateEmployee(c *gin.Context, db dao.GriffinWeb2Conn) {
	args := map[string]bool{
		EMPLOYEE_NAME:         true,
		EMPLOYEE_TYPE:         true,
		EMP_PAY_FREQ:          true,
		EMPLOYEE_POSITION:     false,
		EMPLOYEE_WALLET:       true,
		EMPLOYEE_WALLET_AZTEC: true,
		EMPLOYEE_PAYROLL:      true,
		EMPLOYEE_CURRENCY:     true,
		EMPLOYEE_EMAIL:        true,
		EMPLOYEE_PAYDAY:       true,
		EMPLOYEE_WORKFOR:      true,
		EMPLOYEE_WORKSTART:    true,
		EMPLOYEE_WORKEND:      false,
	}
	argsQuery, err := common.HandleOptionalQueryParam(c, args)
	if err != nil {
		return
	}

	payroll, err := strconv.ParseFloat(argsQuery[EMPLOYEE_PAYROLL], 64)
	if args[EMPLOYEE_PAYROLL] && err != nil {
		msg := api_base.CommonResponse{
			Status:  false,
			Message: REQUEST_WRONG_TYPE + " " + argsQuery[EMPLOYEE_PAYROLL],
		}
		c.JSON(http.StatusBadRequest, msg)
		return
	}
	// Validate whether is it {permanent or freelance}
	err = common.ValidateEmployType(argsQuery[EMPLOYEE_TYPE])
	if err != nil {
		msg := api_base.CommonResponse{
			Status:  false,
			Message: REQUEST_WRONG_TYPE + " " + argsQuery[EMPLOYEE_TYPE],
		}
		c.JSON(http.StatusBadRequest, msg)
		return
	}
	// Validate whether is it "D", "W", or "M". Capitalization required
	err = common.ValidatePayFreq(argsQuery[EMP_PAY_FREQ])
	if err != nil {
		msg := api_base.CommonResponse{
			Status:  false,
			Message: REQUEST_WRONG_TYPE + " " + argsQuery[EMP_PAY_FREQ],
		}
		c.JSON(http.StatusBadRequest, msg)
		return
	}
	payday, err := time.Parse("20060102", argsQuery[EMPLOYEE_PAYDAY])
	if args[EMPLOYEE_PAYDAY] && err != nil {
		msg := api_base.CommonResponse{
			Status:  false,
			Message: REQUEST_WRONG_TYPE + " " + argsQuery[EMPLOYEE_PAYDAY],
		}
		c.JSON(http.StatusBadRequest, msg)
		return
	}
	// If `pay_freq` is `M`, date should be leq 25th
	err = common.ValidateMonthDate(payday, argsQuery[EMP_PAY_FREQ])

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

	obj, err := dao.CreateEmployee(
		freshMeat,
		argsQuery[EMPLOYEE_WORKFOR],
		argsQuery[EMPLOYEE_CURRENCY],
		argsQuery[EMPLOYEE_TYPE],
		argsQuery[EMP_PAY_FREQ],
		ctx,
		db.Conn,
	)
	if err != nil {
		msg := api_base.CommonResponse{
			Status:  false,
			Message: DATABASE_CREATE_FAIL,
		}
		c.JSON(http.StatusBadRequest, msg)
		return
	}
	switch argsQuery[EMPLOYEE_TYPE] {
	case "permanent":
		// Generate 1 year
		ws, err := time.Parse("20060102", argsQuery[EMPLOYEE_WORKSTART])
		if err != nil {
			msg := api_base.CommonResponse{
				Status:  false,
				Message: REQUEST_WRONG_TYPE,
			}
			c.JSON(http.StatusBadRequest, msg)
			return
		}
		err = dao.CreatePermanent(obj, ws, argsQuery[EMP_PAY_FREQ], ctx, db.Conn)
		if err != nil {
			msg := api_base.CommonResponse{
				Status:  false,
				Message: "Permanent worker payment schedule creation failed.",
			}
			c.JSON(http.StatusInternalServerError, msg)
			return
		}
	case "freelance":
		// Generate from WorkStart ~ WorkEnds
		ws, err1 := time.Parse("20060102", argsQuery[EMPLOYEE_WORKSTART])
		we, err2 := time.Parse("20060102", argsQuery[EMPLOYEE_WORKEND])
		if err1 != nil || err2 != nil {
			msg := api_base.CommonResponse{
				Status: false,
				Message: fmt.Sprintf(
					"%s. Error1: %v, Error2: %v",
					REQUEST_WRONG_TYPE,
					err1,
					err2,
				),
			}
			c.JSON(http.StatusBadRequest, msg)
			return
		}
		err = dao.CreateFreelance(obj, ws, we, argsQuery[EMP_PAY_FREQ], ctx, db.Conn)
		if err != nil {
			msg := api_base.CommonResponse{
				Status:  false,
				Message: "Freelance worker payment schedule creation failed.",
			}
			c.JSON(http.StatusInternalServerError, msg)
			return
		}
	}

	msg := api_base.CommonResponse{
		Status:  true,
		Message: DATABASE_CREATE_SUCCESS,
	}
	c.JSON(http.StatusOK, msg)
}

// RemoveEmployee
// @Summary Delete employee from the database.
// @Description Worker's information needed. His/Her Griffin ID (GID) and employer Griffin ID.
// @Accept  json
// @Produce  json
// @Param gid query string true "Employee's griffin id (in uuid form)"
// @Param employer_gid query string true "Employee's information. Corp Gid or Organization Gid"
// @Router /api/v0/employee [delete]
// @Success 200 {object} api_base.CommonResponse
// @Failure 400 {object} api_base.CommonResponse
// @Failure 500 {object} api_base.CommonResponse
func RemoveEmployee(c *gin.Context, db dao.GriffinWeb2Conn) {
	var ctx = context.Background()
	args := map[string]bool{
		EMPLOYEE_GID:     true,
		EMPLOYEE_WORKFOR: true,
	}
	argsQuery, err := common.HandleOptionalQueryParam(c, args)
	if err != nil {
		return
	}

	err = dao.DeleteEmployeewEmployerInd(
		argsQuery[EMPLOYEE_WORKFOR],
		argsQuery[EMPLOYEE_GID],
		ctx,
		db.Conn,
	)
	if err != nil {
		msg := api_base.CommonResponse{
			Status:  false,
			Message: DATABASE_DELETE_FAIL,
		}
		c.JSON(http.StatusBadRequest, msg)
		return
	}

	msg := api_base.CommonResponse{
		Status:  true,
		Message: DATABASE_DELETE_SUCCESS,
	}
	c.JSON(http.StatusOK, msg)
}

// EmployeeSingle
// @Summary Query employee from the database.
// @Description Worker's information needed. Worker is singled out with their griffin id and his employer id.
// @Accept  json
// @Produce  json
// @Param gid query string true "Employee's griffin id (in uuid form)"
// @Param employer_gid query string true "Employee's information. Corp Gid or Organization Gid"
// @Router /api/v0/employee/single [get]
// @Success 200 {object} ent.EMPLOYEE
// @Failure 400 {object} api_base.CommonResponse
// @Failure 500 {object} api_base.CommonResponse
func EmployeeSingle(c *gin.Context, db dao.GriffinWeb2Conn) {
	var ctx = context.Background()
	args := map[string]bool{
		EMPLOYEE_GID:     true,
		EMPLOYEE_WORKFOR: true,
	}
	argsQuery, err := common.HandleOptionalQueryParam(c, args)
	if err != nil {
		return
	}

	result, err := dao.QueryEmployee(argsQuery[EMPLOYEE_GID], argsQuery[EMPLOYEE_WORKFOR], ctx, db.Conn)
	if err != nil {
		msg := api_base.CommonResponse{
			Status:  false,
			Message: DATABASE_SELECT_FAIL,
		}
		c.JSON(http.StatusBadRequest, msg)
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
// @Router /api/v0/employee/multi [get]
// @Success 200 {object} []ent.EMPLOYEE
// @Failure 400 {object} api_base.CommonResponse
// @Failure 500 {object} api_base.CommonResponse
func EmployeeMulti(c *gin.Context, db dao.GriffinWeb2Conn) {
	var ctx = context.Background()
	args := map[string]bool{
		EMPLOYEE_WORKFOR: true,
	}
	argsQuery, err := common.HandleOptionalQueryParam(c, args)
	if err != nil {
		return
	}

	results, err := dao.QueryEmployeewEmployerGid(argsQuery[EMPLOYEE_WORKFOR], ctx, db.Conn)
	if err != nil {
		msg := api_base.CommonResponse{
			Status:  false,
			Message: DATABASE_SELECT_FAIL,
		}
		c.JSON(http.StatusBadRequest, msg)
	}
	c.JSON(http.StatusOK, results)
}
