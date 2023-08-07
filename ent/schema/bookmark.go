package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"time"
)

// Bookmark holds the schema definition for the Bookmark entity.
type Bookmark struct {
	ent.Schema
}

// Fields of the Bookmark.
func (Bookmark) Fields() []ent.Field {
	return []ent.Field{
		field.Bool("saved").Default(false),
		field.Int("user_id"),
		field.Int("comment_id"),
		field.Time("created_at").Default(time.Now),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the Bookmark.
func (Bookmark) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("user", User.Type).Unique().Field("user_id").Required(),
		edge.To("comment", Comment.Type).Unique().Field("comment_id").Required(),
	}
}
