package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"time"
)

// Notification holds the schema definition for the Notification entity.
type Notification struct {
	ent.Schema
}

// Fields of the Notification.
func (Notification) Fields() []ent.Field {
	return []ent.Field{
		field.Enum("type").Values("info", "warning", "error", "success").Default("info"),
		field.String("content"),
		field.Int("user_id"),
		field.Enum("status").Values("sent", "received", "read").Default("sent"),
		field.Time("created_at").Default(time.Now),
	}
}

// Edges of the Notification.
func (Notification) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).Ref("notifications").Unique().Required().Field("user_id"),
	}
}
