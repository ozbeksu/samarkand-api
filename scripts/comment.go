package main

import (
	"github.com/gosimple/slug"
	"github.com/ozbeksu/samarkand-api/ent"
	"strings"
)

func createComment(userCount, communityCount, tagCount, treadMin, treadMax int, hasAttachment bool) *ent.CommentCreate {
	d := getRandDate()
	hs, bs := getScores(d)

	t := faker.LoremIpsumSentence(3)
	s := slug.Make(strings.ToLower(t))
	uID := getRandIntInRange(1, userCount)
	tagIDs := []int{getRandIntInRange(1, tagCount), getRandIntInRange(1, tagCount)}

	comment := db.Comment.Create().
		SetTitle(t).
		SetSlug(s).
		SetHotScore(hs).
		SetBestScore(bs).
		SetCreatedAt(d).
		SetAuthorID(uID).
		SetCommunityID(getRandIntInRange(1, communityCount)).
		AddTagIDs(tagIDs...)

	var co *ent.Content
	if hasAttachment {
		var a *ent.Attachment
		if getRandBool() {
			a = createImage(getRandIntInRange(1, 7), "posts", "jpg", 1000, 1000).SaveX(ctx)
		} else {
			a = createImage(getRandIntInRange(1, 6), "albums", "jpg", 600, 600).SaveX(ctx)
		}
		co = createPostContentWithAttachment("", a.ID)
		comment = comment.AddAttachmentIDs(a.ID)
		db.User.UpdateOneID(uID).AddAttachmentIDs(a.ID).SaveX(ctx)
	} else {
		co = createPostContent("")
	}

	if treadMin > 0 {
		comment.SetParentID(getRandIntInRange(treadMin, treadMax))
	}

	return comment.SetContentID(co.ID)
}

func makeComments(n, userCount, communityCount, tagCount, treadMin, treadMax int) []*ent.Comment {
	bulk := make([]*ent.CommentCreate, n)
	for i := 0; i < n; i++ {
		bulk[i] = createComment(userCount, communityCount, tagCount, treadMin, treadMax, getRandBool())
	}
	return db.Comment.CreateBulk(bulk...).SaveX(ctx)
}

func makeSubComments(n, userCount, communityCount, tagCount, treadMin, treadMax int) []*ent.Comment {
	bulk := make([]*ent.CommentCreate, n)
	for i := 0; i < n; i++ {
		bulk[i] = createComment(userCount, communityCount, tagCount, treadMin, treadMax, false)
	}
	return db.Comment.CreateBulk(bulk...).SaveX(ctx)
}
