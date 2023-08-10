package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// PasswordResetToken holds the schema definition for the PasswordResetToken entity.
type PasswordResetToken struct {
	ent.Schema
}

// Indexes of the Authentication.
func (PasswordResetToken) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("token").Unique(),
	}
}

// Fields of the PasswordResetToken.
func (PasswordResetToken) Fields() []ent.Field {
	return []ent.Field{
		field.Int("user_id"),
		field.String("token"),
		field.Time("expiration_date").Optional(),
	}
}

// Edges of the PasswordResetToken.
func (PasswordResetToken) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("user", User.Type).Unique().Required().Field("user_id"),
	}
}
