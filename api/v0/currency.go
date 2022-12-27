package v0

import (
	"github.com/gin-gonic/gin"
	"griffin-dao/price"
	"net/http"
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
