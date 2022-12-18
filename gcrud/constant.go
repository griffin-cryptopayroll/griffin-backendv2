package gcrud

const (
	// database message success & error
	DATABASE_SELECT_SUCCESS = "database search successful"
	DATABASE_SELECT_FAIL    = "database search failed"
	DATABASE_CREATE_SUCCESS = "database new row added"
	DATABASE_CREATE_FAIL    = "fail to create new row"
	DATABASE_DELETE_SUCCESS = "database requested data deleted"
	DATABASE_DELETE_FAIL    = "fail to delete requested data"
	DATABASE_UPDATE_SUCCESS = "database requested data updated"
	DATABASE_UPDATE_FAIL    = "fail to update requested data"
	// request messages
	REQUEST_WRONG_TYPE    = "wrong type parameter"
	REQUEST_MISSING_PARAM = "missing parameter"
	// price information
	PRICE_GET_ERROR = "binance failure"
	// login information
	LOGIN_SUCCESS = "login successful"
	LOGIN_ERROR   = "login fail"
)

const (
	// exchange code
	BINANCE_CODE = 101
	FTX_CODE     = 102
	UPBIT_CODE   = 103
	MEXC_CODE    = 104
)

const (
	// exchange name. small letter case
	BINANCE = "binance"
	FTX     = "ftx"
	UPBIT   = "upbit"
	MEXC    = "mexc"
)
