package api

import (
	"context"
	"github.com/gin-gonic/gin"
	"griffin-dao/gcrud"
	"os"
	"time"
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

//func delEmployer(c *gin.Context, db *redis.Client) {
//	Employer, err := handleParamEmployerGid(c)
//	if err != nil {
//		return
//	}
//	_ = rdb.JsonDel(
//		db,
//		fmt.Sprintf(PERMANENT_EMPLOYER_KEY, Employer),
//		PERMANENT_EMPLOYER_PATH,
//	)
//	_ = rdb.JsonDel(
//		db,
//		fmt.Sprintf(FREELANCE_EMPLOYER_KEY, Employer),
//		FREELANCE_EMPLOYER_PATH,
//	)
//	c.JSON(http.StatusOK, gin.H{
//		"message": DATABASE_DELETE_SUCCESS,
//	})
//}

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
