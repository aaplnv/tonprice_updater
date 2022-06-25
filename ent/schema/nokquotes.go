package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
)

type NOKQuote struct {
	ent.Schema
}

func (NOKQuote) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "NOK"},
	}
}

func (NOKQuote) Fields() []ent.Field {
	return ChartModel()
}

func (NOKQuote) Edges() []ent.Edge {
	return nil
}
