package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
)

type UAHQuote struct {
	ent.Schema
}

func (UAHQuote) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "UAH"},
	}
}

// Fields of the ChartItem.
func (UAHQuote) Fields() []ent.Field {
	return ChartModel()
}

// Edges of the ChartItem.
func (UAHQuote) Edges() []ent.Edge {
	return nil
}
