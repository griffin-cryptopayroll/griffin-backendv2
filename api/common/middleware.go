package common

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	api_login "griffin-dao/api/login"
	"griffin-dao/util"
	"net/http"
	"os"
)

const (
	AllowOrigin         = "Access-Control-Allow-Origin"
	AllowCredentials    = "Access-Control-Allow-Credentials"
	AllowControlHeaders = "Access-Control-Allow-Headers"
	AllowControlMethods = "Access-Control-Allow-Methods"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("CORS limit at", os.Getenv("CORS_LIMIT"))
		c.Writer.Header().Set(AllowOrigin, os.Getenv("CORS_LIMIT"))
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

// TokenAuthMiddleware
// After this middleware, user can perform certain actions
// iff the user provides legal JWT Token. Allowed Session ID is
// stored in Redis Database and has a lifespan of 600 seconds. (10 mins)
// User will provide JWT Token inside a cookie with key `token`.
// func obtains it via variable `tknRecv`.
func TokenAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tknRecv, err := c.Cookie("token")
		if err != nil {
			util.PrintRedError(err.Error())
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		claims := &api_login.Claims{}
		tkn, err := jwt.ParseWithClaims(tknRecv, claims, func(token *jwt.Token) (interface{}, error) {
			return api_login.JwtKey, nil
		})
		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				m := fmt.Sprintf("Invalid Token access from %s", c.ClientIP())
				util.PrintRedError(m)
				c.AbortWithStatus(http.StatusUnauthorized)
				return
			}
			m := fmt.Sprintf("Invalid Token access from %s: %v", c.ClientIP(), err)
			util.PrintRedError(m)
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}

		if !tkn.Valid {
			m := fmt.Sprintf("Expired Token access from %s", c.ClientIP())
			util.PrintRedError(m)
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		// Successful authentication - proceed with `c.Next()`
		util.PrintGreenStatus("User", claims.Username, "authenticate")
		c.Next()
	}
}
