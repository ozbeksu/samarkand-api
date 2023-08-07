package store

import (
	"context"
	"entgo.io/ent/dialect/sql"
	"github.com/ozbeksu/samarkand-api/ent"
	"github.com/ozbeksu/samarkand-api/ent/tag"
)

type TagStore interface {
	Find(map[string]string) ([]*ent.Tag, error)
	FindOne(map[string]string) (*ent.Tag, error)
	FindOneWithComments(map[string]string, map[string]string) (*ent.Tag, error)
}

type EntTagStore struct {
	client *ent.Client
	ctx    context.Context
}

func NewEntTagStore(client *ent.Client) *EntTagStore {
	return &EntTagStore{client: client, ctx: context.Background()}
}

func (e EntTagStore) Find(q map[string]string) ([]*ent.Tag, error) {
	l, o := paginate(q)
	ts, err := e.client.Tag.Query().
		WithComments(func(query *ent.CommentQuery) {
			query.Select("id")
		}).
		Order(tag.ByCommentsCount(sql.OrderDesc())).
		Limit(l).
		Offset(o).
		All(e.ctx)
	if err != nil {
		return nil, err
	}

	return ts, nil
}

func (e EntTagStore) FindOne(w map[string]string) (*ent.Tag, error) {
	t, err := e.client.Tag.Query().
		Where(tag.Slug(w["slug"])).
		First(e.ctx)
	if err != nil {
		return nil, err
	}

	return t, nil
}

func (e EntTagStore) FindOneWithComments(w map[string]string, q map[string]string) (*ent.Tag, error) {
	t, err := e.client.Tag.Query().
		Where(tag.Slug(w["slug"])).
		WithComments(tagWithPaginatedComments(q)).
		First(e.ctx)
	if err != nil {
		return nil, err
	}

	return t, nil
}
