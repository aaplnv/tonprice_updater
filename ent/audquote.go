// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"main/ent/audquote"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
)

// AUDQuote is the model entity for the AUDQuote schema.
type AUDQuote struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Price holds the value of the "price" field.
	Price float64 `json:"price,omitempty"`
	// Timestamp holds the value of the "Timestamp" field.
	Timestamp time.Time `json:"Timestamp,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*AUDQuote) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case audquote.FieldPrice:
			values[i] = new(sql.NullFloat64)
		case audquote.FieldID:
			values[i] = new(sql.NullInt64)
		case audquote.FieldTimestamp:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type AUDQuote", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the AUDQuote fields.
func (aq *AUDQuote) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case audquote.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			aq.ID = int(value.Int64)
		case audquote.FieldPrice:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field price", values[i])
			} else if value.Valid {
				aq.Price = value.Float64
			}
		case audquote.FieldTimestamp:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field Timestamp", values[i])
			} else if value.Valid {
				aq.Timestamp = value.Time
			}
		}
	}
	return nil
}

// Update returns a builder for updating this AUDQuote.
// Note that you need to call AUDQuote.Unwrap() before calling this method if this AUDQuote
// was returned from a transaction, and the transaction was committed or rolled back.
func (aq *AUDQuote) Update() *AUDQuoteUpdateOne {
	return (&AUDQuoteClient{config: aq.config}).UpdateOne(aq)
}

// Unwrap unwraps the AUDQuote entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (aq *AUDQuote) Unwrap() *AUDQuote {
	tx, ok := aq.config.driver.(*txDriver)
	if !ok {
		panic("ent: AUDQuote is not a transactional entity")
	}
	aq.config.driver = tx.drv
	return aq
}

// String implements the fmt.Stringer.
func (aq *AUDQuote) String() string {
	var builder strings.Builder
	builder.WriteString("AUDQuote(")
	builder.WriteString(fmt.Sprintf("id=%v", aq.ID))
	builder.WriteString(", price=")
	builder.WriteString(fmt.Sprintf("%v", aq.Price))
	builder.WriteString(", Timestamp=")
	builder.WriteString(aq.Timestamp.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// AUDQuotes is a parsable slice of AUDQuote.
type AUDQuotes []*AUDQuote

func (aq AUDQuotes) config(cfg config) {
	for _i := range aq {
		aq[_i].config = cfg
	}
}
