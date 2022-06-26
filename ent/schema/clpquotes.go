package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
)

type CLPQuote struct {
	ent.Schema
}

func (CLPQuote) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "CLP"},
	}
}

func (CLPQuote) Fields() []ent.Field {
	return Quote()
}

func (CLPQuote) Edges() []ent.Edge {
	return nil
}
