package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/ozbeksu/samarkand-api/utils"
)

// Message holds the schema definition for the Message entity.
type Message struct {
	ent.Schema
}

// Fields of the Message.
func (Message) Fields() []ent.Field {
	return []ent.Field{
		field.String("slug").DefaultFunc(func() string { return utils.RandStringBytes(12) }),
		field.String("subject").Optional(),
		field.String("content"),
		field.Enum("type").Values("text", "media", "link").Default("text"),
		field.Enum("context").Values("inbox", "chat").Default("inbox"),
	}
}

// Edges of the Message.
func (Message) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("sender", MessageSender.Type).Unique(),
		edge.To("recipients", MessageRecipient.Type),
	}
}
