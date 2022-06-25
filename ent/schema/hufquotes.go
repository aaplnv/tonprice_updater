package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
)

type HUFQuote struct {
	ent.Schema
}

func (HUFQuote) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "HUF"},
	}
}

func (HUFQuote) Fields() []ent.Field {
	return ChartModel()
}

func (HUFQuote) Edges() []ent.Edge {
	return nil
}
