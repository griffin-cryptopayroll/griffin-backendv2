// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"griffin-dao/ent/crypto_currency"
	"griffin-dao/ent/employ_type"
	"griffin-dao/ent/employee"
	"griffin-dao/ent/employer_user_info"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
)

// EMPLOYEE is the model entity for the EMPLOYEE schema.
type EMPLOYEE struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Gid holds the value of the "gid" field.
	Gid string `json:"gid,omitempty"`
	// EmployerGid holds the value of the "employer_gid" field.
	EmployerGid string `json:"employer_gid,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Position holds the value of the "position" field.
	Position string `json:"position,omitempty"`
	// Wallet holds the value of the "wallet" field.
	Wallet string `json:"wallet,omitempty"`
	// Payroll holds the value of the "payroll" field.
	Payroll float64 `json:"payroll,omitempty"`
	// Currency holds the value of the "currency" field.
	Currency int `json:"currency,omitempty"`
	// Payday holds the value of the "payday" field.
	Payday time.Time `json:"payday,omitempty"`
	// Employ holds the value of the "employ" field.
	Employ int `json:"employ,omitempty"`
	// Email holds the value of the "email" field.
	Email string `json:"email,omitempty"`
	// WorkStart holds the value of the "work_start" field.
	WorkStart string `json:"work_start,omitempty"`
	// WorkEnds holds the value of the "work_ends" field.
	WorkEnds string `json:"work_ends,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// CreatedBy holds the value of the "created_by" field.
	CreatedBy string `json:"created_by,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// UpdatedBy holds the value of the "updated_by" field.
	UpdatedBy string `json:"updated_by,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the EMPLOYEEQuery when eager-loading is set.
	Edges                         EMPLOYEEEdges `json:"edges"`
	crypto_currency_employee_paid *int
	employer_user_info_work_under *int
	employ_type_employee_type_to  *int
}

// EMPLOYEEEdges holds the relations/edges for other nodes in the graph.
type EMPLOYEEEdges struct {
	// EmployeeGets holds the value of the employee_gets edge.
	EmployeeGets *CRYPTO_CURRENCY `json:"employee_gets,omitempty"`
	// EmployeeTypeFrom holds the value of the employee_type_from edge.
	EmployeeTypeFrom *EMPLOY_TYPE `json:"employee_type_from,omitempty"`
	// WorkFor holds the value of the work_for edge.
	WorkFor *EMPLOYER_USER_INFO `json:"work_for,omitempty"`
	// PaymentHistory holds the value of the payment_history edge.
	PaymentHistory []*PAYMENT_HISTORY `json:"payment_history,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [4]bool
}

// EmployeeGetsOrErr returns the EmployeeGets value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e EMPLOYEEEdges) EmployeeGetsOrErr() (*CRYPTO_CURRENCY, error) {
	if e.loadedTypes[0] {
		if e.EmployeeGets == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: crypto_currency.Label}
		}
		return e.EmployeeGets, nil
	}
	return nil, &NotLoadedError{edge: "employee_gets"}
}

// EmployeeTypeFromOrErr returns the EmployeeTypeFrom value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e EMPLOYEEEdges) EmployeeTypeFromOrErr() (*EMPLOY_TYPE, error) {
	if e.loadedTypes[1] {
		if e.EmployeeTypeFrom == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: employ_type.Label}
		}
		return e.EmployeeTypeFrom, nil
	}
	return nil, &NotLoadedError{edge: "employee_type_from"}
}

// WorkForOrErr returns the WorkFor value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e EMPLOYEEEdges) WorkForOrErr() (*EMPLOYER_USER_INFO, error) {
	if e.loadedTypes[2] {
		if e.WorkFor == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: employer_user_info.Label}
		}
		return e.WorkFor, nil
	}
	return nil, &NotLoadedError{edge: "work_for"}
}

// PaymentHistoryOrErr returns the PaymentHistory value or an error if the edge
// was not loaded in eager-loading.
func (e EMPLOYEEEdges) PaymentHistoryOrErr() ([]*PAYMENT_HISTORY, error) {
	if e.loadedTypes[3] {
		return e.PaymentHistory, nil
	}
	return nil, &NotLoadedError{edge: "payment_history"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*EMPLOYEE) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case employee.FieldPayroll:
			values[i] = new(sql.NullFloat64)
		case employee.FieldID, employee.FieldCurrency, employee.FieldEmploy:
			values[i] = new(sql.NullInt64)
		case employee.FieldGid, employee.FieldEmployerGid, employee.FieldName, employee.FieldPosition, employee.FieldWallet, employee.FieldEmail, employee.FieldWorkStart, employee.FieldWorkEnds, employee.FieldCreatedBy, employee.FieldUpdatedBy:
			values[i] = new(sql.NullString)
		case employee.FieldPayday, employee.FieldCreatedAt, employee.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		case employee.ForeignKeys[0]: // crypto_currency_employee_paid
			values[i] = new(sql.NullInt64)
		case employee.ForeignKeys[1]: // employer_user_info_work_under
			values[i] = new(sql.NullInt64)
		case employee.ForeignKeys[2]: // employ_type_employee_type_to
			values[i] = new(sql.NullInt64)
		default:
			return nil, fmt.Errorf("unexpected column %q for type EMPLOYEE", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the EMPLOYEE fields.
func (e *EMPLOYEE) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case employee.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			e.ID = int(value.Int64)
		case employee.FieldGid:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field gid", values[i])
			} else if value.Valid {
				e.Gid = value.String
			}
		case employee.FieldEmployerGid:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field employer_gid", values[i])
			} else if value.Valid {
				e.EmployerGid = value.String
			}
		case employee.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				e.Name = value.String
			}
		case employee.FieldPosition:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field position", values[i])
			} else if value.Valid {
				e.Position = value.String
			}
		case employee.FieldWallet:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field wallet", values[i])
			} else if value.Valid {
				e.Wallet = value.String
			}
		case employee.FieldPayroll:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field payroll", values[i])
			} else if value.Valid {
				e.Payroll = value.Float64
			}
		case employee.FieldCurrency:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field currency", values[i])
			} else if value.Valid {
				e.Currency = int(value.Int64)
			}
		case employee.FieldPayday:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field payday", values[i])
			} else if value.Valid {
				e.Payday = value.Time
			}
		case employee.FieldEmploy:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field employ", values[i])
			} else if value.Valid {
				e.Employ = int(value.Int64)
			}
		case employee.FieldEmail:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field email", values[i])
			} else if value.Valid {
				e.Email = value.String
			}
		case employee.FieldWorkStart:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field work_start", values[i])
			} else if value.Valid {
				e.WorkStart = value.String
			}
		case employee.FieldWorkEnds:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field work_ends", values[i])
			} else if value.Valid {
				e.WorkEnds = value.String
			}
		case employee.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				e.CreatedAt = value.Time
			}
		case employee.FieldCreatedBy:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field created_by", values[i])
			} else if value.Valid {
				e.CreatedBy = value.String
			}
		case employee.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				e.UpdatedAt = value.Time
			}
		case employee.FieldUpdatedBy:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field updated_by", values[i])
			} else if value.Valid {
				e.UpdatedBy = value.String
			}
		case employee.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field crypto_currency_employee_paid", value)
			} else if value.Valid {
				e.crypto_currency_employee_paid = new(int)
				*e.crypto_currency_employee_paid = int(value.Int64)
			}
		case employee.ForeignKeys[1]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field employer_user_info_work_under", value)
			} else if value.Valid {
				e.employer_user_info_work_under = new(int)
				*e.employer_user_info_work_under = int(value.Int64)
			}
		case employee.ForeignKeys[2]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field employ_type_employee_type_to", value)
			} else if value.Valid {
				e.employ_type_employee_type_to = new(int)
				*e.employ_type_employee_type_to = int(value.Int64)
			}
		}
	}
	return nil
}

// QueryEmployeeGets queries the "employee_gets" edge of the EMPLOYEE entity.
func (e *EMPLOYEE) QueryEmployeeGets() *CRYPTOCURRENCYQuery {
	return (&EMPLOYEEClient{config: e.config}).QueryEmployeeGets(e)
}

// QueryEmployeeTypeFrom queries the "employee_type_from" edge of the EMPLOYEE entity.
func (e *EMPLOYEE) QueryEmployeeTypeFrom() *EMPLOYTYPEQuery {
	return (&EMPLOYEEClient{config: e.config}).QueryEmployeeTypeFrom(e)
}

// QueryWorkFor queries the "work_for" edge of the EMPLOYEE entity.
func (e *EMPLOYEE) QueryWorkFor() *EMPLOYERUSERINFOQuery {
	return (&EMPLOYEEClient{config: e.config}).QueryWorkFor(e)
}

// QueryPaymentHistory queries the "payment_history" edge of the EMPLOYEE entity.
func (e *EMPLOYEE) QueryPaymentHistory() *PAYMENTHISTORYQuery {
	return (&EMPLOYEEClient{config: e.config}).QueryPaymentHistory(e)
}

// Update returns a builder for updating this EMPLOYEE.
// Note that you need to call EMPLOYEE.Unwrap() before calling this method if this EMPLOYEE
// was returned from a transaction, and the transaction was committed or rolled back.
func (e *EMPLOYEE) Update() *EMPLOYEEUpdateOne {
	return (&EMPLOYEEClient{config: e.config}).UpdateOne(e)
}

// Unwrap unwraps the EMPLOYEE entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (e *EMPLOYEE) Unwrap() *EMPLOYEE {
	_tx, ok := e.config.driver.(*txDriver)
	if !ok {
		panic("ent: EMPLOYEE is not a transactional entity")
	}
	e.config.driver = _tx.drv
	return e
}

// String implements the fmt.Stringer.
func (e *EMPLOYEE) String() string {
	var builder strings.Builder
	builder.WriteString("EMPLOYEE(")
	builder.WriteString(fmt.Sprintf("id=%v, ", e.ID))
	builder.WriteString("gid=")
	builder.WriteString(e.Gid)
	builder.WriteString(", ")
	builder.WriteString("employer_gid=")
	builder.WriteString(e.EmployerGid)
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(e.Name)
	builder.WriteString(", ")
	builder.WriteString("position=")
	builder.WriteString(e.Position)
	builder.WriteString(", ")
	builder.WriteString("wallet=")
	builder.WriteString(e.Wallet)
	builder.WriteString(", ")
	builder.WriteString("payroll=")
	builder.WriteString(fmt.Sprintf("%v", e.Payroll))
	builder.WriteString(", ")
	builder.WriteString("currency=")
	builder.WriteString(fmt.Sprintf("%v", e.Currency))
	builder.WriteString(", ")
	builder.WriteString("payday=")
	builder.WriteString(e.Payday.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("employ=")
	builder.WriteString(fmt.Sprintf("%v", e.Employ))
	builder.WriteString(", ")
	builder.WriteString("email=")
	builder.WriteString(e.Email)
	builder.WriteString(", ")
	builder.WriteString("work_start=")
	builder.WriteString(e.WorkStart)
	builder.WriteString(", ")
	builder.WriteString("work_ends=")
	builder.WriteString(e.WorkEnds)
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(e.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("created_by=")
	builder.WriteString(e.CreatedBy)
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(e.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_by=")
	builder.WriteString(e.UpdatedBy)
	builder.WriteByte(')')
	return builder.String()
}

// EMPLOYEEs is a parsable slice of EMPLOYEE.
type EMPLOYEEs []*EMPLOYEE

func (e EMPLOYEEs) config(cfg config) {
	for _i := range e {
		e[_i].config = cfg
	}
}
