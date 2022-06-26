package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
)

type RUBQuote struct {
	ent.Schema
}

func (RUBQuote) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "RUB"},
	}
}

func (RUBQuote) Fields() []ent.Field {
	return Quote()
}

func (RUBQuote) Edges() []ent.Edge {
	return nil
}
