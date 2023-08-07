package store

import (
	"context"
	"github.com/ozbeksu/samarkand-api/ent"
	"github.com/ozbeksu/samarkand-api/ent/bookmark"
	"github.com/ozbeksu/samarkand-api/ent/comment"
	"github.com/ozbeksu/samarkand-api/ent/user"
	"github.com/ozbeksu/samarkand-api/ent/vote"
	"github.com/ozbeksu/samarkand-api/utils"
	"strconv"
	"time"
)

type CommentStore interface {
	Find(map[string]string) ([]*ent.Comment, error)
	FindOne(map[string]string) (*ent.Comment, error)
	FindOneWithComments(map[string]string, map[string]string) (*ent.Comment, error)
	IncrementVote(map[string]string) (*ent.Comment, error)
	DecrementVote(map[string]string) (*ent.Comment, error)
	UpdateBookmark(map[string]string) (*ent.Comment, error)
}

type EntCommentStore struct {
	client *ent.Client
	ctx    context.Context
}

func NewEntCommentStore(client *ent.Client) *EntCommentStore {
	return &EntCommentStore{client: client, ctx: context.Background()}
}

func (e EntCommentStore) Find(q map[string]string) ([]*ent.Comment, error) {
	limit, offset := paginate(q)
	tx := e.client.Comment.Query()

	parent, ok := q["parent"]
	if ok {
		if parent == "null" {
			tx = tx.Where(comment.ParentIDIsNil())
		} else {
			id, err := strconv.ParseInt(parent, 10, 64)
			if err == nil {
				tx = tx.Where(comment.ParentIDEQ(int(id)))
			}
		}
	}

	sort, ok := q["sort"]
	if ok {
		switch sort {
		case "best":
			tx = tx.Order(ent.Desc(comment.FieldBestScore))
			break
		case "hot":
			tx = tx.Order(ent.Desc(comment.FieldHotScore))
			break
		default:
			tx = tx.Order(ent.Desc(comment.FieldCreatedAt))
			break
		}
	}

	ts, err := tx.
		Where(comment.TitleNotNil()).
		WithAuthor(userQ).
		WithContent(contentQ).
		WithCommunity(commentWithCommunityQ).
		WithComments(commentsQ).
		WithAttachments().
		WithTags().
		WithBookmarks().
		WithVotes().
		Limit(limit).
		Offset(offset).
		All(e.ctx)
	if err != nil {
		return nil, err
	}

	return ts, nil
}

func (e EntCommentStore) FindOne(w map[string]string) (*ent.Comment, error) {
	t, err := e.client.Comment.Query().
		Where(comment.Slug(w["slug"])).
		WithAuthor(userQ).
		WithContent(contentQ).
		WithCommunity(communityQ).
		WithComments(commentsQ).
		WithTags().
		WithVotes().
		First(e.ctx)
	if err != nil {
		return nil, err
	}

	return t, nil
}

func (e EntCommentStore) FindOneWithComments(w map[string]string, q map[string]string) (*ent.Comment, error) {
	t, err := e.client.Comment.Query().
		Where(comment.Slug(w["slug"])).
		WithAuthor(userQ).
		WithContent(contentQ).
		WithCommunity(commentWithCommunityQ).
		WithComments(commentWithPaginatedQ(q)).
		WithAttachments().
		WithTags().
		WithVotes().
		First(e.ctx)
	if err != nil {
		return nil, err
	}

	return t, nil
}

func (e EntCommentStore) IncrementVote(q map[string]string) (*ent.Comment, error) {
	u, err := e.client.User.Query().Where(user.Username(q["username"])).First(e.ctx)
	if err != nil {
		return nil, err
	}

	c, err := e.client.Comment.Query().Where(comment.Slug(q["slug"])).First(e.ctx)
	if err != nil {
		return nil, err
	}

	v, err := e.client.Vote.Query().Where(vote.UserID(u.ID)).Where(vote.CommentID(c.ID)).First(e.ctx)
	upVotes, downVotes, err := e.incrementVote(u.ID, c, v)
	if err != nil {
		return nil, err
	}

	c, err = e.client.Comment.
		UpdateOneID(c.ID).
		SetUpVotes(upVotes).
		SetDownVotes(downVotes).
		SetHotScore(utils.HotScore(upVotes, downVotes, time.Now())).
		SetBestScore(utils.BestScore(upVotes, upVotes+downVotes)).
		Save(e.ctx)

	return c, nil
}

func (e EntCommentStore) DecrementVote(q map[string]string) (*ent.Comment, error) {
	u, err := e.client.User.Query().Where(user.Username(q["username"])).First(e.ctx)
	if err != nil {
		return nil, err
	}

	c, err := e.client.Comment.Query().Where(comment.Slug(q["slug"])).First(e.ctx)
	if err != nil {
		return nil, err
	}

	v, err := e.client.Vote.Query().Where(vote.UserID(u.ID)).Where(vote.CommentID(c.ID)).First(e.ctx)
	upVotes, downVotes, err := e.decrementVote(u.ID, c, v)
	if err != nil {
		return nil, err
	}

	c, err = e.client.Comment.
		UpdateOneID(c.ID).
		SetUpVotes(upVotes).
		SetDownVotes(downVotes).
		SetHotScore(utils.HotScore(upVotes, downVotes, time.Now())).
		SetBestScore(utils.BestScore(upVotes, upVotes+downVotes)).
		Save(e.ctx)

	return c, nil
}

func (e EntCommentStore) UpdateBookmark(q map[string]string) (*ent.Comment, error) {
	u, err := e.client.User.Query().Where(user.Username(q["username"])).First(e.ctx)
	if err != nil {
		return nil, err
	}

	c, err := e.client.Comment.Query().Where(comment.Slug(q["slug"])).First(e.ctx)
	if err != nil {
		return nil, err
	}

	b, err := e.client.Bookmark.Query().Where(bookmark.UserID(u.ID)).Where(bookmark.CommentID(c.ID)).First(e.ctx)
	if b == nil {
		_, sErr := e.client.Bookmark.Create().SetSaved(true).SetUserID(u.ID).SetCommentID(c.ID).Save(e.ctx)
		if sErr != nil {
			return nil, sErr
		}
	} else {
		_, sErr := e.client.Bookmark.UpdateOne(b).SetSaved(!b.Saved).Save(e.ctx)
		if sErr != nil {
			return nil, sErr
		}
	}

	return c, nil
}

func (e EntCommentStore) incrementVote(uID int, c *ent.Comment, v *ent.Vote) (int, int, error) {
	upVotes := c.UpVotes
	downVotes := c.DownVotes

	if v == nil {
		_, err := e.client.Vote.Create().SetUpVote(true).SetUserID(uID).SetCommentID(c.ID).Save(e.ctx)
		if err != nil {
			return 0, 0, err
		}
		upVotes += 1
	} else {
		tx := e.client.Vote.UpdateOne(v)
		if v.UpVote == true {
			tx = tx.SetUpVote(false).SetDownVote(false)
			upVotes -= 1
		} else if v.DownVote == true {
			tx = tx.SetUpVote(true).SetDownVote(false)
			upVotes += 1
			downVotes -= 1
		} else {
			tx = tx.SetUpVote(true).SetDownVote(false)
			upVotes += 1
		}
		_, err := tx.Save(e.ctx)
		if err != nil {
			return 0, 0, err
		}
	}
	return upVotes, downVotes, nil
}

func (e EntCommentStore) decrementVote(uID int, c *ent.Comment, v *ent.Vote) (int, int, error) {
	upVotes := c.UpVotes
	downVotes := c.DownVotes

	if v == nil {
		_, err := e.client.Vote.Create().SetDownVote(true).SetUserID(uID).SetCommentID(c.ID).Save(e.ctx)
		if err != nil {
			return 0, 0, err
		}
		downVotes += 1
	} else {
		tx := e.client.Vote.UpdateOne(v)
		if v.DownVote == true {
			tx = tx.SetDownVote(false).SetUpVote(false)
			downVotes -= 1
		} else if v.UpVote == true {
			tx = tx.SetDownVote(true).SetUpVote(false)
			downVotes += 1
			upVotes -= 1
		} else {
			tx = tx.SetDownVote(true).SetUpVote(false)
			downVotes += 1
		}
		_, err := tx.Save(e.ctx)
		if err != nil {
			return 0, 0, err
		}
	}
	return upVotes, downVotes, nil
}
