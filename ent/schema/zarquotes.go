package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
)

type ZARQuote struct {
	ent.Schema
}

func (ZARQuote) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "ZAR"},
	}
}

func (ZARQuote) Fields() []ent.Field {
	return ChartModel()
}

func (ZARQuote) Edges() []ent.Edge {
	return nil
}
