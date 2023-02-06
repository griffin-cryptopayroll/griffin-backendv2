package api_login

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/spruceid/siwe-go"
	api_base "griffin-dao/api/base"
	"griffin-dao/dao"
	"griffin-dao/util"
	"net/http"
	"os"
	"time"
)

var ctx = context.Background()

type SignInWithEthMessage struct {
	Message   string `json:"message"`
	Signature string `json:"signature"`
}

type SignInWithEthNOnce struct {
	Status bool   `json:"status,omitempty" example:"true"`
	Nonce  string `json:"nonce,omitempty"`
}

type SessionStore struct {
	ID          string
	Gid         string
	Valid       time.Time
	LastKnownIP string
}

var SessionMapInfo = map[string]SessionStore{}
var SessionMapLogin = map[string]int{}

// SiweNonce
// @Summary Provide Nonce value. (Numbers used only once)
// @Description Provide Nonce value that's needed for SIWE login
// @Accept  json
// @Produce  json
// @Router /api/v1/nonce [get]
// @Success 200 {object} SignInWithEthNOnce
func SiweNonce(c *gin.Context) {
	nonceMessage := SignInWithEthNOnce{
		Status: true,
		Nonce:  siwe.GenerateNonce(),
	}
	c.JSON(http.StatusOK, nonceMessage)
}

// SiweVerify
// @Summary Login using SIWE (Sign in with Ethereum)
// @Description Using strict format string, in json, post login information and get valid SessionID.
// @Accept json
// @Produce json
// @Router /api/v1/verify [post]
// @Success 200 {object} api_base.CommonResponse
// @Failure 400 {object} api_base.CommonResponse
// @Failure 403 {object} api_base.CommonResponse
func SiweVerify(c *gin.Context, db dao.GriffinWeb2Conn) {
	util.PrintYellowStatus("Perform login verification. Provide `sID` if successful")
	// 1. Get Incoming message from requesting body - strict SIWE form
	v := SignInWithEthMessage{}
	err := c.ShouldBindJSON(&v)
	if err != nil {
		util.PrintRedError("SignInWithEth. Message parse to json error", err.Error())
		msg := api_base.CommonResponse{
			Status:  false,
			Message: err.Error(),
		}
		c.JSON(http.StatusBadRequest, msg)
		return
	}

	// 2. Verify using package
	walletAddr, err := Authentication(v)
	if err != nil {
		// Record client ID if failed login.
		util.PrintRedError(
			fmt.Sprintf("User with anonymous key failed to login from [%s]", c.ClientIP()),
		)
		msg := api_base.CommonResponse{
			Status:  false,
			Message: err.Error(),
		}
		c.JSON(http.StatusForbidden, msg)
		return
	}
	// Verify user wallet inside Web2 database.
	userGid, err := dao.RegisteredWallet(walletAddr, ctx, db.Conn)
	if err != nil {
		util.PrintRedError(
			fmt.Sprintf("User with wrong wallet [%s] failed to login from [%s]", walletAddr, c.ClientIP()),
		)
		msg := api_base.CommonResponse{
			Status:  false,
			Message: err.Error(),
		}
		c.JSON(http.StatusForbidden, msg)
		return
	}

	// Verification successful. Give the client session ID
	util.PrintGreenStatus(
		fmt.Sprintf(
			"User with address [%s] logged in from [%s]",
			walletAddr,
			c.ClientIP(),
		),
	)
	sID := uuid.New().String()

	// Save it on server memory (without DB)
	userSessionInfo := SessionStore{
		ID:          sID,
		Gid:         userGid,
		Valid:       time.Now().Add(time.Hour),
		LastKnownIP: c.ClientIP(),
	}
	SessionMapInfo[sID] = userSessionInfo
	SessionMapLogin[sID] = 1

	// "github.com/golang-jwt/jwt/v4"
	fmt.Println(SessionMapInfo)
	fmt.Println(SessionMapLogin)
	msg := api_base.LoginResponse{
		Status: true,
		Message: fmt.Sprintf(
			"Session ID [%s] has been provided in key `sID`",
			sID,
		),
		Gid: userGid,
	}
	// SetCookie function will write `sID` on the response header
	c.SetCookie("sID", sID, 3600, "/", os.Getenv("HOSTNAME"), false, false)
	c.JSON(http.StatusOK, msg)
}

func SiweVerifyToken(c *gin.Context, db dao.GriffinWeb2Conn) {
	util.PrintYellowStatus("Perform login verification. Provide `token` if successful")
	// 1. Get Incoming message from requesting body - strict SIWE form
	v := SignInWithEthMessage{}
	err := c.ShouldBindJSON(&v)
	if err != nil {
		util.PrintRedError("SignInWithEth. Message parse to json error", err.Error())
		msg := api_base.CommonResponse{
			Status:  false,
			Message: err.Error(),
		}
		c.JSON(http.StatusBadRequest, msg)
		return
	}

	// 2.1 Verify SIWE message using package
	walletAddr, err := Authentication(v)
	if err != nil {
		// Record client ID if failed login.
		util.PrintRedError(
			fmt.Sprintf("User with anonymous key failed to login from [%s]", c.ClientIP()),
		)
		msg := api_base.CommonResponse{
			Status:  false,
			Message: err.Error(),
		}
		c.JSON(http.StatusForbidden, msg)
		return
	}
	// 2.2 Verify user wallet inside Web2 database.
	userGid, err := dao.RegisteredWallet(walletAddr, ctx, db.Conn)
	if err != nil {
		util.PrintRedError(
			fmt.Sprintf("User with wrong wallet [%s] failed to login from [%s]", walletAddr, c.ClientIP()),
		)
		msg := api_base.CommonResponse{
			Status:  false,
			Message: err.Error(),
		}
		c.JSON(http.StatusForbidden, msg)
		return
	}
	// 2.3 Distribute JWT Token
	tokenStr, err := createJWTToken(userGid)
	if err != nil {
		msg := api_base.CommonResponse{
			Status:  false,
			Message: err.Error(),
		}
		c.JSON(http.StatusInternalServerError, msg)
		return
	}

	// Verification successful. Give the client session ID
	util.PrintGreenStatus(
		fmt.Sprintf(
			"User with address [%s] logged in from [%s]. Token [%s] distributed",
			walletAddr,
			c.ClientIP(),
			tokenStr,
		),
	)
	msg := api_base.LoginResponse{
		Status:  true,
		Message: fmt.Sprintf("Token [%s] distributed", tokenStr),
		Gid:     userGid,
	}
	c.SetCookie("token", tokenStr, 3600, "/", os.Getenv("HOSTNAME"), false, false)
	c.JSON(http.StatusOK, msg)
}

func Authentication(v SignInWithEthMessage) (string, error) {
	authenticate, err := siwe.ParseMessage(v.Message)
	if err != nil {
		util.PrintRedError("SignInWithEth, SIWE format parse error", err.Error())
		return "", err
	}

	pk, err := authenticate.Verify(v.Signature, nil, nil, nil)
	if err != nil {
		util.PrintRedError("SignInWithEth, Verification error", err.Error())
	}
	addr := crypto.PubkeyToAddress(*pk)
	return addr.String(), nil
}
