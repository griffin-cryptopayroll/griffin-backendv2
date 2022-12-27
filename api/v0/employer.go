package v0

import (
	"context"
	"github.com/gin-gonic/gin"
	"griffin-dao/ent"
	"griffin-dao/gcrud"
	"net/http"
	"os"
)

// EmployerWithGid
// @Summary Query employer from the database
// @Description Employer Griffin ID is in UUID form. Login will give you access to UUID.
// @Accept  json
// @Produce  json
// @Param gid query string true "Employer's griffin id (in uuid form)"
// @Router /employer [get]
// @Success 200 {object} ent.EMPLOYER
// @Failure 400 {object} CommonResponse
// @Failure 500 {object} CommonResponse
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
			"status":  false,
			"message": DATABASE_SELECT_FAIL,
		})
		return
	}
	c.JSON(http.StatusOK, employer)
}

// addEmployer
// @Summary Post employer to the database
// @Description Employer information is registered by google form.
// @Accept  json
// @Produce  json
// @Param id query string false "Employer's user id"
// @Param pw query string true "Employer's user password."
// @Param corp_name query string false "Employer information (corp or organization name)"
// @Param corp_email query string true "Employer information (corp or organization email)"
// @Param wallet query string false "Employer's griffin id (in uuid form)"
// @Router /employer [post]
// @Success 200 {object} CommonResponse
// @Failure 400 {object} CommonResponse
// @Failure 500 {object} CommonResponse
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
	err = gcrud.CreateEmployer(freshMeat, ctx, db.Conn)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": DATABASE_CREATE_FAIL,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": DATABASE_CREATE_SUCCESS,
	})
}

// delEmployer
// @Summary Delete employer to the database
// @Description Deleting the employer will delete all the employee's related to that employer
// @Accept  json
// @Produce  json
// @Param gid query string true "Employer's griffin id (in uuid form)"
// @Router /employer [delete]
// @Success 200 {object} CommonResponse
// @Failure 400 {object} CommonResponse
// @Failure 500 {object} CommonResponse
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
			"status":  false,
			"message": DATABASE_DELETE_FAIL + ": employer all",
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": DATABASE_DELETE_SUCCESS,
	})
}

// UpdateEmployer
// @Summary Delete Employer.
// @Description Not yet implemented
// @Accept json
// @Produce json
// @Router /employer [put]
// @Failure 403 {object} CommonResponse
func updEmployer(c *gin.Context, db gcrud.GriffinWeb2Conn) {
	// NotImplemented
	c.JSON(http.StatusForbidden, gin.H{
		"status":  false,
		"message": "NotImplemented",
	})
}
