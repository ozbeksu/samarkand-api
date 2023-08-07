package main

import (
	"github.com/ozbeksu/samarkand-api/ent"
)

func createContent(t, body string) *ent.Content {
	switch t {
	case "poll":
		return createPollContent(body)
	default:
		return createPostContent(body)
	}
}

func createPostContent(body string) *ent.Content {
	if body == "" {
		body = faker.LoremIpsumParagraph(1, 3, 15, "<br/>\n")
	}
	return db.Content.Create().SetType("post").
		SetBody(body).
		SaveX(ctx)
}

func createPostContentWithAttachment(body string, imgID int) *ent.Content {
	if body == "" {
		body = faker.LoremIpsumParagraph(1, 3, 15, "<br/>\n")
	}
	return db.Content.Create().SetType("post").
		SetBody(body).
		AddAttachmentIDs(imgID).
		SaveX(ctx)
}

func createPollContent(body string) *ent.Content {
	if body == "" {
		body = faker.LoremIpsumParagraph(1, 3, 15, "<br/>\n")
	}
	return db.Content.Create().SetType("poll").
		SetBody(body).
		SaveX(ctx)
}
