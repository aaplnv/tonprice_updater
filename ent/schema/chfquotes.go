package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
)

type CHFQuote struct {
	ent.Schema
}

func (CHFQuote) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "CHF"},
	}
}

func (CHFQuote) Fields() []ent.Field {
	return ChartModel()
}

func (CHFQuote) Edges() []ent.Edge {
	return nil
}
