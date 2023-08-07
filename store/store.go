package store

import (
	"github.com/ozbeksu/samarkand-api/utils"
	"strconv"
)

type Store struct {
	User      UserStore
	Tag       TagStore
	Topic     TopicStore
	Community CommunityStore
	Comment   CommentStore
}

func paginate(q map[string]string) (int, int) {
	page := 1
	limit := 10

	p, ok := q["page"]
	if ok {
		page, _ = strconv.Atoi(p)
		if page <= 0 {
			page = 1
		}
	}

	s, ok := q["size"]
	if ok {
		limit, _ = strconv.Atoi(s)
		limit = utils.Clamp(limit, 0, 100, 10)
	}

	l, ok := q["limit"]
	if ok {
		limit, _ = strconv.Atoi(l)
		limit = utils.Clamp(limit, 0, 100, 10)
	}

	return limit, (page - 1) * limit
}
