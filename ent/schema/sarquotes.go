package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
)

type SARQuote struct {
	ent.Schema
}

func (SARQuote) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "SAR"},
	}
}

func (SARQuote) Fields() []ent.Field {
	return Quote()
}

func (SARQuote) Edges() []ent.Edge {
	return nil
}
