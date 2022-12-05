package v0

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"griffin-dao/gcrud"
	"griffin-dao/price"
	"net/http"
	"strconv"
	"strings"
	"time"
)

// GetPrice
// @Summary Get all the price information that Griffin serves
// @Description ETH, MATIC data from binance.
// @Accept  json
// @Produce  json
// @Router /price [get]
// @Success 200 {object} price.PriceInformation
// @Failure 400 {object} CommonResponse
func getBinanceTrade(c *gin.Context) {
	currNeed := []string{price.ETHEREUM, price.POLYGON, price.USDC}
	result := map[string]float64{}
	for _, m := range currNeed {
		p, err := price.BinancePrice(m)
		if err != nil {
			c.JSON(http.StatusNoContent, gin.H{
				"message": PRICE_GET_ERROR,
			})
			return
		}
		result[m] = p
	}

	// send price information
	p := price.PriceInformation{
		Ethereum: result[price.ETHEREUM],
		Polygon:  result[price.POLYGON],
		USDC:     result[price.USDC],
	}
	c.JSON(http.StatusOK, p)
}

// PaymentHistory
// @Summary Create payment log for employee
// @Description Employee is identified by employer and employee Gid.
// @Description After update recommend updating payment date for employee
// @Accept json
// @Produce json
// @Param gid query string true "Employee's griffin id (in uuid form)"
// @Param employer_gid query string true "Employee's information. Corp Gid or Organization Gid"
// @Router /payment [post]
// @Success 200 {object} CommonResponse
// @Failure 400 {object} CommonResponse
// @Failure 500 {object} CommonResponse
func PaymentHistory(c *gin.Context, db gcrud.GriffinWeb2Conn) {
	var ctx = context.Background()
	args := map[string]bool{
		EMPLOYEE_GID:     true,
		EMPLOYEE_WORKFOR: true,
	}
	argsQuery, err := handleOptionalQueryParam(c, args)
	if err != nil {
		return
	}

	result, err := gcrud.QueryEmployee(argsQuery[EMPLOYEE_GID], argsQuery[EMPLOYEE_WORKFOR], ctx, db.Conn)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": DATABASE_SELECT_FAIL,
		})
		return
	}
	err = gcrud.CreatePaymentHistory(*result, ctx, db.Conn)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  false,
			"message": DATABASE_UPDATE_FAIL,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": DATABASE_UPDATE_SUCCESS,
	})
}

// PaymentMade
// @Summary Query payment history logs.
// @Description Payment history will be searched from standard date to standard date - interval
// @Description For example, if standard date is 2006-02-02, and interval is 1D,
// @Description it will search from 2006-02-01 ~ 2006-02-02.
// @Accept json
// @Produce json
// @Param standard query string true "Standard date which acts as a base point for interval search. YYYYMMDD"
// @Param interval_amount query int true "If 1 day put 1. If 2 week put 2"
// @Param interval_unit query int true "If 1 day put `D`. If 1 week put `W`. Only Offers "D" for dates and "W" for weeks"
// @Router /payment/past [get]
// @Success 200 {object} []ent.PAYMENT_HISTORY
// @Failure 400 {object} CommonResponse
// @Failure 403 {object} CommonResponse
// @Failure 500 {object} CommonResponse
func PaymentMade(c *gin.Context, db gcrud.GriffinWeb2Conn) {
	var ctx = context.Background()
	args := map[string]bool{
		STANDARD_DATE:   true,
		INTERVAL_AMOUNT: true,
		INTERVAL_UNIT:   true,
	}
	argsQuery, err := handleOptionalQueryParam(c, args)
	if err != nil {
		return
	}

	// Process Date information
	standardDate, err := time.Parse("20060102", argsQuery[STANDARD_DATE])
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": fmt.Sprintf("%s failed to meet date format of YYYYMMDD", argsQuery[STANDARD_DATE]),
		})
		return
	}
	var itv time.Duration
	switch strings.ToUpper(argsQuery[INTERVAL_UNIT]) {
	case "D":
		interval, err := strconv.Atoi(argsQuery[INTERVAL_AMOUNT])
		if err != nil {
			c.JSON(http.StatusForbidden, gin.H{
				"status":  false,
				"message": fmt.Sprintf("%s is not in right type", argsQuery[INTERVAL_AMOUNT]),
			})
			return
		}
		itv = time.Hour * time.Duration(24*interval)
	case "W":
		interval, err := strconv.Atoi(argsQuery[INTERVAL_AMOUNT])
		if err != nil {
			c.JSON(http.StatusForbidden, gin.H{
				"status":  false,
				"message": fmt.Sprintf("%s is not in right type", argsQuery[INTERVAL_AMOUNT]),
			})
			return
		}
		itv = time.Hour * time.Duration(24*7*interval)
	default:
		c.JSON(http.StatusForbidden, gin.H{
			"status":  false,
			"message": fmt.Sprintf("%s is not supported", argsQuery[INTERVAL_UNIT]),
		})
		return
	}
	period := gcrud.Period{
		StandardDate: standardDate,
		Interval:     itv,
	}

	obj, err := gcrud.PastPayment(period, ctx, db.Conn)
	c.JSON(http.StatusOK, obj)
}

// PaymentForFuture
// @Summary Query payment that will be made in the future.
// @Description Payment history will be searched from standard date to standard date + interval
// @Description For example, if standard date is 2006-02-02, and interval is 1D,
// @Description it will search from 2006-02-02 ~ 2006-02-03.
// @Accept json
// @Produce json
// @Param standard query string true "Standard date which acts as a base point for interval search. YYYYMMDD"
// @Param interval_amount query int true "If 1 day put 1. If 2 week put 2"
// @Param interval_unit query int true "If 1 day put `D`. If 1 week put `W`. Only Offers "D" for dates and "W" for weeks"
// @Router /payment/future [get]
// @Success 200 {object} []ent.EMPLOYEE
// @Failure 400 {object} CommonResponse
// @Failure 403 {object} CommonResponse
// @Failure 500 {object} CommonResponse
func PaymentForFuture(c *gin.Context, db gcrud.GriffinWeb2Conn) {
	var ctx = context.Background()
	args := map[string]bool{
		STANDARD_DATE:   true,
		INTERVAL_AMOUNT: true,
		INTERVAL_UNIT:   true,
	}
	argsQuery, err := handleOptionalQueryParam(c, args)
	if err != nil {
		return
	}

	// Process Date information
	standardDate, err := time.Parse("20060102", argsQuery[STANDARD_DATE])
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": fmt.Sprintf("%s failed to meet date format of YYYYMMDD", argsQuery[STANDARD_DATE]),
		})
		return
	}
	var itv time.Duration
	switch strings.ToUpper(argsQuery[INTERVAL_UNIT]) {
	case "D":
		interval, err := strconv.Atoi(argsQuery[INTERVAL_AMOUNT])
		if err != nil {
			c.JSON(http.StatusForbidden, gin.H{
				"status":  false,
				"message": fmt.Sprintf("%s is not in right type", argsQuery[INTERVAL_AMOUNT]),
			})
			return
		}
		itv = time.Hour * time.Duration(24*interval)
	case "W":
		interval, err := strconv.Atoi(argsQuery[INTERVAL_AMOUNT])
		if err != nil {
			c.JSON(http.StatusForbidden, gin.H{
				"status":  false,
				"message": fmt.Sprintf("%s is not in right type", argsQuery[INTERVAL_AMOUNT]),
			})
			return
		}
		itv = time.Hour * time.Duration(24*7*interval)
	default:
		c.JSON(http.StatusForbidden, gin.H{
			"status":  false,
			"message": fmt.Sprintf("%s is not supported", argsQuery[INTERVAL_UNIT]),
		})
		return
	}
	period := gcrud.Period{
		StandardDate: standardDate,
		Interval:     itv,
	}

	obj, err := gcrud.FuturePayment(period, ctx, db.Conn)
	c.JSON(http.StatusOK, obj)
}
