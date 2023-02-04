package server

import (
	"database/sql"
	api_base "griffin-dao/api/base"
	"griffin-dao/dao"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

// pingPong
// @Summary Server, DB ping test
// @Description Check 1) server is alive 2) database is alive by pinging it.
// @Description Database ping using internal sql method in golang
// @Accept  json
// @Produce  json
// @Router /ping [get]
// @Success 200 {object} api_base.CommonResponse
// @Failure 500 {object} api_base.CommonResponse
func pingPong(c *gin.Context) {
	// Ping database with internal database/sql module
	newConn := dao.GriffinDataAccess{
		HostAddress: os.Getenv("DATABASE_ADDR"),
		PortAddress: os.Getenv("DATABASE_PORT"),
		Username:    os.Getenv("USERNAME"),
		Password:    os.Getenv("PASSWORD"),
		Name:        os.Getenv("NAME"),
	}
	newConnInfo := newConn.String()

	_, err := sql.Open("mysql", newConnInfo)
	if err != nil {
		msg := api_base.CommonResponse{
			Status:  false,
			Message: "failed to connect to database",
		}
		c.JSON(http.StatusInternalServerError, msg)
		return
	}

	msg := api_base.CommonResponse{
		Status:  true,
		Message: "pong",
	}
	c.JSON(http.StatusOK, msg)
}

// version
// @Summary Read api-version file from environment file.
// @Description env file's parameter is GRIFFIN_WS_VERSION
// @Accept  json
// @Produce  json
// @Router /version [get]
// @Success 200 {object} api_base.CommonResponse
func version(c *gin.Context) {
	msg := api_base.CommonResponse{
		Status:  true,
		Message: os.Getenv("GRIFFIN_WS_VERSION"),
	}
	c.JSON(http.StatusOK, msg)
}
