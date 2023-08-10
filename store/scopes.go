package store

import (
	"github.com/ozbeksu/samarkand-api/ent"
	"github.com/ozbeksu/samarkand-api/ent/bookmark"
	"github.com/ozbeksu/samarkand-api/ent/comment"
	"github.com/ozbeksu/samarkand-api/ent/community"
	"github.com/ozbeksu/samarkand-api/ent/vote"
)

func userQ(q *ent.UserQuery) {
	q.
		WithProfile(profileQ)
}

func avatarQ(q *ent.ProfileQuery) {
	q.
		WithAvatar(func(query *ent.AttachmentQuery) {
			query.Select("url")
		})
}

func profileQ(q *ent.ProfileQuery) {
	q.
		WithAvatar(mediaQ).
		WithCover(mediaQ)
}

func followQ(q *ent.UserQuery) {
	q.
		Select("id")
}

func messageSQ(q *ent.MessageSenderQuery) {
	q.
		WithMessage(func(query *ent.MessageQuery) {
			query.WithRecipients(func(recipientQuery *ent.MessageRecipientQuery) {
				recipientQuery.
					WithCommunity(communityQ).
					WithUser(userQ)
			})
		})
}

func messageRQ(q *ent.MessageRecipientQuery) {
	q.
		WithMessage(func(query *ent.MessageQuery) {
			query.WithSender(func(senderQuery *ent.MessageSenderQuery) {
				senderQuery.
					WithCommunity(communityQ).
					WithUser(userQ)
			})
		})
}

func communityQ(q *ent.CommunityQuery) {
	q.
		WithAvatar(mediaQ).
		WithCover(mediaQ).
		WithTopics()
}

func bookmarkWithCommentQ(q *ent.BookmarkQuery) {
	q.
		WithComment(commentsQ)
}

func voteWithCommentQ(q *ent.VoteQuery) {
	q.
		WithComment(commentsQ)
}

func commentWithCommunityQ(q *ent.CommunityQuery) {
	q.
		Select("title", "slug").
		WithAvatar(mediaQ)
}

func commentsQ(q *ent.CommentQuery) {
	q.
		WithAuthor(userQ).
		WithContent(contentQ).
		WithVotes().
		WithComments(commentsWithCommentQ).
		Order(ent.Desc(comment.FieldCreatedAt)).
		Limit(5)
}

func commentsWithCommentQ(q *ent.CommentQuery) {
	q.
		WithAuthor(userQ).
		WithContent(contentQ).
		WithVotes().
		WithComments(commentsWithCommentWithCommentsQ).
		Order(ent.Desc(comment.FieldCreatedAt)).
		Limit(5)
}

func commentsWithCommentWithCommentsQ(q *ent.CommentQuery) {
	q.
		WithAuthor(userQ).
		WithContent(contentQ).
		WithVotes().
		WithComments().
		Order(ent.Desc(comment.FieldCreatedAt)).
		Limit(5)
}

func contentQ(q *ent.ContentQuery) {
	q.
		WithAttachments()
}

func mediaQ(q *ent.AttachmentQuery) {
	q.
		Select("url", "width", "height")
}

func fileQ(q *ent.AttachmentQuery) {
	q.
		Select("url", "file_name", "mime_type")
}

func commentWithParent(q *ent.CommentQuery) {
	q.
		Where(comment.ParentIDNotNil())
}

func commentWithoutParent(q *ent.CommentQuery) {
	q.
		Where(comment.ParentIDIsNil())
}

func commentWithCommunity(q *ent.CommentQuery) {
	q.
		Where(comment.CommunityIDNotNil())
}

func commentWithoutCommunity(q *ent.CommentQuery) {
	q.
		Where(comment.CommunityIDIsNil())
}

func commentWithBookmarked(q *ent.CommentQuery) {
	q.
		WithBookmarks(bookmarkSaved)
}

func commentWithUpVotes(q *ent.CommentQuery) {
	q.
		WithVotes(voteUpVoted)
}

func commentWithDownVotes(q *ent.CommentQuery) {
	q.
		WithVotes(voteDownVoted)
}

func bookmarkSaved(q *ent.BookmarkQuery) {
	q.
		Where(bookmark.SavedEQ(true))
}

func voteUpVoted(q *ent.VoteQuery) {
	q.
		Where(vote.UpVoteEQ(true))
}

func voteDownVoted(q *ent.VoteQuery) {
	q.
		Where(vote.DownVoteEQ(true))
}

func communityWithPaginatedCommentsQ(q map[string]string) func(query *ent.CommentQuery) {
	l, o := paginate(q)

	return func(query *ent.CommentQuery) {
		query.
			Where(comment.ParentIDIsNil()).
			WithAuthor(userQ).
			WithContent(contentQ).
			WithCommunity(commentWithCommunityQ).
			WithComments(commentsQ).
			WithAttachments().
			WithTags().
			WithBookmarks().
			WithVotes().
			Order(ent.Desc(community.FieldCreatedAt)).
			Limit(l).
			Offset(o)
	}
}

func commentWithPaginatedQ(q map[string]string) func(query *ent.CommentQuery) {
	l, o := paginate(q)
	return func(query *ent.CommentQuery) {
		query.
			WithAuthor(userQ).
			WithContent(contentQ).
			WithVotes().
			WithComments(commentsWithCommentQ).
			Order(ent.Desc(comment.FieldCreatedAt)).
			Limit(l).
			Offset(o)
	}
}

func tagWithPaginatedComments(q map[string]string) func(query *ent.CommentQuery) {
	l, o := paginate(q)

	return func(query *ent.CommentQuery) {
		query.
			WithComments(commentsQ).
			Order(ent.Desc(comment.FieldCreatedAt)).
			Limit(l).
			Offset(o)
	}
}

func topicWithPaginatedCommunitiesQ(q map[string]string) func(query *ent.CommunityQuery) {
	l, o := paginate(q)

	return func(query *ent.CommunityQuery) {
		query.
			WithAvatar(mediaQ).
			Order(ent.Desc(community.FieldCreatedAt)).
			Limit(l).
			Offset(o)
	}
}
