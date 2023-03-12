package api_v0

import (
	"context"
	"github.com/gin-gonic/gin"
	api_base "griffin-dao/api/base"
	"griffin-dao/api/common"
	"griffin-dao/dao"
	"griffin-dao/ent"
	"net/http"
	"os"
)

// EmployerWithGid
// @Summary Query employer from the database
// @Description Employer Griffin ID is in UUID form. Login will give you access to UUID.
// @Accept  json
// @Produce  json
// @Param gid query string true "Employer's griffin id (in uuid form)"
// @Router /api/v0/employer [get]
// @Success 200 {object} ent.EMPLOYER
// @Failure 400 {object} api_base.CommonResponse
// @Failure 500 {object} api_base.CommonResponse
func EmployerWithGid(c *gin.Context, db dao.GriffinWeb2Conn) {
	var ctx = context.Background()
	args := map[string]bool{
		EMPLOYER_GID: true,
	}
	argsQuery, err := common.HandleOptionalQueryParam(c, args)
	if err != nil {
		return
	}

	employer, err := dao.QueryEmployerwEmployerGid(argsQuery[EMPLOYER_GID], ctx, db.Conn)
	if err != nil {
		msg := api_base.CommonResponse{
			Status:  false,
			Message: DATABASE_SELECT_FAIL,
		}
		c.JSON(http.StatusBadRequest, msg)
		return
	}

	c.JSON(http.StatusOK, employer)
}

// AddEmployer
// @Summary Post employer to the database
// @Description Employer information is registered by google form.
// @Accept  json
// @Produce  json
// @Param id query string false "Employer's user id"
// @Param pw query string true "Employer's user password."
// @Param corp_name query string false "Employer information (corp or organization name)"
// @Param corp_email query string true "Employer information (corp or organization email)"
// @Param wallet query string true "Employer's wallet"
// @Param wallet_aztec string true "Employer's aztec wallet"
// @Router /api/v0/employer [post]
// @Success 200 {object} api_base.CommonResponse
// @Failure 400 {object} api_base.CommonResponse
// @Failure 500 {object} api_base.CommonResponse
func AddEmployer(c *gin.Context, db dao.GriffinWeb2Conn) {
	var ctx = context.Background()
	args := map[string]bool{
		EMPLOYER_ID:           false,
		EMPLOYER_PW:           true,
		EMPLOYER_CORPNAME:     false,
		EMPLOYER_EMAIL:        true,
		EMPLOYER_WALLET:       true,
		EMPLOYER_WALLET_AZTEC: true,
	}
	argsQuery, err := common.HandleOptionalQueryParam(c, args)
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
	err = dao.CreateEmployer(freshMeat, ctx, db.Conn)
	if err != nil {
		msg := api_base.CommonResponse{
			Status:  false,
			Message: DATABASE_CREATE_FAIL,
		}
		c.JSON(http.StatusBadRequest, msg)
		return
	}

	msg := api_base.CommonResponse{
		Status:  true,
		Message: DATABASE_CREATE_SUCCESS,
	}
	c.JSON(http.StatusOK, msg)
}

// DelEmployer
// @Summary Delete employer to the database
// @Description Deleting the employer will delete all the employee's related to that employer
// @Accept  json
// @Produce  json
// @Param gid query string true "Employer's griffin id (in uuid form)"
// @Router /api/v0/employer [delete]
// @Success 200 {object} api_base.CommonResponse
// @Failure 400 {object} api_base.CommonResponse
// @Failure 500 {object} api_base.CommonResponse
func DelEmployer(c *gin.Context, db dao.GriffinWeb2Conn) {
	var ctx = context.Background()
	args := map[string]bool{
		EMPLOYER_GID: true,
	}
	argsQuery, err := common.HandleOptionalQueryParam(c, args)
	if err != nil {
		return
	}

	err = dao.DeleteEmployeewEmployerAll(argsQuery[EMPLOYER_GID], ctx, db.Conn)
	if err != nil {
		msg := api_base.CommonResponse{
			Status:  false,
			Message: DATABASE_DELETE_FAIL + ": employer all",
		}
		c.JSON(http.StatusBadRequest, msg)
	}

	msg := api_base.CommonResponse{
		Status:  true,
		Message: DATABASE_DELETE_SUCCESS,
	}
	c.JSON(http.StatusOK, msg)
}

// UpdEmployer
// @Summary Delete Employer.
// @Description Not yet implemented
// @Accept json
// @Produce json
// @Router /api/v0/employer [put]
// @Failure 403 {object} api_base.CommonResponse
func UpdEmployer(c *gin.Context, db dao.GriffinWeb2Conn) {
	// NotImplemented
	msg := api_base.CommonResponse{
		Status:  false,
		Message: "Not implemented",
	}
	c.JSON(http.StatusForbidden, msg)
}
