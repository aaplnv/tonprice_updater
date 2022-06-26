package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
)

type IDRQuote struct {
	ent.Schema
}

func (IDRQuote) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "IDR"},
	}
}

func (IDRQuote) Fields() []ent.Field {
	return Quote()
}

func (IDRQuote) Edges() []ent.Edge {
	return nil
}
