package api

import (
	"fmt"
	"griffin-dao/gcrud"

	"github.com/gin-gonic/gin"
)

func postEmployee(c *gin.Context, db gcrud.GriffinWeb2Conn) {
	a := []string{
		EMPLOYEE_NAME, EMPLOYEE_EMAIL, EMPLOYEE_POSITION,
		EMPLOYEE_ACCOUNT, EMPLOYEE_PAYROLL, EMPLOYEE_PAYDAY, EMPLOYEE_CURRENCY,
	}
	argsQuery, err := handleQueryParam(c, a...)
	if err != nil {
		return
	}
	fmt.Println(argsQuery)
	fmt.Println(c.QueryMap(EMPLOYEE_NAME))
}
