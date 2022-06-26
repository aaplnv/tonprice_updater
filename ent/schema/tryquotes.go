package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
)

type TRYQuote struct {
	ent.Schema
}

func (TRYQuote) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "TRY"},
	}
}

func (TRYQuote) Fields() []ent.Field {
	return Quote()
}

func (TRYQuote) Edges() []ent.Edge {
	return nil
}
