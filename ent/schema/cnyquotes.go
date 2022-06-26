package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
)

type CNYQuote struct {
	ent.Schema
}

func (CNYQuote) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "CNY"},
	}
}

func (CNYQuote) Fields() []ent.Field {
	return Quote()
}

func (CNYQuote) Edges() []ent.Edge {
	return nil
}
