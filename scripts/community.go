package main

import (
	"github.com/gosimple/slug"
	"github.com/ozbeksu/samarkand-api/ent"
)

func createCommunity(topicCount, memberCount int) *ent.CommunityCreate {
	t := faker.LoremIpsumSentence(2)
	d := faker.LoremIpsumParagraph(2, 2, 5, "<br/>\n")
	creatorID := getRandIntInRange(1, memberCount)
	topicIDs := []int{getRandIntInRange(1, topicCount), getRandIntInRange(1, topicCount)}
	memberIDs := []int{getRandIntInRange(1, memberCount), getRandIntInRange(1, memberCount), getRandIntInRange(1, memberCount)}
	modID := getRandIntInRange(1, memberCount)
	a := createImage(getRandIntInRange(1, 10), "communities", "svg", 64, 64).SaveX(ctx)
	c := createImage(getRandIntInRange(1, 10), "covers", "jpg", 800, 392).SaveX(ctx)

	return db.Community.Create().
		SetTitle(t).
		SetSlug(slug.Make(t)).
		SetDescription(d).
		SetCreatorID(creatorID).
		SetAvatarID(a.ID).
		SetCoverID(c.ID).
		AddTopicIDs(topicIDs...).
		AddMemberIDs(memberIDs...).
		AddModeratorIDs(modID)
}

func makeCommunities(n int, topicCount, memberCount int) []*ent.Community {
	bulk := make([]*ent.CommunityCreate, n)
	for i := 0; i < n; i++ {
		bulk[i] = createCommunity(topicCount, memberCount)
	}
	return db.Community.CreateBulk(bulk...).SaveX(ctx)
}
