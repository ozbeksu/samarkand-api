package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/ozbeksu/samarkand-api/utils"
	"time"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("username").Unique().DefaultFunc(func() string { return utils.RandStringBytes(8) }),
		field.String("email").Unique(),
		field.String("password").Sensitive(),
		field.Bool("active").Default(false),
		field.Bool("staff").Default(false),
		field.Time("created_at").Default(time.Now),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("profile", Profile.Type).Unique(),
		edge.To("groups", Group.Type),
		edge.To("comments", Comment.Type),
		edge.To("messages", Message.Type),
		edge.To("notifications", Notification.Type),
		edge.To("attachments", Attachment.Type),
		edge.To("following", User.Type).From("followers"),
		edge.From("votes", Vote.Type).Ref("user"),
		edge.From("bookmarks", Bookmark.Type).Ref("user"),
		edge.From("sent_messages", Message.Type).Ref("sender"),
		edge.From("received_messages", Message.Type).Ref("receiver"),
		edge.From("communities", Community.Type).Ref("members"),
		edge.From("moderating", Community.Type).Ref("moderators"),
	}
}
