package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

func Quote() []ent.Field {
	return []ent.Field{
		field.Int("id").Unique(),
		field.Float("price"),
		field.Time("Timestamp"),
	}
}
