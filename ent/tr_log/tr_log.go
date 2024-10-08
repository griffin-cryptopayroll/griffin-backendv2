// Code generated by ent, DO NOT EDIT.

package tr_log

const (
	// Label holds the string label denoting the tr_log type in the database.
	Label = "tr_log"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldTrType holds the string denoting the tr_type field in the database.
	FieldTrType = "tr_type"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// Table holds the table name of the tr_log in the database.
	Table = "tr_log"
)

// Columns holds all SQL columns for tr_log fields.
var Columns = []string{
	FieldID,
	FieldTrType,
	FieldCreatedAt,
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
