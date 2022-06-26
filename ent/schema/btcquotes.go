package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
)

type BTCQuote struct {
	ent.Schema
}

func (BTCQuote) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "BTC"},
	}
}

func (BTCQuote) Fields() []ent.Field {
	return Quote()
}

func (BTCQuote) Edges() []ent.Edge {
	return nil
}
