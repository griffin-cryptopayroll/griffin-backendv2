package v1

import (
	"context"
	"github.com/gin-gonic/gin"
	"griffin-dao/gcrud"
	"net/http"
)

// addEmpType
// @Summary Add Employ type.
// @Description Whether the employer is full-time worker(fulltime) or contract worker(contract)
// @Accept  json
// @Produce  json
// @Param isPerma query string true "Enumerator (fulltime or contract)"
// @Param payFreq query string true "Single formed frequency, such as D(Day), W(Week), M(Month), Y(Year)"
// @Router /employType [post]
// @Success 200 {object} CommonResponse
// @Failure 400 {object} CommonResponse
// @Failure 403 {object} CommonResponse
// @Failure 500 {object} CommonResponse
func addEmpType(c *gin.Context, db gcrud.GriffinWeb2Conn) {
	var ctx = context.Background()
	args := map[string]bool{
		EMP_TYPE:     true,
		EMP_PAY_FREQ: true, // if EMP_PAY_FREQ is -1, then inf
	}
	argsQuery, err := handleOptionalQueryParam(c, args)
	if err != nil {
		return
	}
	err = gcrud.CreateEmployType(
		argsQuery[EMP_TYPE],
		argsQuery[EMP_PAY_FREQ],
		ctx,
		db.Conn,
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

// getEmpType
// @Summary Get Employ type.
// @Description Whether the employer is full-time worker(fulltime) or contract worker(contract)
// @Accept  json
// @Produce  json
// @Param isPerma query string true "Enumerator (fulltime or contract)"
// @Param payFreq query string true "Single formed frequency, such as D(Day), W(Week), M(Month), Y(Year)"
// @Router /employType [get]
// @Success 200 {object} ent.EMPLOY_TYPE
// @Failure 400 {object} CommonResponse
// @Failure 403 {object} CommonResponse
// @Failrue 500 {object} CommonResponse
func getEmpType(c *gin.Context, db gcrud.GriffinWeb2Conn) {
	var ctx = context.Background()
	args := map[string]bool{
		EMP_TYPE:     true,
		EMP_PAY_FREQ: true,
	}
	argsQuery, err := handleOptionalQueryParam(c, args)
	if err != nil {
		return
	}

	et, err := gcrud.QueryEmployType(
		argsQuery[EMP_TYPE],
		argsQuery[EMP_PAY_FREQ],
		ctx,
		db.Conn,
	)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": DATABASE_SELECT_FAIL,
		})
		return
	}
	c.JSON(http.StatusOK, et)
}

// delEmpType
// @Summary Delete Employ type.
// @Description Not yet implemented
// @Accept json
// @Produce json
// @Router /employType [delete]
// @Failure 403 {object} CommonResponse
func delEmpType(c *gin.Context, db gcrud.GriffinWeb2Conn) {
	// Not Implemented
	c.JSON(http.StatusForbidden, gin.H{
		"message": "NotImplemented",
	})
}
