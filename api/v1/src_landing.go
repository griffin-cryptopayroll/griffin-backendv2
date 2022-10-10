package v1

import (
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
