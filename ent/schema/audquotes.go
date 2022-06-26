package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
)

type AUDQuote struct {
	ent.Schema
}

func (AUDQuote) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "AUD"},
	}
}

func (AUDQuote) Fields() []ent.Field {
	return Quote()
}

func (AUDQuote) Edges() []ent.Edge {
	return nil
}
