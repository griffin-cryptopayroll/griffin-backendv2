package v1

import (
	"context"
	"github.com/gin-gonic/gin"
	"griffin-dao/ent"
	"griffin-dao/gcrud"
	"net/http"
	"os"
)

func EmployerWithGid(c *gin.Context, db gcrud.GriffinWeb2Conn) {
	var ctx = context.Background()
	args := map[string]bool{
		EMPLOYER_GID: true,
	}
	argsQuery, err := handleOptionalQueryParam(c, args)
	if err != nil {
		return
	}

	employer, err := gcrud.QueryEmployerwEmployerGid(
		argsQuery[EMPLOYER_GID],
		ctx,
		db.Conn,
	)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": DATABASE_SELECT_FAIL,
		})
		return
	}
	c.JSON(http.StatusOK, employer)
}

func addEmployer(c *gin.Context, db gcrud.GriffinWeb2Conn) {
	var ctx = context.Background()
	args := map[string]bool{
		EMPLOYER_ID:       false,
		EMPLOYER_PW:       true,
		EMPLOYER_CORPNAME: false,
		EMPLOYER_EMAIL:    true,
		EMPLOYER_WALLET:   false,
	}
	argsQuery, err := handleOptionalQueryParam(c, args)
	if err != nil {
		return
	}

	freshMeat := ent.EMPLOYER{
		Username:  argsQuery[EMPLOYER_ID],
		Password:  argsQuery[EMPLOYER_PW],
		CorpName:  argsQuery[EMPLOYER_CORPNAME],
		CorpEmail: argsQuery[EMPLOYER_EMAIL],
		Wallet:    argsQuery[EMPLOYER_WALLET],
		CreatedBy: os.Getenv("UPDATER"),
		UpdatedBy: os.Getenv("UPDATER"),
	}
	err = gcrud.CreateEmployerUserInfo(freshMeat, ctx, db.Conn)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": DATABASE_CREATE_FAIL,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": DATABASE_CREATE_SUCCESS,
	})
}

func delEmployer(c *gin.Context, db gcrud.GriffinWeb2Conn) {
	var ctx = context.Background()
	args := map[string]bool{
		EMPLOYER_GID: true,
	}
	argsQuery, err := handleOptionalQueryParam(c, args)
	if err != nil {
		return
	}

	err = gcrud.DeleteEmployeewEmployerAll(
		argsQuery[EMPLOYER_GID],
		ctx,
		db.Conn,
	)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": DATABASE_DELETE_FAIL + ": employer all",
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"message": DATABASE_DELETE_SUCCESS,
	})
}

func updEmployer(c *gin.Context, db gcrud.GriffinWeb2Conn) {
	// NotImplemented
	c.JSON(http.StatusForbidden, gin.H{
		"message": "NotImplemented",
	})
}
