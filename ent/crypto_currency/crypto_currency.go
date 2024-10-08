// Code generated by ent, DO NOT EDIT.

package crypto_currency

const (
	// Label holds the string label denoting the crypto_currency type in the database.
	Label = "crypto_currency"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldTicker holds the string denoting the ticker field in the database.
	FieldTicker = "ticker"
	// FieldSourceID holds the string denoting the source_id field in the database.
	FieldSourceID = "source_id"
	// EdgeCurrencyFromSource holds the string denoting the currency_from_source edge name in mutations.
	EdgeCurrencyFromSource = "currency_from_source"
	// EdgeCurrencyOfEmployee holds the string denoting the currency_of_employee edge name in mutations.
	EdgeCurrencyOfEmployee = "currency_of_employee"
	// EdgeCurrencyOfPaymentHistory holds the string denoting the currency_of_payment_history edge name in mutations.
	EdgeCurrencyOfPaymentHistory = "currency_of_payment_history"
	// EdgeCurrencyOfPayment holds the string denoting the currency_of_payment edge name in mutations.
	EdgeCurrencyOfPayment = "currency_of_payment"
	// Table holds the table name of the crypto_currency in the database.
	Table = "crypto_currency"
	// CurrencyFromSourceTable is the table that holds the currency_from_source relation/edge.
	CurrencyFromSourceTable = "crypto_currency"
	// CurrencyFromSourceInverseTable is the table name for the CRYPTO_PRC_SOURCE entity.
	// It exists in this package in order to avoid circular dependency with the "crypto_prc_source" package.
	CurrencyFromSourceInverseTable = "crypto_prc_source"
	// CurrencyFromSourceColumn is the table column denoting the currency_from_source relation/edge.
	CurrencyFromSourceColumn = "source_id"
	// CurrencyOfEmployeeTable is the table that holds the currency_of_employee relation/edge.
	CurrencyOfEmployeeTable = "employee"
	// CurrencyOfEmployeeInverseTable is the table name for the EMPLOYEE entity.
	// It exists in this package in order to avoid circular dependency with the "employee" package.
	CurrencyOfEmployeeInverseTable = "employee"
	// CurrencyOfEmployeeColumn is the table column denoting the currency_of_employee relation/edge.
	CurrencyOfEmployeeColumn = "crypto_currency_id"
	// CurrencyOfPaymentHistoryTable is the table that holds the currency_of_payment_history relation/edge.
	CurrencyOfPaymentHistoryTable = "payment_history"
	// CurrencyOfPaymentHistoryInverseTable is the table name for the PAYMENT_HISTORY entity.
	// It exists in this package in order to avoid circular dependency with the "payment_history" package.
	CurrencyOfPaymentHistoryInverseTable = "payment_history"
	// CurrencyOfPaymentHistoryColumn is the table column denoting the currency_of_payment_history relation/edge.
	CurrencyOfPaymentHistoryColumn = "currency_id"
	// CurrencyOfPaymentTable is the table that holds the currency_of_payment relation/edge.
	CurrencyOfPaymentTable = "payment"
	// CurrencyOfPaymentInverseTable is the table name for the PAYMENT entity.
	// It exists in this package in order to avoid circular dependency with the "payment" package.
	CurrencyOfPaymentInverseTable = "payment"
	// CurrencyOfPaymentColumn is the table column denoting the currency_of_payment relation/edge.
	CurrencyOfPaymentColumn = "crypto_currency_id"
)

// Columns holds all SQL columns for crypto_currency fields.
var Columns = []string{
	FieldID,
	FieldTicker,
	FieldSourceID,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}
