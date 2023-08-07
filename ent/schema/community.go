package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"time"
)

// Community holds the schema definition for the Community entity.
type Community struct {
	ent.Schema
}

// Fields of the Community.
func (Community) Fields() []ent.Field {
	return []ent.Field{
		field.String("title"),
		field.String("slug"),
		field.String("description"),
		field.Int("avatar_id").Optional(),
		field.Int("cover_id").Optional(),
		field.Int("creator_id"),
		field.Enum("access").Values("public", "restricted", "private").Default("public"),
		field.Time("created_at").Default(time.Now),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the Community.
func (Community) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("avatar", Attachment.Type).Unique().Field("avatar_id"),
		edge.To("cover", Attachment.Type).Unique().Field("cover_id"),
		edge.To("creator", User.Type).Unique().Required().Field("creator_id"),
		edge.To("comments", Comment.Type),
		edge.To("topics", Topic.Type),
		edge.To("members", User.Type),
		edge.To("moderators", User.Type),
	}
}
