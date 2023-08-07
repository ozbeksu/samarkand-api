package main

import (
	"fmt"
	"github.com/ozbeksu/samarkand-api/ent"
	"github.com/ozbeksu/samarkand-api/ent/attachment"
	"mime"
)

func iToDigit(i int, ext string) string {
	var f string
	if i < 10 {
		f = fmt.Sprintf("0%d.%s", i, ext)
	} else {
		f = fmt.Sprintf("%d.%s", i, ext)
	}
	return f
}

func createAttachment(fileName, url string, t attachment.Type) *ent.AttachmentCreate {
	return db.Attachment.Create().SetType(t).SetFileName(fileName).SetURL(url)
}

func createImage(i int, folder, ext string, w, h int) *ent.AttachmentCreate {
	f := iToDigit(i, ext)
	u := fmt.Sprintf("/assets/%s/%s", folder, f)
	return createAttachment(f, u, attachment.TypeAvatar).SetMimeType(mime.TypeByExtension(ext)).SetWidth(w).SetHeight(h)
}

func makeImage(n int, folder, ext string, w, h int) []*ent.Attachment {
	bulk := make([]*ent.AttachmentCreate, n)
	for i := 0; i < n; i++ {
		bulk[i] = createImage(i, folder, ext, w, h)
	}
	return db.Attachment.CreateBulk(bulk...).SaveX(ctx)
}
