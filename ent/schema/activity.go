package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"time"
)

// Activity holds the schema definition for the Activity entity.
type Activity struct {
	ent.Schema
}

// Fields of the Activity.
func (Activity) Fields() []ent.Field {
	return []ent.Field{
		field.String("log").Optional(),
		field.String("description"),
		field.String("causer_id").Optional(),
		field.String("causer_type").Optional(),
		field.String("subject_id").Optional(),
		field.String("subject_type").Optional(),
		field.JSON("data", map[string]any{}),
		field.Time("created_at").Default(time.Now),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the Activity.
func (Activity) Edges() []ent.Edge {
	return nil
}
