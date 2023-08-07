package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"time"
)

// Message holds the schema definition for the Message entity.
type Message struct {
	ent.Schema
}

// Fields of the Message.
func (Message) Fields() []ent.Field {
	return []ent.Field{
		field.String("subject").Optional(),
		field.String("body"),
		field.Int("sender_id"),
		field.Int("receiver_id"),
		field.Time("sent_at").Default(time.Now),
	}
}

// Edges of the Message.
func (Message) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("sender", User.Type).Unique().Required().Field("sender_id"),
		edge.To("receiver", User.Type).Unique().Required().Field("receiver_id"),
		edge.From("users", User.Type).Ref("messages"),
	}
}
