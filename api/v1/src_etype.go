package v1

import (
	"context"
	"github.com/gin-gonic/gin"
	"griffin-dao/gcrud"
	"net/http"
)

func addEmpType(c *gin.Context, db gcrud.GriffinWeb2Conn) {
	args := map[string]bool{
		EMP_TYPE:     true,
		EMP_PAY_FREQ: true, // if EMP_PAY_FREQ is -1, then inf
	}
	argsQuery, err := handleOptionalQueryParam(c, args)
	if err != nil {
		return
	}
	err = gcrud.CreateEmployType(argsQuery[EMP_TYPE], argsQuery[EMP_PAY_FREQ], context.Background(), db.Conn)
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

func getEmpType(c *gin.Context, db gcrud.GriffinWeb2Conn) {
	args := map[string]bool{
		EMP_TYPE:     true,
		EMP_PAY_FREQ: true,
	}
	argsQuery, err := handleOptionalQueryParam(c, args)
	if err != nil {
		return
	}

	et, err := gcrud.QueryEmployType(argsQuery[EMP_TYPE], argsQuery[EMP_PAY_FREQ], context.Background(), db.Conn)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": DATABASE_SELECT_FAIL,
		})
		return
	}
	c.JSON(http.StatusOK, et)
}

func delEmpType(c *gin.Context, db gcrud.GriffinWeb2Conn) {
	// Not Implemented
	c.JSON(http.StatusForbidden, gin.H{
		"message": "NotImplemented",
	})
}
