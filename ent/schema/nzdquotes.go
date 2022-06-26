package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
)

type NZDQuote struct {
	ent.Schema
}

func (NZDQuote) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "NZD"},
	}
}

func (NZDQuote) Fields() []ent.Field {
	return Quote()
}

func (NZDQuote) Edges() []ent.Edge {
	return nil
}
