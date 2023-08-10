package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Content holds the schema definition for the Content entity.
type Content struct {
	ent.Schema
}

// Mixin of the Content.
func (Content) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}

// Fields of the Content.
func (Content) Fields() []ent.Field {
	return []ent.Field{
		field.Enum("type").Values("post", "poll").Default("post"),
		field.String("body").Optional(),
		field.Int("comment_id").Optional(),
	}
}

// Edges of the Content.
func (Content) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("attachments", Attachment.Type),
		edge.From("comment", Comment.Type).Ref("content").Unique().Field("comment_id"),
	}
}
