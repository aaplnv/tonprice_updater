package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
)

type BRLQuote struct {
	ent.Schema
}

func (BRLQuote) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "BRL"},
	}
}

func (BRLQuote) Fields() []ent.Field {
	return Quote()
}

func (BRLQuote) Edges() []ent.Edge {
	return nil
}
