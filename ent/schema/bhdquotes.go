package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
)

type BHDQuote struct {
	ent.Schema
}

func (BHDQuote) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "BHD"},
	}
}

func (BHDQuote) Fields() []ent.Field {
	return Quote()
}

func (BHDQuote) Edges() []ent.Edge {
	return nil
}
