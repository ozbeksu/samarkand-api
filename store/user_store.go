package store

import (
	"context"
	"entgo.io/ent/dialect/sql"
	"github.com/ozbeksu/samarkand-api/ent"
	"github.com/ozbeksu/samarkand-api/ent/user"
	"github.com/ozbeksu/samarkand-api/types"
)

type UserStore interface {
	Find(map[string]string) ([]*ent.User, error)
	FindOne(map[string]string) (*ent.User, error)
	FindOneWithProfile(map[string]string, map[string]string) (*ent.User, error)
	FindOneWithMessages(map[string]string, map[string]string) (*ent.User, error)
	FindOneWithPosts(map[string]string, map[string]string) (*ent.User, error)
	FindOneWithTreads(map[string]string, map[string]string) (*ent.User, error)
	FindOneWithComments(map[string]string, map[string]string) (*ent.User, error)
	FindOneWithMedia(map[string]string, map[string]string) (*ent.User, error)
	FindOneWithBookmark(map[string]string, map[string]string) (*ent.User, error)
	FindOneWithUpVoted(map[string]string, map[string]string) (*ent.User, error)
	FindOneWithDownVoted(map[string]string, map[string]string) (*ent.User, error)
	Create(*types.AuthParams) (*ent.User, error)
}

type EntUserStore struct {
	client *ent.Client
	ctx    context.Context
}

func NewEntUserStore(client *ent.Client) *EntUserStore {
	return &EntUserStore{client: client, ctx: context.Background()}
}

func (s EntUserStore) Find(q map[string]string) ([]*ent.User, error) {
	l, o := paginate(q)
	us, err := s.client.User.Query().
		WithProfile(profileQ).
		WithComments(func(query *ent.CommentQuery) {
			query.Select("id")
		}).
		Order(user.ByCommentsCount(sql.OrderDesc())).
		Limit(l).
		Offset(o).
		All(s.ctx)
	if err != nil {
		return nil, err
	}

	return us, nil
}

func (s EntUserStore) FindOne(q map[string]string) (*ent.User, error) {
	tx := s.client.User.Query()

	email, ok := q["email"]
	if ok {
		tx.Where(user.Email(email))
	}

	username, ok := q["username"]
	if ok {
		tx.Where(user.Username(username))
	}

	u, err := tx.WithProfile(avatarQ).First(s.ctx)
	if err != nil {
		return nil, err
	}

	return u, nil
}

func (s EntUserStore) FindOneWithProfile(w map[string]string, q map[string]string) (*ent.User, error) {
	u, err := s.client.User.Query().
		Where(user.Username(w["username"])).
		WithProfile(profileQ).
		WithFollowers(userQ).
		WithFollowing(userQ).
		First(s.ctx)
	if err != nil {
		return nil, err
	}

	return u, nil
}
func (s EntUserStore) FindOneWithMessages(w map[string]string, q map[string]string) (*ent.User, error) {
	u, err := s.client.User.Query().
		Where(user.Username(w["username"])).
		WithSentMessages(messageSQ).
		WithReceivedMessages(messageRQ).
		First(s.ctx)
	if err != nil {
		return nil, err
	}

	return u, nil
}

func (s EntUserStore) FindOneWithPosts(w map[string]string, q map[string]string) (*ent.User, error) {
	u, err := s.client.User.Query().
		Where(user.Username(w["username"])).
		WithComments(commentWithoutParent, commentWithoutCommunity, commentWithPaginatedQ(q)).
		First(s.ctx)
	if err != nil {
		return nil, err
	}

	return u, nil
}

func (s EntUserStore) FindOneWithTreads(w map[string]string, q map[string]string) (*ent.User, error) {
	u, err := s.client.User.Query().
		Where(user.Username(w["username"])).
		WithComments(commentWithoutParent, commentWithCommunity, commentWithPaginatedQ(q)).
		First(s.ctx)
	if err != nil {
		return nil, err
	}

	return u, nil
}

func (s EntUserStore) FindOneWithComments(w map[string]string, q map[string]string) (*ent.User, error) {
	u, err := s.client.User.Query().
		Where(user.Username(w["username"])).
		WithComments(commentWithParent, commentWithCommunity, commentWithPaginatedQ(q)).
		First(s.ctx)
	if err != nil {
		return nil, err
	}

	return u, nil
}

func (s EntUserStore) FindOneWithMedia(w map[string]string, q map[string]string) (*ent.User, error) {
	u, err := s.client.User.Query().
		Where(user.Username(w["username"])).
		WithAttachments().
		First(s.ctx)
	if err != nil {
		return nil, err
	}

	return u, nil
}

func (s EntUserStore) FindOneWithBookmark(w map[string]string, q map[string]string) (*ent.User, error) {
	u, err := s.client.User.Query().
		Where(user.Username(w["username"])).
		WithBookmarks(bookmarkWithCommentQ).
		First(s.ctx)
	if err != nil {
		return nil, err
	}

	return u, nil
}

func (s EntUserStore) FindOneWithUpVoted(w map[string]string, q map[string]string) (*ent.User, error) {
	u, err := s.client.User.Query().
		Where(user.Username(w["username"])).
		WithVotes(voteUpVoted, voteWithCommentQ).
		First(s.ctx)
	if err != nil {
		return nil, err
	}

	return u, nil
}

func (s EntUserStore) FindOneWithDownVoted(w map[string]string, q map[string]string) (*ent.User, error) {
	u, err := s.client.User.Query().
		Where(user.Username(w["username"])).
		WithVotes(voteDownVoted, voteWithCommentQ).
		First(s.ctx)
	if err != nil {
		return nil, err
	}

	return u, nil
}

func (s EntUserStore) Create(p *types.AuthParams) (*ent.User, error) {
	pass, err := types.HashPassword(p.Password)
	if err != nil {
		return nil, err
	}

	u, err := s.client.User.Create().
		SetEmail(p.Email).
		SetPassword(pass).
		Save(s.ctx)
	if err != nil {
		return nil, err
	}

	return u, nil
}
