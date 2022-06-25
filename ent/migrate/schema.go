// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// EUROColumns holds the columns for the "EURO" table.
	EUROColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "price", Type: field.TypeFloat64},
		{Name: "timestamp", Type: field.TypeTime},
	}
	// EUROTable holds the schema information for the "EURO" table.
	EUROTable = &schema.Table{
		Name:       "EURO",
		Columns:    EUROColumns,
		PrimaryKey: []*schema.Column{EUROColumns[0]},
	}
	// RUBColumns holds the columns for the "RUB" table.
	RUBColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "price", Type: field.TypeFloat64},
		{Name: "timestamp", Type: field.TypeTime},
	}
	// RUBTable holds the schema information for the "RUB" table.
	RUBTable = &schema.Table{
		Name:       "RUB",
		Columns:    RUBColumns,
		PrimaryKey: []*schema.Column{RUBColumns[0]},
	}
	// UAHColumns holds the columns for the "UAH" table.
	UAHColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "price", Type: field.TypeFloat64},
		{Name: "timestamp", Type: field.TypeTime},
	}
	// UAHTable holds the schema information for the "UAH" table.
	UAHTable = &schema.Table{
		Name:       "UAH",
		Columns:    UAHColumns,
		PrimaryKey: []*schema.Column{UAHColumns[0]},
	}
	// USDColumns holds the columns for the "USD" table.
	USDColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "price", Type: field.TypeFloat64},
		{Name: "timestamp", Type: field.TypeTime},
	}
	// USDTable holds the schema information for the "USD" table.
	USDTable = &schema.Table{
		Name:       "USD",
		Columns:    USDColumns,
		PrimaryKey: []*schema.Column{USDColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		EUROTable,
		RUBTable,
		UAHTable,
		USDTable,
	}
)

func init() {
	EUROTable.Annotation = &entsql.Annotation{
		Table: "EURO",
	}
	RUBTable.Annotation = &entsql.Annotation{
		Table: "RUB",
	}
	UAHTable.Annotation = &entsql.Annotation{
		Table: "UAH",
	}
	USDTable.Annotation = &entsql.Annotation{
		Table: "USD",
	}
}
