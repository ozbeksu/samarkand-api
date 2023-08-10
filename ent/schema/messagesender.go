package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// MessageSender holds the schema definition for the MessageSender entity.
type MessageSender struct {
	ent.Schema
}

// Fields of the MessageSender.
func (MessageSender) Fields() []ent.Field {
	return []ent.Field{
		field.Enum("sender_type").Values("user", "group", "community"),
		field.Int("message_id").Optional(),
		field.Int("user_id").Optional(),
		field.Int("group_id").Optional(),
		field.Int("community_id").Optional(),
	}
}

// Edges of the MessageSender.
func (MessageSender) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("message", Message.Type).Ref("sender").Unique().Field("message_id"),
		edge.To("user", User.Type).Unique().Field("user_id"),
		edge.To("group", Group.Type).Unique().Field("group_id"),
		edge.To("community", Community.Type).Unique().Field("community_id"),
	}
}
