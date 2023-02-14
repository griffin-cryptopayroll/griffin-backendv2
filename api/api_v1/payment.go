package api_v1

import (
	"github.com/gin-gonic/gin"
	api_base "griffin-dao/api/base"
	"griffin-dao/api/common"
	"griffin-dao/dao"
	"griffin-dao/util"
	"net/http"
)

// PaymentByEmployer
// @Summary Query payment log for employer
// @Description Gives you 1. scheduled payment, 2. scheduled and executed payment, and 3. oneoff payment
// @Description for whole employee cohort working under employer
// @Accept json
// @Produce json
// @Param employer_gid query string true "Employee's information. Corp Gid or Organization Gid"
// @Router /api/v1/payment/employer [get]
// @Success 200 {object} api_base.PaymentType
// @Failure 400 {object} api_base.CommonResponse
func PaymentByEmployer(c *gin.Context, db dao.GriffinWeb2Conn) {
	args := map[string]bool{
		EMPLOYEE_WORKFOR: true,
	}
	argsQuery, err := common.HandleOptionalQueryParam(c, args)
	if err != nil {
		return
	}
	payments, err := dao.QueryPaymentEmployer(argsQuery[EMPLOYEE_WORKFOR], ctx, db.Conn)
	res := api_base.PaymentTypeV2{
		Scheduled: util.FlattenPayment(dao.PaymentScheduled(payments)),
		Executed:  util.FlattenPayment(dao.PaymentExecuted(payments)),
		OneOff:    util.FlattenPayment(dao.PaymentOneOff(payments)),
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
// @Router /api/v1/payment/future [get]
// @Success 200 {object} api_base.PaymentTimeV2
// @Failure 400 {object} api_base.CommonResponse
func PaymentFuture(c *gin.Context, db dao.GriffinWeb2Conn) {
	args := map[string]bool{
		EMPLOYEE_WORKFOR: true,
		INTERVAL:         true,
	}
	argsQuery, err := common.HandleOptionalQueryParam(c, args)
	if err != nil {
		return
	}
	payments, err := dao.QueryPaymentEmployer(argsQuery[EMPLOYEE_WORKFOR], ctx, db.Conn)
	f, err := dao.PaymentFuture(payments, argsQuery[INTERVAL])
	if err != nil {
		msg := api_base.CommonResponse{
			Status:  false,
			Message: err.Error(),
		}
		c.JSON(http.StatusBadRequest, msg)
		return
	}
	res := api_base.PaymentTimeV2{
		Future: util.FlattenPayment(f),
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
// @Router /api/v1/payment/past [get]
// @Success 200 {object} api_base.PaymentTimeV2
// @Failure 400 {object} api_base.CommonResponse
func PaymentPast(c *gin.Context, db dao.GriffinWeb2Conn) {
	args := map[string]bool{
		EMPLOYEE_WORKFOR: true,
		INTERVAL:         true,
	}
	argsQuery, err := common.HandleOptionalQueryParam(c, args)
	if err != nil {
		return
	}
	payments, err := dao.QueryPaymentEmployer(argsQuery[EMPLOYEE_WORKFOR], ctx, db.Conn)
	p, err := dao.PaymentPast(payments, argsQuery[INTERVAL])
	if err != nil {
		msg := api_base.CommonResponse{
			Status:  false,
			Message: err.Error(),
		}
		c.JSON(http.StatusBadRequest, msg)
		return
	}
	res := api_base.PaymentTimeV2{
		Past: util.FlattenPayment(p),
	}
	c.JSON(http.StatusOK, res)
}

// PaymentMissed
// @Summary Query missed scheduled payment for an employer.
// @Description Gives you missed scheduled payment. NO interval needed
// @Accept json
// @Produce json
// @Param employer_gid query string true "Employee's information. Corp Gid or Organization Gid"
// @Router /api/v1/payment/miss [get]
// @Success 200 {object} api_base.PaymentTimeV2
// @Failure 400 {object} api_base.CommonResponse
func PaymentMissed(c *gin.Context, db dao.GriffinWeb2Conn) {
	args := map[string]bool{
		EMPLOYEE_WORKFOR: true,
	}
	argsQuery, err := common.HandleOptionalQueryParam(c, args)
	if err != nil {
		return
	}
	payments, err := dao.QueryPaymentEmployer(argsQuery[EMPLOYEE_WORKFOR], ctx, db.Conn)
	if err != nil {
		msg := api_base.CommonResponse{
			Status:  false,
			Message: err.Error(),
		}
		c.JSON(http.StatusBadRequest, msg)
		return
	}
	p := dao.PaymentMissed(payments)
	res := api_base.PaymentTimeV2{
		Missed: util.FlattenPayment(p),
	}
	c.JSON(http.StatusOK, res)
}

// TotalPayment
// @Summary Query missed scheduled payment for an employer.
// @Description Gives you missed scheduled payment. NO interval needed
// @Accept json
// @Produce json
// @Param employer_gid query string true "Employee's information. Corp Gid or Organization Gid"
// @Router /api/v1/payment/miss [get]
// @Success 200 {object} api_base.PaymentTime
// @Failure 400 {object} api_base.CommonResponse
func TotalPayment(c *gin.Context, db dao.GriffinWeb2Conn) {
	args := map[string]bool{
		EMPLOYEE_WORKFOR: true,
	}
	_, errMsg, err := common.HandleOptionalQueryParamV2(c, args)
	if err != nil {
		c.JSON(http.StatusBadRequest, errMsg)
		return
	}
}
