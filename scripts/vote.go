package main

import "github.com/ozbeksu/samarkand-api/ent"

func createVote(isUpVote bool, userID, commentID int) *ent.Vote {
	vote := db.Vote.Create()
	if isUpVote {
		vote = vote.SetUpVote(true)
	} else {
		vote = vote.SetDownVote(true)
	}

	return vote.SetUserID(userID).SetCommentID(commentID).SaveX(ctx)
}
