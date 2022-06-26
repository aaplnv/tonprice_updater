package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
)

type CADQuote struct {
	ent.Schema
}

func (CADQuote) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "CAD"},
	}
}

func (CADQuote) Fields() []ent.Field {
	return Quote()
}

func (CADQuote) Edges() []ent.Edge {
	return nil
}
