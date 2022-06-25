// Code generated by entc, DO NOT EDIT.

package chartitem

const (
	// Label holds the string label denoting the chartitem type in the database.
	Label = "chart_item"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldPrice holds the string denoting the price field in the database.
	FieldPrice = "price"
	// Table holds the table name of the chartitem in the database.
	Table = "chart_items"
)

// Columns holds all SQL columns for chartitem fields.
var Columns = []string{
	FieldID,
	FieldPrice,
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
