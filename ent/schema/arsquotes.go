package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
)

type ARSQuote struct {
	ent.Schema
}

func (ARSQuote) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "ARS"},
	}
}

func (ARSQuote) Fields() []ent.Field {
	return Quote()
}

func (ARSQuote) Edges() []ent.Edge {
	return nil
}
