package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
)

type PKRQuote struct {
	ent.Schema
}

func (PKRQuote) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "PKR"},
	}
}

func (PKRQuote) Fields() []ent.Field {
	return ChartModel()
}

func (PKRQuote) Edges() []ent.Edge {
	return nil
}
