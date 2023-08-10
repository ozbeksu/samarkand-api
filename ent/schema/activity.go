package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Activity holds the schema definition for the Activity entity.
type Activity struct {
	ent.Schema
}

// Mixin of the User.
func (Activity) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
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
	}
}

// Edges of the Activity.
func (Activity) Edges() []ent.Edge {
	return nil
}
