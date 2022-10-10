package v1

// Access control
const (
	ALLOW_ORIGIN            = "Access-Control-Allow-Origin"
	ALLOW_ORIGIN_VALUE      = "*"
	ALLOW_CREDENTIALS       = "Access-Control-Allow-Credentials"
	ALLOW_CREDENTIALS_VALUE = "true"
)

// Function get employer query
const (
	EMPLOYER_GID      = "gid"
	EMPLOYER_ID       = "id"
	EMPLOYER_PW       = "pw"
	EMPLOYER_CORPNAME = "corp_name"
	EMPLOYER_EMAIL    = "corp_email"
	EMPLOYER_WALLET   = "wallet"
)

// Function new employer query
const (
	EMPLOYEE_LNAME     = "last_name"
	EMPLOYEE_FNAME     = "first_name"
	EMPLOYEE_GID       = "gid"
	EMPLOYEE_WORKFOR   = "employer_gid"
	EMPLOYEE_POSITION  = "position"
	EMPLOYEE_WALLET    = "wallet"
	EMPLOYEE_PAYROLL   = "payroll"
	EMPLOYEE_CURRENCY  = "currency"
	EMPLOYEE_PAYDAY    = "payday"
	EMPLOYEE_EMAIL     = "email"
	EMPLOYEE_TYPE      = "employ_type"
	EMPLOYEE_WORKSTART = "work_start"
	CONTRACT_MONTH     = "contract_month"
)

// Function employment type query
const (
	EMP_TYPE  = "empType"
	EMP_MONTH = "empMonth"
)

// Message
const (
	// database message success & error
	DATABASE_SELECT_SUCCESS = "database search successful"
	DATABASE_SELECT_FAIL    = "database search failed"
	DATABASE_CREATE_SUCCESS = "database new row added"
	DATABASE_CREATE_FAIL    = "fail to create new row"
	DATABASE_DELETE_SUCCESS = "database requested data deleted"
	DATABASE_DELETE_FAIL    = "fail to delete requested data"
	// request messages
	REQUEST_WRONG_TYPE    = "wrong type parameter"
	REQUEST_MISSING_PARAM = "missing parameter"
	// price information
	PRICE_GET_ERROR = "binance failure"
	// login information
	LOGIN_SUCCESS = "login successful"
	LOGIN_ERROR   = "login fail"
)

type CommonResponse struct {
	Message string `json:"message" example:"database search successful / failed"`
}
