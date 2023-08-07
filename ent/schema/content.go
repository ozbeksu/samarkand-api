package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"time"
)

// Content holds the schema definition for the Content entity.
type Content struct {
	ent.Schema
}

// Fields of the Content.
func (Content) Fields() []ent.Field {
	return []ent.Field{
		field.Enum("type").Values("post", "poll").Default("post"),
		field.String("body").Optional(),
		field.Int("comment_id").Optional(),
		field.Time("created_at").Default(time.Now),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the Content.
func (Content) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("attachments", Attachment.Type),
		edge.From("comment", Comment.Type).Ref("content").Unique().Field("comment_id"),
	}
}
