package main

import (
	"github.com/gosimple/slug"
	"github.com/ozbeksu/samarkand-api/ent"
	"github.com/ozbeksu/samarkand-api/ent/messagesender"
	"strings"
)

func createMessage() *ent.MessageCreate {
	t := faker.LoremIpsumParagraph(1, 1, 6, "")
	s := slug.Make(strings.ToLower(t))
	c := faker.LoremIpsumParagraph(2, 5, 10, "<br/>\n")

	return db.Message.Create().
		SetSubject(t).
		SetSlug(s).
		SetContent(c)
}

func createMessageSender(mID, userCount, communityCount int) *ent.MessageSenderCreate {
	d := getRandDate()
	tx := db.MessageSender.Create().
		SetMessageID(mID).
		SetSentAt(d)

	if getRandBool() {
		tx = tx.SetUserID(getRandIntInRange(1, userCount)).
			SetType(messagesender.TypeUser)
	} else {
		tx = tx.SetCommunityID(getRandIntInRange(1, communityCount)).
			SetType(messagesender.TypeCommunity)
	}

	return tx
}

func createMessageRecipient(mID, userCount int) *ent.MessageRecipientCreate {
	d := getRandDate()
	return db.MessageRecipient.Create().
		SetMessageID(mID).
		SetUserID(getRandIntInRange(1, userCount)).
		SetReadAt(d)
}

func makeMessages(n, userCount, communityCount int) []*ent.Message {
	mBulk := make([]*ent.MessageCreate, n)
	for i := 0; i < n; i++ {
		mBulk[i] = createMessage()
	}
	messages := db.Message.CreateBulk(mBulk...).SaveX(ctx)

	msBulk := make([]*ent.MessageSenderCreate, n)
	for i := 0; i < n; i++ {
		msBulk[i] = createMessageSender(i+1, userCount, communityCount)
	}
	db.MessageSender.CreateBulk(msBulk...).SaveX(ctx)

	mrBulk := make([]*ent.MessageRecipientCreate, n)
	for i := 0; i < n; i++ {
		mrBulk[i] = createMessageRecipient(i+1, userCount)
	}
	db.MessageRecipient.CreateBulk(mrBulk...).SaveX(ctx)

	return messages
}
