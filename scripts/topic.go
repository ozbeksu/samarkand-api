package main

import (
	"github.com/gosimple/slug"
	"github.com/ozbeksu/samarkand-api/ent"
	"strings"
)

func createTopic() *ent.TopicCreate {
	t := faker.LoremIpsumSentence(3)
	s := slug.Make(strings.ToLower(t))

	return db.Topic.Create().SetName(t).SetSlug(s)
}

func makeTopics(n int) []*ent.Topic {
	bulk := make([]*ent.TopicCreate, n)
	for i := 0; i < n; i++ {
		bulk[i] = createTopic()
	}
	return db.Topic.CreateBulk(bulk...).SaveX(ctx)
}
