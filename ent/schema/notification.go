package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/ozbeksu/samarkand-api/utils"
	"time"
)

// Notification holds the schema definition for the Notification entity.
type Notification struct {
	ent.Schema
}

// Fields of the Notification.
func (Notification) Fields() []ent.Field {
	return []ent.Field{
		field.String("slug").DefaultFunc(func() string { return utils.RandStringBytes(12) }),
		field.Enum("type").Values("message", "mention", "comment", "follow_request", "community_invite"),
		field.String("content"),
		field.Int("reference_id"),
		field.Int("user_id"),
		field.Enum("status").Values("sent", "received", "read", "actioned").Default("sent"),
		field.Time("created_at").Default(time.Now),
	}
}

// Edges of the Notification.
func (Notification) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).Ref("notifications").Unique().Required().Field("user_id"),
	}
}
