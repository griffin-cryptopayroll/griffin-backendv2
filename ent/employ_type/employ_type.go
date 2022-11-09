// Code generated by ent, DO NOT EDIT.

package employ_type

const (
	// Label holds the string label denoting the employ_type type in the database.
	Label = "employ_type"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldIsPermanent holds the string denoting the is_permanent field in the database.
	FieldIsPermanent = "is_permanent"
	// FieldPayFreq holds the string denoting the pay_freq field in the database.
	FieldPayFreq = "pay_freq"
	// EdgeEmployeeTypeTo holds the string denoting the employee_type_to edge name in mutations.
	EdgeEmployeeTypeTo = "employee_type_to"
	// Table holds the table name of the employ_type in the database.
	Table = "employ_type"
	// EmployeeTypeToTable is the table that holds the employee_type_to relation/edge.
	EmployeeTypeToTable = "employee"
	// EmployeeTypeToInverseTable is the table name for the EMPLOYEE entity.
	// It exists in this package in order to avoid circular dependency with the "employee" package.
	EmployeeTypeToInverseTable = "employee"
	// EmployeeTypeToColumn is the table column denoting the employee_type_to relation/edge.
	EmployeeTypeToColumn = "employ"
)

// Columns holds all SQL columns for employ_type fields.
var Columns = []string{
	FieldID,
	FieldIsPermanent,
	FieldPayFreq,
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
