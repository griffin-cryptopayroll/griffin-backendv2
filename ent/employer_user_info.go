// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"griffin-dao/ent/employer_user_info"
	"strings"

	"entgo.io/ent/dialect/sql"
)

// EMPLOYER_USER_INFO is the model entity for the EMPLOYER_USER_INFO schema.
type EMPLOYER_USER_INFO struct {
	config
	// ID of the ent.
	ID int `json:"id,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*EMPLOYER_USER_INFO) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case employer_user_info.FieldID:
			values[i] = new(sql.NullInt64)
		default:
			return nil, fmt.Errorf("unexpected column %q for type EMPLOYER_USER_INFO", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the EMPLOYER_USER_INFO fields.
func (eui *EMPLOYER_USER_INFO) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case employer_user_info.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			eui.ID = int(value.Int64)
		}
	}
	return nil
}

// Update returns a builder for updating this EMPLOYER_USER_INFO.
// Note that you need to call EMPLOYER_USER_INFO.Unwrap() before calling this method if this EMPLOYER_USER_INFO
// was returned from a transaction, and the transaction was committed or rolled back.
func (eui *EMPLOYER_USER_INFO) Update() *EMPLOYERUSERINFOUpdateOne {
	return (&EMPLOYER_USER_INFOClient{config: eui.config}).UpdateOne(eui)
}

// Unwrap unwraps the EMPLOYER_USER_INFO entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (eui *EMPLOYER_USER_INFO) Unwrap() *EMPLOYER_USER_INFO {
	_tx, ok := eui.config.driver.(*txDriver)
	if !ok {
		panic("ent: EMPLOYER_USER_INFO is not a transactional entity")
	}
	eui.config.driver = _tx.drv
	return eui
}

// String implements the fmt.Stringer.
func (eui *EMPLOYER_USER_INFO) String() string {
	var builder strings.Builder
	builder.WriteString("EMPLOYER_USER_INFO(")
	builder.WriteString(fmt.Sprintf("id=%v", eui.ID))
	builder.WriteByte(')')
	return builder.String()
}

// EMPLOYER_USER_INFOs is a parsable slice of EMPLOYER_USER_INFO.
type EMPLOYER_USER_INFOs []*EMPLOYER_USER_INFO

func (eui EMPLOYER_USER_INFOs) config(cfg config) {
	for _i := range eui {
		eui[_i].config = cfg
	}
}
