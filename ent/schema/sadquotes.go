package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
)

type SADQuote struct {
	ent.Schema
}

func (SADQuote) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "SAD"},
	}
}

func (SADQuote) Fields() []ent.Field {
	return ChartModel()
}

func (SADQuote) Edges() []ent.Edge {
	return nil
}
