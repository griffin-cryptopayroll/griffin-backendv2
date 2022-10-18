package v1

import (
	"context"
	"database/sql"
	"griffin-dao/gcrud"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

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
			"message": "failed to connect to database",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

func version(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": os.Getenv("GRIFFIN_WS_VERSION"),
	})
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header(ALLOW_ORIGIN, ALLOW_ORIGIN_VALUE)
		c.Header(ALLOW_CREDENTIALS, ALLOW_CREDENTIALS_VALUE)
		c.Next()
	}
}

func login(c *gin.Context, db gcrud.GriffinWeb2Conn) {
	args := map[string]bool{
		LOGIN_USERNAME: true,
		LOGIN_PASSWORD: true,
	}
	argsQuery, err := handleOptionalQueryParam(c, args)
	if err != nil {
		return
	}
	employer, err := gcrud.QueryLoginPassword(argsQuery[LOGIN_USERNAME], context.Background(), db.Conn)
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{
			"message": "no username found",
		})
	}
	if argsQuery[LOGIN_PASSWORD] != employer.Password {
		c.JSON(http.StatusForbidden, gin.H{
			"message": "incorrect password",
		})
		return
	}
	c.JSON(http.StatusOK, employer)
}
