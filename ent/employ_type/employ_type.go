// Code generated by ent, DO NOT EDIT.

package employ_type

const (
	// Label holds the string label denoting the employ_type type in the database.
	Label = "employ_type"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldIsPermanent holds the string denoting the is_permanent field in the database.
	FieldIsPermanent = "is_permanent"
	// FieldContractStart holds the string denoting the contract_start field in the database.
	FieldContractStart = "contract_start"
	// FieldContractPeriod holds the string denoting the contract_period field in the database.
	FieldContractPeriod = "contract_period"
	// EdgeEmployeeTypeTo holds the string denoting the employee_type_to edge name in mutations.
	EdgeEmployeeTypeTo = "employee_type_to"
	// Table holds the table name of the employ_type in the database.
	Table = "employ_typ_es"
	// EmployeeTypeToTable is the table that holds the employee_type_to relation/edge.
	EmployeeTypeToTable = "employe_es"
	// EmployeeTypeToInverseTable is the table name for the EMPLOYEE entity.
	// It exists in this package in order to avoid circular dependency with the "employee" package.
	EmployeeTypeToInverseTable = "employe_es"
	// EmployeeTypeToColumn is the table column denoting the employee_type_to relation/edge.
	EmployeeTypeToColumn = "employ_type_employee_type_to"
)

// Columns holds all SQL columns for employ_type fields.
var Columns = []string{
	FieldID,
	FieldIsPermanent,
	FieldContractStart,
	FieldContractPeriod,
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
