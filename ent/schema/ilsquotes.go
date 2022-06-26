package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
)

type ILSQuote struct {
	ent.Schema
}

func (ILSQuote) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "ILS"},
	}
}

func (ILSQuote) Fields() []ent.Field {
	return Quote()
}

func (ILSQuote) Edges() []ent.Edge {
	return nil
}
