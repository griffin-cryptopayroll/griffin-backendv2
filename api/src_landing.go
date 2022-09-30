package api

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func pingPong(c *gin.Context) {
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
