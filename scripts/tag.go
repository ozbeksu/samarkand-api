package main

import (
	"github.com/gosimple/slug"
	"github.com/ozbeksu/samarkand-api/ent"
	"strings"
)

func createTag() *ent.TagCreate {
	t := faker.LoremIpsumWord()
	s := slug.Make(strings.ToLower(t))

	return db.Tag.Create().SetName(t).SetSlug(s)
}

func makeTags(n int) []*ent.Tag {
	bulk := make([]*ent.TagCreate, n)
	for i := 0; i < n; i++ {
		bulk[i] = createTag()
	}
	return db.Tag.CreateBulk(bulk...).SaveX(ctx)
}
