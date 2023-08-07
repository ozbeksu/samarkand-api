package store

import (
	"context"
	"github.com/ozbeksu/samarkand-api/ent"
	"github.com/ozbeksu/samarkand-api/ent/comment"
	"github.com/ozbeksu/samarkand-api/ent/community"
)

type CommunityStore interface {
	Find(map[string]string) ([]*ent.Community, error)
	FindOne(map[string]string) (*ent.Community, error)
	FindOneWithComments(map[string]string, map[string]string) (*ent.Community, error)
}

type EntCommunityStore struct {
	client *ent.Client
	ctx    context.Context
}

func NewEntCommunityStore(client *ent.Client) *EntCommunityStore {
	return &EntCommunityStore{client: client, ctx: context.Background()}
}

func (e EntCommunityStore) Find(q map[string]string) ([]*ent.Community, error) {
	l, o := paginate(q)
	ts, err := e.client.Community.Query().
		WithAvatar(mediaQ).
		WithCover(mediaQ).
		WithTopics().
		WithComments(func(query *ent.CommentQuery) {
			query.
				Where(comment.ParentIDIsNil()).
				WithAuthor(userQ).
				WithContent(contentQ).
				WithComments(commentsQ).
				WithAttachments().
				WithTags().
				WithBookmarks().
				WithVotes()
		}).
		Order(ent.Asc(community.FieldCreatedAt)).
		Limit(l).
		Offset(o).
		All(e.ctx)
	if err != nil {
		return nil, err
	}

	return ts, nil
}

func (e EntCommunityStore) FindOne(w map[string]string) (*ent.Community, error) {
	t, err := e.client.Community.Query().
		Where(community.Slug(w["slug"])).
		WithAvatar(mediaQ).
		WithCover(mediaQ).
		WithTopics().
		First(e.ctx)
	if err != nil {
		return nil, err
	}

	return t, nil
}

func (e EntCommunityStore) FindOneWithComments(w map[string]string, q map[string]string) (*ent.Community, error) {
	t, err := e.client.Community.Query().
		Where(community.Slug(w["slug"])).
		WithAvatar(mediaQ).
		WithCover(mediaQ).
		WithTopics().
		WithComments(communityWithPaginatedCommentsQ(q)).
		First(e.ctx)
	if err != nil {
		return nil, err
	}

	return t, nil
}
