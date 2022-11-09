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
	// Source holds the value of the "source" field.
	Source int `json:"source,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the CRYPTO_CURRENCYQuery when eager-loading is set.
	Edges                      CRYPTO_CURRENCYEdges `json:"edges"`
	crypto_prc_source_price_of *int
}

// CRYPTO_CURRENCYEdges holds the relations/edges for other nodes in the graph.
type CRYPTO_CURRENCYEdges struct {
	// SourceOf holds the value of the source_of edge.
	SourceOf *CRYPTO_PRC_SOURCE `json:"source_of,omitempty"`
	// CurrencyEmployee holds the value of the currency_employee edge.
	CurrencyEmployee []*EMPLOYEE `json:"currency_employee,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// SourceOfOrErr returns the SourceOf value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e CRYPTO_CURRENCYEdges) SourceOfOrErr() (*CRYPTO_PRC_SOURCE, error) {
	if e.loadedTypes[0] {
		if e.SourceOf == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: crypto_prc_source.Label}
		}
		return e.SourceOf, nil
	}
	return nil, &NotLoadedError{edge: "source_of"}
}

// CurrencyEmployeeOrErr returns the CurrencyEmployee value or an error if the edge
// was not loaded in eager-loading.
func (e CRYPTO_CURRENCYEdges) CurrencyEmployeeOrErr() ([]*EMPLOYEE, error) {
	if e.loadedTypes[1] {
		return e.CurrencyEmployee, nil
	}
	return nil, &NotLoadedError{edge: "currency_employee"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*CRYPTO_CURRENCY) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case crypto_currency.FieldID, crypto_currency.FieldSource:
			values[i] = new(sql.NullInt64)
		case crypto_currency.FieldTicker:
			values[i] = new(sql.NullString)
		case crypto_currency.ForeignKeys[0]: // crypto_prc_source_price_of
			values[i] = new(sql.NullInt64)
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
		case crypto_currency.FieldSource:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field source", values[i])
			} else if value.Valid {
				cc.Source = int(value.Int64)
			}
		case crypto_currency.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field crypto_prc_source_price_of", value)
			} else if value.Valid {
				cc.crypto_prc_source_price_of = new(int)
				*cc.crypto_prc_source_price_of = int(value.Int64)
			}
		}
	}
	return nil
}

// QuerySourceOf queries the "source_of" edge of the CRYPTO_CURRENCY entity.
func (cc *CRYPTO_CURRENCY) QuerySourceOf() *CRYPTOPRCSOURCEQuery {
	return (&CRYPTO_CURRENCYClient{config: cc.config}).QuerySourceOf(cc)
}

// QueryCurrencyEmployee queries the "currency_employee" edge of the CRYPTO_CURRENCY entity.
func (cc *CRYPTO_CURRENCY) QueryCurrencyEmployee() *EMPLOYEEQuery {
	return (&CRYPTO_CURRENCYClient{config: cc.config}).QueryCurrencyEmployee(cc)
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
	builder.WriteString("source=")
	builder.WriteString(fmt.Sprintf("%v", cc.Source))
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
