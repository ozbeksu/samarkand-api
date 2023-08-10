package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// MessageRecipient holds the schema definition for the MessageRecipient entity.
type MessageRecipient struct {
	ent.Schema
}

// Fields of the MessageRecipient.
func (MessageRecipient) Fields() []ent.Field {
	return []ent.Field{
		field.Enum("status").Values("new", "read", "unread").Default("new"),
		field.Bool("archived").Default(false),
		field.Bool("deleted").Default(false),
		field.Int("message_id").Optional(),
		field.Int("user_id").Optional(),
		field.Int("group_id").Optional(),
		field.Int("community_id").Optional(),
	}
}

// Edges of the MessageRecipient.
func (MessageRecipient) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("message", Message.Type).Ref("recipients").Unique().Field("message_id"),
		edge.To("user", User.Type).Unique().Field("user_id"),
		edge.To("group", Group.Type).Unique().Field("group_id"),
		edge.To("community", Community.Type).Unique().Field("community_id"),
	}
}
