package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
)

type JPYQuote struct {
	ent.Schema
}

func (JPYQuote) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "JPY"},
	}
}

func (JPYQuote) Fields() []ent.Field {
	return ChartModel()
}

func (JPYQuote) Edges() []ent.Edge {
	return nil
}
