// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"griffin-dao/ent/employ_type"
	"strings"

	"entgo.io/ent/dialect/sql"
)

// EMPLOY_TYPE is the model entity for the EMPLOY_TYPE schema.
type EMPLOY_TYPE struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// IsPermanent holds the value of the "is_permanent" field.
	IsPermanent string `json:"is_permanent,omitempty"`
	// PayFreq holds the value of the "pay_freq" field.
	PayFreq string `json:"pay_freq,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the EMPLOY_TYPEQuery when eager-loading is set.
	Edges EMPLOY_TYPEEdges `json:"edges"`
}

// EMPLOY_TYPEEdges holds the relations/edges for other nodes in the graph.
type EMPLOY_TYPEEdges struct {
	// EmployeeTypeTo holds the value of the employee_type_to edge.
	EmployeeTypeTo []*EMPLOYEE `json:"employee_type_to,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// EmployeeTypeToOrErr returns the EmployeeTypeTo value or an error if the edge
// was not loaded in eager-loading.
func (e EMPLOY_TYPEEdges) EmployeeTypeToOrErr() ([]*EMPLOYEE, error) {
	if e.loadedTypes[0] {
		return e.EmployeeTypeTo, nil
	}
	return nil, &NotLoadedError{edge: "employee_type_to"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*EMPLOY_TYPE) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case employ_type.FieldID:
			values[i] = new(sql.NullInt64)
		case employ_type.FieldIsPermanent, employ_type.FieldPayFreq:
			values[i] = new(sql.NullString)
		default:
			return nil, fmt.Errorf("unexpected column %q for type EMPLOY_TYPE", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the EMPLOY_TYPE fields.
func (et *EMPLOY_TYPE) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case employ_type.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			et.ID = int(value.Int64)
		case employ_type.FieldIsPermanent:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field is_permanent", values[i])
			} else if value.Valid {
				et.IsPermanent = value.String
			}
		case employ_type.FieldPayFreq:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field pay_freq", values[i])
			} else if value.Valid {
				et.PayFreq = value.String
			}
		}
	}
	return nil
}

// QueryEmployeeTypeTo queries the "employee_type_to" edge of the EMPLOY_TYPE entity.
func (et *EMPLOY_TYPE) QueryEmployeeTypeTo() *EMPLOYEEQuery {
	return (&EMPLOY_TYPEClient{config: et.config}).QueryEmployeeTypeTo(et)
}

// Update returns a builder for updating this EMPLOY_TYPE.
// Note that you need to call EMPLOY_TYPE.Unwrap() before calling this method if this EMPLOY_TYPE
// was returned from a transaction, and the transaction was committed or rolled back.
func (et *EMPLOY_TYPE) Update() *EMPLOYTYPEUpdateOne {
	return (&EMPLOY_TYPEClient{config: et.config}).UpdateOne(et)
}

// Unwrap unwraps the EMPLOY_TYPE entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (et *EMPLOY_TYPE) Unwrap() *EMPLOY_TYPE {
	_tx, ok := et.config.driver.(*txDriver)
	if !ok {
		panic("ent: EMPLOY_TYPE is not a transactional entity")
	}
	et.config.driver = _tx.drv
	return et
}

// String implements the fmt.Stringer.
func (et *EMPLOY_TYPE) String() string {
	var builder strings.Builder
	builder.WriteString("EMPLOY_TYPE(")
	builder.WriteString(fmt.Sprintf("id=%v, ", et.ID))
	builder.WriteString("is_permanent=")
	builder.WriteString(et.IsPermanent)
	builder.WriteString(", ")
	builder.WriteString("pay_freq=")
	builder.WriteString(et.PayFreq)
	builder.WriteByte(')')
	return builder.String()
}

// EMPLOY_TYPEs is a parsable slice of EMPLOY_TYPE.
type EMPLOY_TYPEs []*EMPLOY_TYPE

func (et EMPLOY_TYPEs) config(cfg config) {
	for _i := range et {
		et[_i].config = cfg
	}
}
