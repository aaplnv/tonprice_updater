package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
)

type INRQuote struct {
	ent.Schema
}

func (INRQuote) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "INR"},
	}
}

func (INRQuote) Fields() []ent.Field {
	return ChartModel()
}

func (INRQuote) Edges() []ent.Edge {
	return nil
}
