package api_v1

import (
	api_base "griffin-dao/api/base"
	"griffin-dao/api/common"
	"griffin-dao/dao"
	"griffin-dao/ent"
	"net/http"

	"github.com/gin-gonic/gin"
)

// AddEmpType
// @Summary Add Employ type.
// @Description Whether the employer is full-time worker(permanent) or contract worker(freelance)
// @Accept  json
// @Produce  json
// @Param is_perma query string true "Enumerator (permanent, freelance)"
// @Param pay_freq query string true "Single formed frequency, such as D(Day), W(Week)"
// @Router /api/v1/employType [post]
// @Success 200 {object} api_base.CommonResponse
// @Failure 400 {object} api_base.CommonResponse
// @Failure 403 {object} api_base.CommonResponse
// @Failure 500 {object} api_base.CommonResponse
func AddEmpType(c *gin.Context, db dao.GriffinWeb2Conn) {
	args := map[string]bool{
		EMP_TYPE: true,
		// If EMP_PAY_FREQ == -1, then inf payment
		EMP_PAY_FREQ: true,
	}
	argsQuery, err := common.HandleOptionalQueryParam(c, args)
	if err != nil {
		return
	}

	err = dao.CreateEmployType(argsQuery[EMP_TYPE], argsQuery[EMP_PAY_FREQ], ctx, db.Conn)
	if err != nil {
		msg := api_base.CommonResponse{
			Status:  false,
			Message: DATABASE_CREATE_FAIL,
		}
		c.JSON(http.StatusInternalServerError, msg)
		return
	}

	msg := api_base.CommonResponse{
		Status:  true,
		Message: DATABASE_CREATE_SUCCESS,
	}
	c.JSON(http.StatusOK, msg)
}

// GetEmpType
// @Summary Get Employ type.
// @Description Whether the employer is full-time worker(permanent) or contract worker(freelance)
// @Accept  json
// @Produce  json
// @Param is_perma query string true "Enumerator (permanent or freelance)"
// @Param pay_freq query string true "Single formed frequency, such as D(Day), W(Week)"
// @Router /api/v1/employType [get]
// @Success 200 {object} api_base.CommonResponseData[ent.EMPLOY_TYPE]
// @Failure 403 {object} api_base.CommonResponse
// @Failure 500 {object} api_base.CommonResponse
func GetEmpType(c *gin.Context, db dao.GriffinWeb2Conn) {
	args := map[string]bool{
		EMP_TYPE:     true,
		EMP_PAY_FREQ: true,
	}
	argsQuery, err := common.HandleOptionalQueryParam(c, args)
	if err != nil {
		return
	}

	et, err := dao.QueryEmployType(argsQuery[EMP_TYPE], argsQuery[EMP_PAY_FREQ], ctx, db.Conn)
	if err != nil {
		msg := api_base.CommonResponse{
			Status:  false,
			Message: DATABASE_SELECT_FAIL,
		}
		c.JSON(http.StatusInternalServerError, msg)
		return
	}

	packet := api_base.CommonResponseData[*ent.EMPLOY_TYPE]{
		Status:        true,
		Message:       DATABASE_SELECT_SUCCESS,
		DataContainer: et,
	}
	c.JSON(http.StatusOK, packet)
}

// DelEmpType
// @Summary Delete Employ type.
// @Description Not yet implemented
// @Accept json
// @Produce json
// @Router /api/v1/employType [delete]
// @Failure 403 {object} api_base.CommonResponse
func DelEmpType(c *gin.Context, db dao.GriffinWeb2Conn) {
	// Not Implemented
	msg := api_base.CommonResponse{
		Status:  false,
		Message: "Not implemented",
	}
	c.JSON(http.StatusForbidden, msg)
}
