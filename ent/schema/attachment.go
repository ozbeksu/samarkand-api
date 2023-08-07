package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"time"
)

// Attachment holds the schema definition for the Attachment entity.
type Attachment struct {
	ent.Schema
}

// Fields of the Attachment.
func (Attachment) Fields() []ent.Field {
	return []ent.Field{
		field.Enum("type").Values("avatar", "cover", "media", "file").Default("media"),
		field.String("file_name"),
		field.String("mime_type").Optional(),
		field.String("url"),
		field.Int("width").Optional(),
		field.Int("height").Optional(),
		field.Time("created_at").Default(time.Now),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the Attachment.
func (Attachment) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("comments", User.Type).Ref("attachments"),
		edge.From("contents", Content.Type).Ref("attachments"),
		edge.From("users", User.Type).Ref("attachments"),
	}
}
