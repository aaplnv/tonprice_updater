package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

func ChartModel() []ent.Field {
	return []ent.Field{
		field.Int("id").Unique(),
		field.Float("price"),
		field.Time("Timestamp"),
	}
}
