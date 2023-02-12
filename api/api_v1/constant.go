package api_v1

import "context"

var ctx = context.Background()

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
)

// Function payment log
const (
	INTERVAL = "interval"
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
// 1. Database message success & error
const (
	DATABASE_SELECT_SUCCESS = "database search successful"
	DATABASE_SELECT_FAIL    = "database search failed"
	DATABASE_CREATE_SUCCESS = "database new row added"
	DATABASE_CREATE_FAIL    = "fail to create new row"
	DATABASE_DELETE_SUCCESS = "database requested data deleted"
	DATABASE_DELETE_FAIL    = "fail to delete requested data"
	DATABASE_UPDATE_SUCCESS = "database requested data updated"
	DATABASE_UPDATE_FAIL    = "fail to update requested data"
)

// 2. HTTP request messages
const (
	REQUEST_WRONG_TYPE    = "wrong type parameter"
	REQUEST_MISSING_PARAM = "missing parameter"
)

// 3. Price information request to Binance.
const (
	PRICE_GET_ERROR = "binance failure"
)

// 4. Meta information
const (
	META_GET_PRICE_INFO = "Price Information gathered from binance. For U/S use different server"
)
