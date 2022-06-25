// Code generated by entc, DO NOT EDIT.

package uahquote

const (
	// Label holds the string label denoting the uahquote type in the database.
	Label = "uah_quote"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldPrice holds the string denoting the price field in the database.
	FieldPrice = "price"
	// FieldTimestamp holds the string denoting the timestamp field in the database.
	FieldTimestamp = "timestamp"
	// Table holds the table name of the uahquote in the database.
	Table = "UAH"
)

// Columns holds all SQL columns for uahquote fields.
var Columns = []string{
	FieldID,
	FieldPrice,
	FieldTimestamp,
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
