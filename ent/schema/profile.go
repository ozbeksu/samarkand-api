package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Profile holds the schema definition for the Profile entity.
type Profile struct {
	ent.Schema
}

// Mixin of the User.
func (Profile) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}

// Fields of the Profile.
func (Profile) Fields() []ent.Field {
	return []ent.Field{
		field.String("first_name").Optional(),
		field.String("last_name").Optional(),
		field.String("bio").Optional(),
		field.String("location").Optional(),
		field.Time("date_of_birth").Optional(),
		field.Int("avatar_id").Optional(),
		field.Int("cover_id").Optional(),
		field.Int("user_id"),
	}
}

// Edges of the Profile.
func (Profile) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("avatar", Attachment.Type).Unique().Field("avatar_id"),
		edge.To("cover", Attachment.Type).Unique().Field("cover_id"),
		edge.From("user", User.Type).Ref("profile").Unique().Required().Field("user_id"),
	}
}
