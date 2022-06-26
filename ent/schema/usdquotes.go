package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
)

type USDQuote struct {
	ent.Schema
}

func (USDQuote) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "USD"},
	}
}

func (USDQuote) Fields() []ent.Field {
	return Quote()
}

func (USDQuote) Edges() []ent.Edge {
	return nil
}
