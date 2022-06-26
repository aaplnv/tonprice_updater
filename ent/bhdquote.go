// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"main/ent/bhdquote"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
)

// BHDQuote is the model entity for the BHDQuote schema.
type BHDQuote struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Price holds the value of the "price" field.
	Price float64 `json:"price,omitempty"`
	// Timestamp holds the value of the "Timestamp" field.
	Timestamp time.Time `json:"Timestamp,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*BHDQuote) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case bhdquote.FieldPrice:
			values[i] = new(sql.NullFloat64)
		case bhdquote.FieldID:
			values[i] = new(sql.NullInt64)
		case bhdquote.FieldTimestamp:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type BHDQuote", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the BHDQuote fields.
func (bq *BHDQuote) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case bhdquote.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			bq.ID = int(value.Int64)
		case bhdquote.FieldPrice:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field price", values[i])
			} else if value.Valid {
				bq.Price = value.Float64
			}
		case bhdquote.FieldTimestamp:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field Timestamp", values[i])
			} else if value.Valid {
				bq.Timestamp = value.Time
			}
		}
	}
	return nil
}

// Update returns a builder for updating this BHDQuote.
// Note that you need to call BHDQuote.Unwrap() before calling this method if this BHDQuote
// was returned from a transaction, and the transaction was committed or rolled back.
func (bq *BHDQuote) Update() *BHDQuoteUpdateOne {
	return (&BHDQuoteClient{config: bq.config}).UpdateOne(bq)
}

// Unwrap unwraps the BHDQuote entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (bq *BHDQuote) Unwrap() *BHDQuote {
	tx, ok := bq.config.driver.(*txDriver)
	if !ok {
		panic("ent: BHDQuote is not a transactional entity")
	}
	bq.config.driver = tx.drv
	return bq
}

// String implements the fmt.Stringer.
func (bq *BHDQuote) String() string {
	var builder strings.Builder
	builder.WriteString("BHDQuote(")
	builder.WriteString(fmt.Sprintf("id=%v", bq.ID))
	builder.WriteString(", price=")
	builder.WriteString(fmt.Sprintf("%v", bq.Price))
	builder.WriteString(", Timestamp=")
	builder.WriteString(bq.Timestamp.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// BHDQuotes is a parsable slice of BHDQuote.
type BHDQuotes []*BHDQuote

func (bq BHDQuotes) config(cfg config) {
	for _i := range bq {
		bq[_i].config = cfg
	}
}