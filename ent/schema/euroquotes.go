package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
)

type EUROQuote struct {
	ent.Schema
}

func (EUROQuote) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "EURO"},
	}
}

func (EUROQuote) Fields() []ent.Field {
	return Quote()
}

func (EUROQuote) Edges() []ent.Edge {
	return nil
}
