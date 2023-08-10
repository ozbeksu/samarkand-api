package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/ozbeksu/samarkand-api/utils"
)

// Comment holds the schema definition for the Comment entity.
type Comment struct {
	ent.Schema
}

// Mixin of the Comment.
func (Comment) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}

// Fields of the Comment.
func (Comment) Fields() []ent.Field {
	return []ent.Field{
		field.String("title").Optional(),
		field.String("slug").DefaultFunc(func() string { return utils.RandStringBytes(12) }),
		field.Int("parent_id").Optional(),
		field.Int("community_id").Optional(),
		field.Int("author_id"),
		field.Float("hot_score"),
		field.Float("best_score"),
		field.Int("up_votes").Default(0),
		field.Int("down_votes").Default(0),
		field.Bool("is_moderated").Default(false),
	}
}

// Edges of the Comment.
func (Comment) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("content", Content.Type).Unique().Required(),
		edge.To("tags", Tag.Type),
		edge.To("attachments", Attachment.Type),
		edge.To("comments", Comment.Type).From("parent").Unique().Field("parent_id"),
		edge.From("votes", Vote.Type).Ref("comment"),
		edge.From("bookmarks", Bookmark.Type).Ref("comment"),
		edge.From("community", Community.Type).Ref("comments").Unique().Field("community_id"),
		edge.From("author", User.Type).Ref("comments").Unique().Required().Field("author_id"),
	}
}
