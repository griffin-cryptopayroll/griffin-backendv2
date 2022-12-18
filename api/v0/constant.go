package v0

import "griffin-dao/ent"

const (
	LOGIN_USERNAME = "username"
	LOGIN_PASSWORD = "password"
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
	EMPLOYEE_GID       = "gid"
	EMPLOYEE_NAME      = "name"
	EMPLOYEE_WORKFOR   = "employer_gid"
	EMPLOYEE_POSITION  = "position"
	EMPLOYEE_WALLET    = "wallet"
	EMPLOYEE_PAYROLL   = "payroll"
	EMPLOYEE_CURRENCY  = "currency"
	EMPLOYEE_PAYDAY    = "payday"
	EMPLOYEE_EMAIL     = "email"
	EMPLOYEE_TYPE      = "employ_type"
	EMPLOYEE_WORKSTART = "work_start"
	EMPLOYEE_WORKEND   = "work_end"
	CONTRACT_MONTH     = "contract_month"
)

// Function payment history
const (
	STANDARD_DATE   = "standard"
	INTERVAL_UNIT   = "interval_unit"
	INTERVAL_AMOUNT = "interval_amount"
)

// Function payment
const (
	SCHEDULED_PAYMENT = "schd_date"
	EXECUTED_PAYMENT  = "exec_date"
	ONEOFF_PAYMENT    = "onof_date"
)

// Function employment type query
const (
	EMP_TYPE     = "is_perma"
	EMP_PAY_FREQ = "pay_freq"
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

type CommonResponse struct {
	Status  bool   `json:"status" example:"true"`
	Message string `json:"message" example:"database (create / delete) (successful / failed)"`
}

type CommonResponseToken struct {
	Status  bool   `json:"status" example:"true"`
	Message string `json:"message" example:"<employer gid>"`
	Token   string `json:"token" example:"<employer JWT token>"`
}

type PaymentType struct {
	Scheduled []*ent.PAYMENT `json:"scheduled"`
	Executed  []*ent.PAYMENT `json:"executed"`
	OneOff    []*ent.PAYMENT `json:"oneoff"`
}
