package v1

import (
	"context"
	"griffin-dao/gcrud"
	"griffin-dao/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func addEmpType(c *gin.Context, db gcrud.GriffinWeb2Conn) {
	args := map[string]bool{
		EMP_TYPE:  true,
		EMP_MONTH: true, // if EMP_MONTH is -1, then inf
	}
	argsQuery, err := handleOptionalQueryParam(c, args)
	if err != nil {
		return
	}

	contractPeriod, err := strconv.Atoi(argsQuery[EMP_MONTH])
	err = service.ValidateContractPeriod(contractPeriod, argsQuery[EMP_TYPE])
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	switch {
	case contractPeriod <= 0:
		gcrud.CreateEmployType("permanent", contractPeriod, context.Background(), db.Conn)
	case contractPeriod > 0:
		gcrud.CreateEmployType("contract", contractPeriod, context.Background(), db.Conn)
	}
}

func getEmpType(c *gin.Context, db gcrud.GriffinWeb2Conn) {
	args := map[string]bool{
		EMP_MONTH: true,
	}
	argsQuery, err := handleOptionalQueryParam(c, args)
	if err != nil {
		return
	}

	contractPeriod, err := strconv.Atoi(argsQuery[EMP_MONTH])
	err = service.ValidateContractPeriod(contractPeriod, argsQuery[EMP_TYPE])
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	gcrud.QueryEmployType(contractPeriod, context.Background(), db.Conn)
	c.JSON(http.StatusOK, gin.H{})
}

func delEmpType(c *gin.Context, db gcrud.GriffinWeb2Conn) {
	args := map[string]bool{
		EMP_MONTH: true,
	}
	argsQuery, err := handleOptionalQueryParam(c, args)
	if err != nil {
		return
	}

	contractPeriod, err := strconv.Atoi(argsQuery[EMP_MONTH])
	err = service.ValidateContractPeriod(contractPeriod, argsQuery[EMP_TYPE])
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	gcrud.DeleteEmployType(contractPeriod, context.Background(), db.Conn)
}
