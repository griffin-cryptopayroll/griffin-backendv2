package price

const (
	// server point
	BINANCE    = "https://api.binance.com/api/v3"
	BINANCE_US = "https://api.binance.us/api/v3"
	// request endpoint
	TRADE = "/trades"
	// needed currency
	ETHEREUM = "ETHUSDT"
	POLYGON  = "MATICUSDT"
	USDC     = "USDCUSDT"
)

type BinanceRequest struct {
	Symbol string
	Limit  string
}

type BinanceResponse struct {
	Id            int    `json:"id"`
	Price         string `json:"price"`
	Quantity      string `json:"qty"`
	QuoteQuantity string `json:"quoteQty"`
	UnixTime      int    `json:"time"`
	IsBuyerMaker  bool   `json:"isBuyerMaker"`
	IsBestMatch   bool   `json:"isBestMatch"`
}

type PriceInformation struct {
	Ethereum float64 `json:"ethusdt"`
	Polygon  float64 `json:"maticusdt"`
	USDC     float64 `json:"usdcusdt"`
}

type MonthlyPayHistory struct {
	Date     string `json:"date"`
	Ethereum int    `json:"eth"`
	Polygon  int    `json:"matic"`
	USDC     int    `json:"usdc"`
}

type PriceInfoContainer struct {
	Meta      string             `json:"meta" example:"some meta information"`
	Dimension Dim                `json:"dimension,omitempty"`
	Data      PriceInformationV2 `json:"data"`
}

type PriceInformationV2 struct {
	Ethereum float64 `json:"eth" example:"1605.10"`
	Polygon  float64 `json:"matic" example:"0.9985"`
	USDC     float64 `json:"usdc" example:"0.9999"`
}

type Dim struct {
	RowLength    int `json:"rowl,omitempty" example:"1"`
	ColumnLength int `json:"coll,omitempty" example:"2"`
}
