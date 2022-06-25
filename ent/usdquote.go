// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"main/ent/usdquote"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
)

// USDQuote is the model entity for the USDQuote schema.
type USDQuote struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Price holds the value of the "price" field.
	Price float64 `json:"price,omitempty"`
	// Timestamp holds the value of the "Timestamp" field.
	Timestamp time.Time `json:"Timestamp,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*USDQuote) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case usdquote.FieldPrice:
			values[i] = new(sql.NullFloat64)
		case usdquote.FieldID:
			values[i] = new(sql.NullInt64)
		case usdquote.FieldTimestamp:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type USDQuote", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the USDQuote fields.
func (uq *USDQuote) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case usdquote.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			uq.ID = int(value.Int64)
		case usdquote.FieldPrice:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field price", values[i])
			} else if value.Valid {
				uq.Price = value.Float64
			}
		case usdquote.FieldTimestamp:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field Timestamp", values[i])
			} else if value.Valid {
				uq.Timestamp = value.Time
			}
		}
	}
	return nil
}

// Update returns a builder for updating this USDQuote.
// Note that you need to call USDQuote.Unwrap() before calling this method if this USDQuote
// was returned from a transaction, and the transaction was committed or rolled back.
func (uq *USDQuote) Update() *USDQuoteUpdateOne {
	return (&USDQuoteClient{config: uq.config}).UpdateOne(uq)
}

// Unwrap unwraps the USDQuote entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (uq *USDQuote) Unwrap() *USDQuote {
	tx, ok := uq.config.driver.(*txDriver)
	if !ok {
		panic("ent: USDQuote is not a transactional entity")
	}
	uq.config.driver = tx.drv
	return uq
}

// String implements the fmt.Stringer.
func (uq *USDQuote) String() string {
	var builder strings.Builder
	builder.WriteString("USDQuote(")
	builder.WriteString(fmt.Sprintf("id=%v", uq.ID))
	builder.WriteString(", price=")
	builder.WriteString(fmt.Sprintf("%v", uq.Price))
	builder.WriteString(", Timestamp=")
	builder.WriteString(uq.Timestamp.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// USDQuotes is a parsable slice of USDQuote.
type USDQuotes []*USDQuote

func (uq USDQuotes) config(cfg config) {
	for _i := range uq {
		uq[_i].config = cfg
	}
}
