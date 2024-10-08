// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"griffin-dao/ent/crypto_currency"
	"griffin-dao/ent/crypto_prc_source"
	"strings"

	"entgo.io/ent/dialect/sql"
)

// CRYPTO_CURRENCY is the model entity for the CRYPTO_CURRENCY schema.
type CRYPTO_CURRENCY struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Ticker holds the value of the "ticker" field.
	Ticker string `json:"ticker,omitempty"`
	// SourceID holds the value of the "source_id" field.
	SourceID int `json:"source_id,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the CRYPTO_CURRENCYQuery when eager-loading is set.
	Edges CRYPTO_CURRENCYEdges `json:"edges"`
}

// CRYPTO_CURRENCYEdges holds the relations/edges for other nodes in the graph.
type CRYPTO_CURRENCYEdges struct {
	// CurrencyFromSource holds the value of the currency_from_source edge.
	CurrencyFromSource *CRYPTO_PRC_SOURCE `json:"currency_from_source,omitempty"`
	// CurrencyOfEmployee holds the value of the currency_of_employee edge.
	CurrencyOfEmployee []*EMPLOYEE `json:"currency_of_employee,omitempty"`
	// CurrencyOfPaymentHistory holds the value of the currency_of_payment_history edge.
	CurrencyOfPaymentHistory []*PAYMENT_HISTORY `json:"currency_of_payment_history,omitempty"`
	// CurrencyOfPayment holds the value of the currency_of_payment edge.
	CurrencyOfPayment []*PAYMENT `json:"currency_of_payment,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [4]bool
}

// CurrencyFromSourceOrErr returns the CurrencyFromSource value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e CRYPTO_CURRENCYEdges) CurrencyFromSourceOrErr() (*CRYPTO_PRC_SOURCE, error) {
	if e.loadedTypes[0] {
		if e.CurrencyFromSource == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: crypto_prc_source.Label}
		}
		return e.CurrencyFromSource, nil
	}
	return nil, &NotLoadedError{edge: "currency_from_source"}
}

// CurrencyOfEmployeeOrErr returns the CurrencyOfEmployee value or an error if the edge
// was not loaded in eager-loading.
func (e CRYPTO_CURRENCYEdges) CurrencyOfEmployeeOrErr() ([]*EMPLOYEE, error) {
	if e.loadedTypes[1] {
		return e.CurrencyOfEmployee, nil
	}
	return nil, &NotLoadedError{edge: "currency_of_employee"}
}

// CurrencyOfPaymentHistoryOrErr returns the CurrencyOfPaymentHistory value or an error if the edge
// was not loaded in eager-loading.
func (e CRYPTO_CURRENCYEdges) CurrencyOfPaymentHistoryOrErr() ([]*PAYMENT_HISTORY, error) {
	if e.loadedTypes[2] {
		return e.CurrencyOfPaymentHistory, nil
	}
	return nil, &NotLoadedError{edge: "currency_of_payment_history"}
}

// CurrencyOfPaymentOrErr returns the CurrencyOfPayment value or an error if the edge
// was not loaded in eager-loading.
func (e CRYPTO_CURRENCYEdges) CurrencyOfPaymentOrErr() ([]*PAYMENT, error) {
	if e.loadedTypes[3] {
		return e.CurrencyOfPayment, nil
	}
	return nil, &NotLoadedError{edge: "currency_of_payment"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*CRYPTO_CURRENCY) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case crypto_currency.FieldID, crypto_currency.FieldSourceID:
			values[i] = new(sql.NullInt64)
		case crypto_currency.FieldTicker:
			values[i] = new(sql.NullString)
		default:
			return nil, fmt.Errorf("unexpected column %q for type CRYPTO_CURRENCY", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the CRYPTO_CURRENCY fields.
func (cc *CRYPTO_CURRENCY) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case crypto_currency.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			cc.ID = int(value.Int64)
		case crypto_currency.FieldTicker:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field ticker", values[i])
			} else if value.Valid {
				cc.Ticker = value.String
			}
		case crypto_currency.FieldSourceID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field source_id", values[i])
			} else if value.Valid {
				cc.SourceID = int(value.Int64)
			}
		}
	}
	return nil
}

// QueryCurrencyFromSource queries the "currency_from_source" edge of the CRYPTO_CURRENCY entity.
func (cc *CRYPTO_CURRENCY) QueryCurrencyFromSource() *CRYPTOPRCSOURCEQuery {
	return (&CRYPTO_CURRENCYClient{config: cc.config}).QueryCurrencyFromSource(cc)
}

// QueryCurrencyOfEmployee queries the "currency_of_employee" edge of the CRYPTO_CURRENCY entity.
func (cc *CRYPTO_CURRENCY) QueryCurrencyOfEmployee() *EMPLOYEEQuery {
	return (&CRYPTO_CURRENCYClient{config: cc.config}).QueryCurrencyOfEmployee(cc)
}

// QueryCurrencyOfPaymentHistory queries the "currency_of_payment_history" edge of the CRYPTO_CURRENCY entity.
func (cc *CRYPTO_CURRENCY) QueryCurrencyOfPaymentHistory() *PAYMENTHISTORYQuery {
	return (&CRYPTO_CURRENCYClient{config: cc.config}).QueryCurrencyOfPaymentHistory(cc)
}

// QueryCurrencyOfPayment queries the "currency_of_payment" edge of the CRYPTO_CURRENCY entity.
func (cc *CRYPTO_CURRENCY) QueryCurrencyOfPayment() *PAYMENTQuery {
	return (&CRYPTO_CURRENCYClient{config: cc.config}).QueryCurrencyOfPayment(cc)
}

// Update returns a builder for updating this CRYPTO_CURRENCY.
// Note that you need to call CRYPTO_CURRENCY.Unwrap() before calling this method if this CRYPTO_CURRENCY
// was returned from a transaction, and the transaction was committed or rolled back.
func (cc *CRYPTO_CURRENCY) Update() *CRYPTOCURRENCYUpdateOne {
	return (&CRYPTO_CURRENCYClient{config: cc.config}).UpdateOne(cc)
}

// Unwrap unwraps the CRYPTO_CURRENCY entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (cc *CRYPTO_CURRENCY) Unwrap() *CRYPTO_CURRENCY {
	_tx, ok := cc.config.driver.(*txDriver)
	if !ok {
		panic("ent: CRYPTO_CURRENCY is not a transactional entity")
	}
	cc.config.driver = _tx.drv
	return cc
}

// String implements the fmt.Stringer.
func (cc *CRYPTO_CURRENCY) String() string {
	var builder strings.Builder
	builder.WriteString("CRYPTO_CURRENCY(")
	builder.WriteString(fmt.Sprintf("id=%v, ", cc.ID))
	builder.WriteString("ticker=")
	builder.WriteString(cc.Ticker)
	builder.WriteString(", ")
	builder.WriteString("source_id=")
	builder.WriteString(fmt.Sprintf("%v", cc.SourceID))
	builder.WriteByte(')')
	return builder.String()
}

// CRYPTO_CURRENCYs is a parsable slice of CRYPTO_CURRENCY.
type CRYPTO_CURRENCYs []*CRYPTO_CURRENCY

func (cc CRYPTO_CURRENCYs) config(cfg config) {
	for _i := range cc {
		cc[_i].config = cfg
	}
}
