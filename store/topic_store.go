package store

import (
	"context"
	"github.com/ozbeksu/samarkand-api/ent"
	"github.com/ozbeksu/samarkand-api/ent/topic"
)

type TopicStore interface {
	Find(map[string]string) ([]*ent.Topic, error)
	FindOne(map[string]string) (*ent.Topic, error)
	FindOneWithCommunities(map[string]string, map[string]string) (*ent.Topic, error)
}

type EntTopicStore struct {
	client *ent.Client
	ctx    context.Context
}

func NewEntTopicStore(client *ent.Client) *EntTopicStore {
	return &EntTopicStore{client: client, ctx: context.Background()}
}

func (e EntTopicStore) Find(q map[string]string) ([]*ent.Topic, error) {
	l, o := paginate(q)
	ts, err := e.client.Topic.Query().
		Limit(l).
		Offset(o).
		All(e.ctx)
	if err != nil {
		return nil, err
	}

	return ts, nil
}

func (e EntTopicStore) FindOne(w map[string]string) (*ent.Topic, error) {
	t, err := e.client.Topic.Query().
		Where(topic.Slug(w["slug"])).
		First(e.ctx)
	if err != nil {
		return nil, err
	}

	return t, nil
}

func (e EntTopicStore) FindOneWithCommunities(w map[string]string, q map[string]string) (*ent.Topic, error) {
	t, err := e.client.Topic.Query().
		Where(topic.Slug(w["slug"])).
		WithCommunities(topicWithPaginatedCommunitiesQ(q)).
		First(e.ctx)
	if err != nil {
		return nil, err
	}

	return t, nil
}
