package api

import (
	"context"
	"griffin-dao/gcrud"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

func getEmployer(c *gin.Context, db gcrud.GriffinWeb2Conn) {
	args := map[string]bool{
		EMPLOYER_GID: true,
	}
	argsQuery, err := handleOptionalQueryParam(c, args)
	if err != nil {
		return
	}
	employer := gcrud.QueryEmployer(argsQuery[EMPLOYER_GID], context.Background(), db.Conn)
	meet := gcrud.EmployerJson{
		Username:  employer.Username,
		Password:  employer.Password,
		GriffinID: employer.Gid,
		CorpName:  employer.CorpName,
		CorpEmail: employer.CorpEmail,
		Wallet:    employer.Wallet,
	}
	c.JSON(http.StatusOK, meet)
}

func addEmployer(c *gin.Context, db gcrud.GriffinWeb2Conn) {
	args := map[string]bool{
		EMPLOYER_GID:      true,
		EMPLOYER_ID:       true,
		EMPLOYER_PW:       true,
		EMPLOYER_CORPNAME: false,
		EMPLOYER_EMAIL:    false,
		EMPLOYER_WALLET:   false,
	}
	argsQuery, err := handleOptionalQueryParam(c, args)
	if err != nil {
		return
	}
	freshMeet := gcrud.EmployerJson{
		Username:  argsQuery[EMPLOYER_ID],
		Password:  argsQuery[EMPLOYER_PW],
		GriffinID: argsQuery[EMPLOYER_GID],
		CorpName:  argsQuery[EMPLOYER_CORPNAME],
		CorpEmail: argsQuery[EMPLOYER_EMAIL],
		Wallet:    argsQuery[EMPLOYER_WALLET],
		CreatedAt: time.Now(),
		CreatedBy: os.Getenv("UPDATER"),
		UpdatedAt: time.Now(),
		UpdatedBy: os.Getenv("UPDATER"),
	}
	gcrud.CreateEmployerUserInfo(freshMeet, context.Background(), db.Conn)
	c.JSON(http.StatusOK, gin.H{
		"message": DATABASE_CREATE_SUCCESS,
	})
}

func delEmployer(c *gin.Context, db gcrud.GriffinWeb2Conn) {
	args := map[string]bool{
		EMPLOYER_GID: true,
	}
	argsQuery, err := handleOptionalQueryParam(c, args)
	if err != nil {
		return
	}
	gcrud.DeleteEmployer(argsQuery[EMPLOYER_GID], context.Background(), db.Conn)
	gcrud.DeleteEmployeewEmployerAll(argsQuery[EMPLOYER_GID], context.Background(), db.Conn)

	c.JSON(http.StatusOK, gin.H{
		"message": DATABASE_DELETE_SUCCESS,
	})
}

func updEmployer(c *gin.Context, db gcrud.GriffinWeb2Conn) {

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
