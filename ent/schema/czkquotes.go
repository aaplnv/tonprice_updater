package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
)

type CZKQuote struct {
	ent.Schema
}

func (CZKQuote) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "CZK"},
	}
}

func (CZKQuote) Fields() []ent.Field {
	return ChartModel()
}

func (CZKQuote) Edges() []ent.Edge {
	return nil
}
