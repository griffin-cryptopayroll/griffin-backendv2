package api

import (
	"context"
	"griffin-dao/gcrud"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

func postEmployer(c *gin.Context, db gcrud.GriffinWeb2Conn) {
	argsQuery, err := handleQueryParam(c, EMPLOYER_GID, EMPLOYER_ID, EMPLOYER_PW)
	if err != nil {
		return
	}
	freshMeet := gcrud.EmployerJson{
		GriffinID: argsQuery[EMPLOYER_GID],
		Username:  argsQuery[EMPLOYER_ID],
		Password:  argsQuery[EMPLOYER_PW],
		CreatedAt: time.Now(),
		CreatedBy: os.Getenv("UPDATER"),
		UpdatedAt: time.Now(),
		UpdatedBy: os.Getenv("UPDATER"),
	}
	gcrud.CreateEmployerUserInfo(freshMeet, context.Background(), db.Conn)
}

func delEmployer(c *gin.Context, db gcrud.GriffinWeb2Conn) {
	argsQuery, err := handleQueryParam(c, EMPLOYER_GID)
	if err != nil {
		return
	}
	gcrud.DeleteEmployer(argsQuery[EMPLOYER_GID], context.Background(), db.Conn)
	gcrud.DeleteEmployeewEmployer(argsQuery[EMPLOYER_GID], context.Background(), db.Conn)
	c.JSON(http.StatusOK, gin.H{
		"message": DATABASE_DELETE_SUCCESS,
	})
}

func updateEmployer(c *gin.Context, db gcrud.GriffinWeb2Conn) {

}

//
//func loginEmployer(c *gin.Context, db *redis.Client) {
//	username, err := handleParamEmployerGid(c)
//	if err != nil {
//		return
//	}
//	password, err := handleParamEmployerPw(c)
//	if err != nil {
//		return
//	}
//
//	var registered [][]rdb.Login
//	_ = rdb.JsonGet(
//		db,
//		LOGIN_KEY,
//		LOGIN_PATH,
//		&registered,
//	)
//	id, pw := isRegistered(registered[0], username, password)
//	switch {
//	case id && pw:
//		c.JSON(http.StatusOK, gin.H{
//			"message":    LOGIN_SUCCESS,
//			"employerId": username,
//			"employerPw": password,
//		})
//	case id:
//		// if wrong password StatusUnauthorized
//		c.JSON(http.StatusUnauthorized, gin.H{
//			"message": LOGIN_ERROR + " wrong password",
//		})
//	default:
//		// if no id StatusForbidden
//		c.JSON(http.StatusForbidden, gin.H{
//			"message": LOGIN_ERROR + " id unrecognized",
//		})
//	}
//}
//
