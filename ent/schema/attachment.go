package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Attachment holds the schema definition for the Attachment entity.
type Attachment struct {
	ent.Schema
}

// Mixin of the Attachment.
func (Attachment) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
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
