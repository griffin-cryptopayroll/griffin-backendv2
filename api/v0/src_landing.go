package v0

import (
	"context"
	"database/sql"
	"griffin-dao/api/common"
	"griffin-dao/gcrud"
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
// @Success 200 {object} CommonResponse
// @Failure 500 {object} CommonResponse
func pingPong(c *gin.Context) {
	// Ping database with internal database/sql module
	newConn := gcrud.GriffinDataAccess{
		HostAddress: os.Getenv("DATABASE_ADDR"),
		PortAddress: os.Getenv("DATABASE_PORT"),
		Username:    os.Getenv("USERNAME"),
		Password:    os.Getenv("PASSWORD"),
		Name:        os.Getenv("NAME"),
	}
	newConnInfo := newConn.String()

	_, err := sql.Open("mysql", newConnInfo)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  false,
			"message": "failed to connect to database",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "pong",
	})
}

// version
// @Summary Read api-version file from environment file.
// @Description env file's parameter is GRIFFIN_WS_VERSION
// @Accept  json
// @Produce  json
// @Router /version [get]
// @Success 200 {object} CommonResponse
func version(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": os.Getenv("GRIFFIN_WS_VERSION"),
	})
}

// login
// @Summary Login into griffin payroll service
// @Description Matches Username with Password. If Username does not exists, 500 error.
// @Accept  json
// @Produce  json
// @Param username query string true "Employer's username (in email form)"
// @Param password query string true "Employer's password"
// @Router /login [get]
// @Success 200 {object} CommonResponseToken
// @Failure 400 {object} CommonResponse
// @Failure 403 {object} CommonResponse
// @Failure 500 {object} CommonResponse
func login(c *gin.Context, db gcrud.GriffinWeb2Conn) {
	args := map[string]bool{
		LOGIN_USERNAME: true,
		LOGIN_PASSWORD: true,
	}
	argsQuery, err := handleOptionalQueryParam(c, args)
	if err != nil {
		return
	}
	employer, err := gcrud.LoginHandler(argsQuery[LOGIN_USERNAME], context.Background(), db.Conn)
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{
			"status":  false,
			"message": "no username found",
		})
		return
	}
	if argsQuery[LOGIN_PASSWORD] != employer.Password {
		c.JSON(http.StatusForbidden, gin.H{
			"status":  false,
			"message": "incorrect password",
		})
		return
	}
	tk, _ := common.GenerateToken(employer.Gid)

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": employer.Gid,
		"token":   tk,
	})
	return
}
