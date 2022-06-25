package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// ChartItem holds the schema definition for the ChartItem entity.
type ChartItem struct {
	ent.Schema
}

// Fields of the ChartItem.
func (ChartItem) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").Unique(),
		field.Float("price"),
	}
}

// Edges of the ChartItem.
func (ChartItem) Edges() []ent.Edge {
	return nil
}
