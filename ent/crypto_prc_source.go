// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"griffin-dao/ent/crypto_prc_source"
	"strings"

	"entgo.io/ent/dialect/sql"
)

// CRYPTO_PRC_SOURCE is the model entity for the CRYPTO_PRC_SOURCE schema.
type CRYPTO_PRC_SOURCE struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the CRYPTO_PRC_SOURCEQuery when eager-loading is set.
	Edges CRYPTO_PRC_SOURCEEdges `json:"edges"`
}

// CRYPTO_PRC_SOURCEEdges holds the relations/edges for other nodes in the graph.
type CRYPTO_PRC_SOURCEEdges struct {
	// SourceOfCurrency holds the value of the source_of_currency edge.
	SourceOfCurrency []*CRYPTO_CURRENCY `json:"source_of_currency,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// SourceOfCurrencyOrErr returns the SourceOfCurrency value or an error if the edge
// was not loaded in eager-loading.
func (e CRYPTO_PRC_SOURCEEdges) SourceOfCurrencyOrErr() ([]*CRYPTO_CURRENCY, error) {
	if e.loadedTypes[0] {
		return e.SourceOfCurrency, nil
	}
	return nil, &NotLoadedError{edge: "source_of_currency"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*CRYPTO_PRC_SOURCE) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case crypto_prc_source.FieldID:
			values[i] = new(sql.NullInt64)
		case crypto_prc_source.FieldName:
			values[i] = new(sql.NullString)
		default:
			return nil, fmt.Errorf("unexpected column %q for type CRYPTO_PRC_SOURCE", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the CRYPTO_PRC_SOURCE fields.
func (cps *CRYPTO_PRC_SOURCE) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case crypto_prc_source.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			cps.ID = int(value.Int64)
		case crypto_prc_source.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				cps.Name = value.String
			}
		}
	}
	return nil
}

// QuerySourceOfCurrency queries the "source_of_currency" edge of the CRYPTO_PRC_SOURCE entity.
func (cps *CRYPTO_PRC_SOURCE) QuerySourceOfCurrency() *CRYPTOCURRENCYQuery {
	return (&CRYPTO_PRC_SOURCEClient{config: cps.config}).QuerySourceOfCurrency(cps)
}

// Update returns a builder for updating this CRYPTO_PRC_SOURCE.
// Note that you need to call CRYPTO_PRC_SOURCE.Unwrap() before calling this method if this CRYPTO_PRC_SOURCE
// was returned from a transaction, and the transaction was committed or rolled back.
func (cps *CRYPTO_PRC_SOURCE) Update() *CRYPTOPRCSOURCEUpdateOne {
	return (&CRYPTO_PRC_SOURCEClient{config: cps.config}).UpdateOne(cps)
}

// Unwrap unwraps the CRYPTO_PRC_SOURCE entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (cps *CRYPTO_PRC_SOURCE) Unwrap() *CRYPTO_PRC_SOURCE {
	_tx, ok := cps.config.driver.(*txDriver)
	if !ok {
		panic("ent: CRYPTO_PRC_SOURCE is not a transactional entity")
	}
	cps.config.driver = _tx.drv
	return cps
}

// String implements the fmt.Stringer.
func (cps *CRYPTO_PRC_SOURCE) String() string {
	var builder strings.Builder
	builder.WriteString("CRYPTO_PRC_SOURCE(")
	builder.WriteString(fmt.Sprintf("id=%v, ", cps.ID))
	builder.WriteString("name=")
	builder.WriteString(cps.Name)
	builder.WriteByte(')')
	return builder.String()
}

// CRYPTO_PRC_SOURCEs is a parsable slice of CRYPTO_PRC_SOURCE.
type CRYPTO_PRC_SOURCEs []*CRYPTO_PRC_SOURCE

func (cps CRYPTO_PRC_SOURCEs) config(cfg config) {
	for _i := range cps {
		cps[_i].config = cfg
	}
}
