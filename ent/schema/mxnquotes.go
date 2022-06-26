package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
)

type MXNQuote struct {
	ent.Schema
}

func (MXNQuote) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "MXN"},
	}
}

func (MXNQuote) Fields() []ent.Field {
	return ChartModel()
}

func (MXNQuote) Edges() []ent.Edge {
	return nil
}
