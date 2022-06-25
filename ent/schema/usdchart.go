package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
)

// ChartItem holds the schema definition for the ChartItem entity.
type USDChart struct {
	ent.Schema
}

func (USDChart) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "USD"},
	}
}

// Fields of the ChartItem.
func (USDChart) Fields() []ent.Field {
	return ChartModel()
}

// Edges of the ChartItem.
func (USDChart) Edges() []ent.Edge {
	return nil
}
