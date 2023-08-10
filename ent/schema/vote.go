package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Vote holds the schema definition for the Vote entity.
type Vote struct {
	ent.Schema
}

// Mixin of the Vote.
func (Vote) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}

// Fields of the Vote.
func (Vote) Fields() []ent.Field {
	return []ent.Field{
		field.Bool("up_vote").Default(false),
		field.Bool("down_vote").Default(false),
		field.Int("user_id"),
		field.Int("comment_id"),
	}
}

// Edges of the Vote.
func (Vote) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("user", User.Type).Unique().Field("user_id").Required(),
		edge.To("comment", Comment.Type).Unique().Field("comment_id").Required(),
	}
}
