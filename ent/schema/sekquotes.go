package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
)

type SEKQuote struct {
	ent.Schema
}

func (SEKQuote) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "SEK"},
	}
}

func (SEKQuote) Fields() []ent.Field {
	return ChartModel()
}

func (SEKQuote) Edges() []ent.Edge {
	return nil
}
