package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
)

type HKDQuote struct {
	ent.Schema
}

func (HKDQuote) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "HKD"},
	}
}

func (HKDQuote) Fields() []ent.Field {
	return ChartModel()
}

func (HKDQuote) Edges() []ent.Edge {
	return nil
}
