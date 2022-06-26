package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
)

type GBPQuote struct {
	ent.Schema
}

func (GBPQuote) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "GBP"},
	}
}

func (GBPQuote) Fields() []ent.Field {
	return Quote()
}

func (GBPQuote) Edges() []ent.Edge {
	return nil
}
