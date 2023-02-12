package api_v1

import (
	"fmt"
	api_base "griffin-dao/api/base"
	"griffin-dao/util"
	"griffin-dao/util/price"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetPriceInfo
// @Summary Get price information that Griffin serves.
// @Description ETH, MATIC and USDC data from binance.
// @Accept json
// @Produce json
// @Router /api/v1/price [get]
// @Success 200 {object} price.PriceInfoContainer
// @Success 204 {object} api_base.CommonResponse
func GetPriceInfo(c *gin.Context) {
	currNeed := []string{price.ETHEREUM, price.POLYGON, price.USDC}

	ok := false
	result := map[string]float64{}
	for _, m := range currNeed {
		p, err := price.BinancePrice(m)
		if err != nil {
			msg := fmt.Sprintf("%s: asset: %s", PRICE_GET_ERROR, m)
			util.PrintRedError(msg)
			continue
		}
		ok = true
		result[m] = p
	}

	if !ok {
		// No single price information is present.
		msg := api_base.CommonResponse{
			Status:  false,
			Message: fmt.Sprintf("%s: asset: ALL ASSET", PRICE_GET_ERROR),
		}
		c.JSON(http.StatusNoContent, msg)
		return
	}

	p := price.PriceInfoContainer{
		Meta: META_GET_PRICE_INFO,
		Data: price.PriceInformationV2{
			Ethereum: result[price.ETHEREUM],
			Polygon:  result[price.POLYGON],
			USDC:     result[price.USDC],
		},
	}
	c.JSON(http.StatusOK, p)
}
