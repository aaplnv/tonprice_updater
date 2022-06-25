package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
)

type TWDQuote struct {
	ent.Schema
}

func (TWDQuote) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "TWD"},
	}
}

func (TWDQuote) Fields() []ent.Field {
	return ChartModel()
}

func (TWDQuote) Edges() []ent.Edge {
	return nil
}
