package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
)

// ChartItem holds the schema definition for the ChartItem entity.
type EUROQuote struct {
	ent.Schema
}

func (EUROQuote) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "EURO"},
	}
}

// Fields of the ChartItem.
func (EUROQuote) Fields() []ent.Field {
	return ChartModel()
}

// Edges of the ChartItem.
func (EUROQuote) Edges() []ent.Edge {
	return nil
}
