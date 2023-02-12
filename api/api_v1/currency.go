package api_v0

import (
	"github.com/gin-gonic/gin"
	api_base "griffin-dao/api/base"
	"griffin-dao/util/price"
	"net/http"
)

// GetBinanceTrade
// @Summary Get all the price information that Griffin serves
// @Description ETH, MATIC data from binance.
// @Accept  json
// @Produce  json
// @Router /api/v0/price [get]
// @Success 200 {object} price.PriceInformation
// @Success 204 {object} api_base.CommonResponse
func GetBinanceTrade(c *gin.Context) {
	currNeed := []string{price.ETHEREUM, price.POLYGON, price.USDC}
	result := map[string]float64{}
	for _, m := range currNeed {
		p, err := price.BinancePrice(m)
		if err != nil {
			msg := api_base.CommonResponse{
				Status:  false,
				Message: PRICE_GET_ERROR,
			}
			c.JSON(http.StatusNoContent, msg)
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
