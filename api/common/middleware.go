package common

import (
	"github.com/gin-gonic/gin"
)

const (
	AllowOrigin         = "Access-Control-Allow-Origin"
	AllowCredentials    = "Access-Control-Allow-Credentials"
	AllowControlHeaders = "Access-Control-Allow-Headers"
	AllowControlMethods = "Access-Control-Allow-Methods"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set(AllowOrigin, "*")
		c.Writer.Header().Set(AllowCredentials, "true")
		c.Writer.Header().Set(AllowControlHeaders, "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set(AllowControlMethods, "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	}
}

//
//func JwtAuthMiddleware() gin.HandlerFunc {
//	return func(c *gin.Context) {
//		err := TokenValid(c)
//		if err != nil {
//			c.JSON(http.StatusUnauthorized, gin.H{
//				"message": "Unauthorized",
//			})
//			c.Abort()
//			return
//		}
//		c.Next()
//	}
//}
