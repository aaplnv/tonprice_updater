package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
)

type AEDQuote struct {
	ent.Schema
}

func (AEDQuote) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "AED"},
	}
}

func (AEDQuote) Fields() []ent.Field {
	return Quote()
}

func (AEDQuote) Edges() []ent.Edge {
	return nil
}
