// Code generated by entc, DO NOT EDIT.

package usdchart

const (
	// Label holds the string label denoting the usdchart type in the database.
	Label = "usd_chart"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldPrice holds the string denoting the price field in the database.
	FieldPrice = "price"
	// Table holds the table name of the usdchart in the database.
	Table = "USD"
)

// Columns holds all SQL columns for usdchart fields.
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
