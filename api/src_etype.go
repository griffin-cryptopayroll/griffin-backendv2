package api

import (
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
	
}

func getEmpType(c *gin.Context, db gcrud.GriffinWeb2Conn) {

}

func delEmpType(c *gin.Context, db gcrud.GriffinWeb2Conn) {

}
