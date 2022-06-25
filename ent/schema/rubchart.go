package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
)

// ChartItem holds the schema definition for the ChartItem entity.
type RUBChart struct {
	ent.Schema
}

func (RUBChart) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "RUB"},
	}
}

// Fields of the ChartItem.
func (RUBChart) Fields() []ent.Field {
	return ChartModel()
}

// Edges of the ChartItem.
func (RUBChart) Edges() []ent.Edge {
	return nil
}
