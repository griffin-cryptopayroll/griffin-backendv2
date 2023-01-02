package v0

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"griffin-dao/gcrud"
	"net/http"
	"time"
)

// PaymentExecute
// @Summary Create payment log for employee. Scheduled payment is executed
// @Description Employee is identified by employer and employee Gid.
// @Description After update recommend updating payment date for employee
// @Accept json
// @Produce json
// @Param gid query string true "Employee's griffin id (in uuid form)"
// @Param employer_gid query string true "Employee's information. Corp Gid or Organization Gid"
// @Param schd_date query string true "Payment schedule date. `gid` `employer_gid` and `schd_date` will make a unique key. Format should be YYYYMMDD"
// @Param exec_date query string true "Payment executed date. In format YYYYMMDD"
// @Router /payment/execute [put]
// @Success 200 {object} CommonResponse
// @Failure 400 {object} CommonResponse
// @Failure 500 {object} CommonResponse
func PaymentExecute(c *gin.Context, db gcrud.GriffinWeb2Conn) {
	var ctx = context.Background()
	args := map[string]bool{
		EMPLOYEE_GID:      true,
		EMPLOYEE_WORKFOR:  true,
		SCHEDULED_PAYMENT: true,
		EXECUTED_PAYMENT:  true,
	}
	argsQuery, err := handleOptionalQueryParam(c, args)
	if err != nil {
		return
	}

	// Payment executed to `executeTo` employee
	executeTo, err := gcrud.QueryEmployee(
		argsQuery[EMPLOYEE_GID],
		argsQuery[EMPLOYEE_WORKFOR],
		ctx,
		db.Conn,
	)
	// Parse Date
	schd, err1 := time.Parse("20060102", argsQuery[SCHEDULED_PAYMENT])
	exec, err2 := time.Parse("20060102", argsQuery[EXECUTED_PAYMENT])
	if err1 != nil || err2 != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": false,
			"message": fmt.Sprintf(
				"%s, %s and %s should be YYYYMMDD form",
				REQUEST_WRONG_TYPE,
				argsQuery[SCHEDULED_PAYMENT],
				argsQuery[EXECUTED_PAYMENT],
			),
		})
		return
	}
	err = gcrud.UpdatePaymentExecuted(executeTo, schd, exec, ctx, db.Conn)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  false,
			"message": err.Error(),
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": DATABASE_UPDATE_SUCCESS,
	})
}

// PaymentOneOff
// @Summary Create payment log for employee. Oneoff payment. Not scheduled.
// @Description Employee is identified by employer and employee Gid.
// @Description After update recommend updating payment date for employee
// @Accept json
// @Produce json
// @Param gid query string true "Employee's griffin id (in uuid form)"
// @Param employer_gid query string true "Employer's information. Corp Gid or Organization Gid"
// @Param onof_date query string true "Payment executed date. In format YYYYMMDD"
// @Router /payment/oneoff [post]
// @Success 200 {object} CommonResponse
// @Failure 400 {object} CommonResponse
// @Failure 500 {object} CommonResponse
func PaymentOneOff(c *gin.Context, db gcrud.GriffinWeb2Conn) {
	var ctx = context.Background()
	args := map[string]bool{
		EMPLOYEE_GID:     true,
		EMPLOYEE_WORKFOR: true,
		ONEOFF_PAYMENT:   true,
	}
	argsQuery, err := handleOptionalQueryParam(c, args)
	if err != nil {
		return
	}

	// Payment executed to `executeTo` employee
	executeTo, err := gcrud.QueryEmployee(
		argsQuery[EMPLOYEE_GID],
		argsQuery[EMPLOYEE_WORKFOR],
		ctx,
		db.Conn,
	)
	// ParseDate
	oneOff, err := time.Parse("20060102", argsQuery[ONEOFF_PAYMENT])
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": false,
			"message": fmt.Sprintf(
				"%s, %s should be YYYYMMDD form",
				REQUEST_WRONG_TYPE,
				argsQuery[ONEOFF_PAYMENT],
			),
		})
		return
	}
	err = gcrud.CreatePaymentOneOff(executeTo, oneOff, ctx, db.Conn)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  false,
			"message": err.Error(),
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": DATABASE_CREATE_SUCCESS,
	})
}

// PaymentByEmployee
// @Summary Query payment log for employee.
// @Description Employee is identified by employer and employee Gid.
// @Description Gives you 1. scheduled payment, 2. scheduled and executed payment, and 3. oneoff payment
// @Accept json
// @Produce json
// @Param gid query string true "Employee's griffin id (in uuid form)"
// @Param employer_gid query string true "Employer's griffin id (in uuid form)"
// @Router /payment/employee [get]
// @Success 200 {object} PaymentType
// @Failure 400 {object} CommonResponse
func PaymentByEmployee(c *gin.Context, db gcrud.GriffinWeb2Conn) {
	var ctx = context.Background()
	args := map[string]bool{
		EMPLOYEE_GID:     true,
		EMPLOYEE_WORKFOR: true,
	}
	argsQuery, err := handleOptionalQueryParam(c, args)
	if err != nil {
		return
	}
	payments, err := gcrud.QueryPaymentEmployee(argsQuery[EMPLOYEE_GID], ctx, db.Conn)
	res := PaymentType{
		Scheduled: gcrud.PaymentScheduled(payments),
		Executed:  gcrud.PaymentExecuted(payments),
		OneOff:    gcrud.PaymentOneOff(payments),
	}
	c.JSON(http.StatusOK, res)
}

// PaymentByEmployer
// @Summary Query payment log for employer
// @Description Gives you 1. scheduled payment, 2. scheduled and executed payment, and 3. oneoff payment
// @Description for whole employee cohort working under employer
// @Accept json
// @Produce json
// @Param employer_gid query string true "Employee's information. Corp Gid or Organization Gid"
// @Router /payment/employer [get]
// @Success 200 {object} PaymentType
// @Failure 400 {object} CommonResponse
func PaymentByEmployer(c *gin.Context, db gcrud.GriffinWeb2Conn) {
	var ctx = context.Background()
	args := map[string]bool{
		EMPLOYEE_WORKFOR: true,
	}
	argsQuery, err := handleOptionalQueryParam(c, args)
	if err != nil {
		return
	}
	payments, err := gcrud.QueryPaymentEmployer(argsQuery[EMPLOYEE_WORKFOR], ctx, db.Conn)
	res := PaymentType{
		Scheduled: gcrud.PaymentScheduled(payments),
		Executed:  gcrud.PaymentExecuted(payments),
		OneOff:    gcrud.PaymentOneOff(payments),
	}
	c.JSON(http.StatusOK, res)
}

// PaymentFuture
// @Summary Query future scheduled payment for an employer.
// @Description Gives you future scheduled payment. Future determined by interval
// @Accept json
// @Produce json
// @Param employer_gid query string true "Employee's information. Corp Gid or Organization Gid"
// @Param interval query string true "time interval. supports 2 length string. (O) 1d, 2m, 3y | (X) 10d, 20m, 30y"
// @Router /payment/future [get]
// @Success 200 {object} PaymentTime
// @Failure 400 {object} CommonResponse
func PaymentFuture(c *gin.Context, db gcrud.GriffinWeb2Conn) {
	var ctx = context.Background()
	args := map[string]bool{
		EMPLOYEE_WORKFOR: true,
		INTERVAL:         true,
	}
	argsQuery, err := handleOptionalQueryParam(c, args)
	if err != nil {
		return
	}
	payments, err := gcrud.QueryPaymentEmployer(argsQuery[EMPLOYEE_WORKFOR], ctx, db.Conn)
	f, err := gcrud.PaymentFuture(payments, argsQuery[INTERVAL])
	if err != nil {
		msg := CommonResponse{
			Status:  false,
			Message: err.Error(),
		}
		c.JSON(http.StatusBadRequest, msg)
	}
	res := PaymentTime{
		Future: f,
	}
	c.JSON(http.StatusOK, res)
}

// PaymentPast
// @Summary Query past scheduled payment for an employer.
// @Description Gives you past executed payment. Future determined by interval
// @Accept json
// @Produce json
// @Param employer_gid query string true "Employee's information. Corp Gid or Organization Gid"
// @Param interval query string true "time interval. supports 2 length string. (O) 1d, 2m, 3y | (X) 10d, 20m, 30y"
// @Router /payment/past [get]
// @Success 200 {object} PaymentTime
// @Failure 400 {object} CommonResponse
func PaymentPast(c *gin.Context, db gcrud.GriffinWeb2Conn) {
	var ctx = context.Background()
	args := map[string]bool{
		EMPLOYEE_WORKFOR: true,
		INTERVAL:         true,
	}
	argsQuery, err := handleOptionalQueryParam(c, args)
	if err != nil {
		return
	}
	payments, err := gcrud.QueryPaymentEmployer(argsQuery[EMPLOYEE_WORKFOR], ctx, db.Conn)
	p, err := gcrud.PaymentPast(payments, argsQuery[INTERVAL])
	if err != nil {
		msg := CommonResponse{
			Status:  false,
			Message: err.Error(),
		}
		c.JSON(http.StatusBadRequest, msg)
	}
	res := PaymentTime{
		Past: p,
	}
	c.JSON(http.StatusOK, res)
}

// PaymentMissed
// @Summary Query missed scheduled payment for an employer.
// @Description Gives you missed scheduled payment. NO interval needed
// @Accept json
// @Produce json
// @Param employer_gid query string true "Employee's information. Corp Gid or Organization Gid"
// @Router /payment/miss [get]
// @Success 200 {object} PaymentTime
// @Failure 400 {object} CommonResponse
func PaymentMissed(c *gin.Context, db gcrud.GriffinWeb2Conn) {
	var ctx = context.Background()
	args := map[string]bool{
		EMPLOYEE_WORKFOR: true,
	}
	argsQuery, err := handleOptionalQueryParam(c, args)
	if err != nil {
		return
	}
	payments, err := gcrud.QueryPaymentEmployer(argsQuery[EMPLOYEE_WORKFOR], ctx, db.Conn)
	p := gcrud.PaymentMissed(payments)
	res := PaymentTime{
		Missed: p,
	}
	c.JSON(http.StatusOK, res)
}
