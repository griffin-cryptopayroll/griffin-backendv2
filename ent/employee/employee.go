// Code generated by ent, DO NOT EDIT.

package employee

const (
	// Label holds the string label denoting the employee type in the database.
	Label = "employee"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldGid holds the string denoting the gid field in the database.
	FieldGid = "gid"
	// FieldEmployerGid holds the string denoting the employer_gid field in the database.
	FieldEmployerGid = "employer_gid"
	// FieldLastName holds the string denoting the last_name field in the database.
	FieldLastName = "last_name"
	// FieldFirstName holds the string denoting the first_name field in the database.
	FieldFirstName = "first_name"
	// FieldPosition holds the string denoting the position field in the database.
	FieldPosition = "position"
	// FieldWallet holds the string denoting the wallet field in the database.
	FieldWallet = "wallet"
	// FieldPayroll holds the string denoting the payroll field in the database.
	FieldPayroll = "payroll"
	// FieldCurrency holds the string denoting the currency field in the database.
	FieldCurrency = "currency"
	// FieldPayday holds the string denoting the payday field in the database.
	FieldPayday = "payday"
	// FieldEmploy holds the string denoting the employ field in the database.
	FieldEmploy = "employ"
	// FieldEmail holds the string denoting the email field in the database.
	FieldEmail = "email"
	// FieldWorkStart holds the string denoting the work_start field in the database.
	FieldWorkStart = "work_start"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldCreatedBy holds the string denoting the created_by field in the database.
	FieldCreatedBy = "created_by"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldUpdatedBy holds the string denoting the updated_by field in the database.
	FieldUpdatedBy = "updated_by"
	// EdgeEmployeeGets holds the string denoting the employee_gets edge name in mutations.
	EdgeEmployeeGets = "employee_gets"
	// EdgeEmployeeTypeFrom holds the string denoting the employee_type_from edge name in mutations.
	EdgeEmployeeTypeFrom = "employee_type_from"
	// EdgeWorkFor holds the string denoting the work_for edge name in mutations.
	EdgeWorkFor = "work_for"
	// EdgePaymentHistory holds the string denoting the payment_history edge name in mutations.
	EdgePaymentHistory = "payment_history"
	// Table holds the table name of the employee in the database.
	Table = "employe_es"
	// EmployeeGetsTable is the table that holds the employee_gets relation/edge.
	EmployeeGetsTable = "employe_es"
	// EmployeeGetsInverseTable is the table name for the CRYPTO_CURRENCY entity.
	// It exists in this package in order to avoid circular dependency with the "crypto_currency" package.
	EmployeeGetsInverseTable = "crypto_currenc_ys"
	// EmployeeGetsColumn is the table column denoting the employee_gets relation/edge.
	EmployeeGetsColumn = "crypto_currency_employee_paid"
	// EmployeeTypeFromTable is the table that holds the employee_type_from relation/edge.
	EmployeeTypeFromTable = "employe_es"
	// EmployeeTypeFromInverseTable is the table name for the EMPLOY_TYPE entity.
	// It exists in this package in order to avoid circular dependency with the "employ_type" package.
	EmployeeTypeFromInverseTable = "employ_typ_es"
	// EmployeeTypeFromColumn is the table column denoting the employee_type_from relation/edge.
	EmployeeTypeFromColumn = "employ_type_employee_type_to"
	// WorkForTable is the table that holds the work_for relation/edge.
	WorkForTable = "employe_es"
	// WorkForInverseTable is the table name for the EMPLOYER_USER_INFO entity.
	// It exists in this package in order to avoid circular dependency with the "employer_user_info" package.
	WorkForInverseTable = "employer_user_inf_os"
	// WorkForColumn is the table column denoting the work_for relation/edge.
	WorkForColumn = "employer_user_info_work_under"
	// PaymentHistoryTable is the table that holds the payment_history relation/edge.
	PaymentHistoryTable = "payment_histor_ys"
	// PaymentHistoryInverseTable is the table name for the PAYMENT_HISTORY entity.
	// It exists in this package in order to avoid circular dependency with the "payment_history" package.
	PaymentHistoryInverseTable = "payment_histor_ys"
	// PaymentHistoryColumn is the table column denoting the payment_history relation/edge.
	PaymentHistoryColumn = "employee_payment_history"
)

// Columns holds all SQL columns for employee fields.
var Columns = []string{
	FieldID,
	FieldGid,
	FieldEmployerGid,
	FieldLastName,
	FieldFirstName,
	FieldPosition,
	FieldWallet,
	FieldPayroll,
	FieldCurrency,
	FieldPayday,
	FieldEmploy,
	FieldEmail,
	FieldWorkStart,
	FieldCreatedAt,
	FieldCreatedBy,
	FieldUpdatedAt,
	FieldUpdatedBy,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "employe_es"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"crypto_currency_employee_paid",
	"employer_user_info_work_under",
	"employ_type_employee_type_to",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}
