package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
)

type PLNQuote struct {
	ent.Schema
}

func (PLNQuote) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "PLN"},
	}
}

func (PLNQuote) Fields() []ent.Field {
	return ChartModel()
}

func (PLNQuote) Edges() []ent.Edge {
	return nil
}
