package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ozbeksu/samarkand-api/store"
)

type TopicHandler struct {
	store *store.Store
}

func NewTopicHandler(s *store.Store) *TopicHandler {
	return &TopicHandler{store: s}
}

func (h *TopicHandler) HandleGetTopics(c *fiber.Ctx) error {
	ts, err := h.store.Topic.Find(c.Queries())
	if err != nil {
		return NotFound(c)
	}

	return ResponseOk(ts, c)
}

func (h *TopicHandler) HandleGetTopic(c *fiber.Ctx) error {
	ps := map[string]string{"slug": c.Params("slug")}
	t, err := h.store.Topic.FindOne(ps)
	if err != nil {
		return NotFound(c)
	}

	return ResponseOk(t, c)
}

func (h *TopicHandler) HandleGetTopicWithCommunities(c *fiber.Ctx) error {
	ps := map[string]string{"slug": c.Params("slug")}
	ts, err := h.store.Topic.FindOneWithCommunities(ps, c.Queries())
	if err != nil {
		return NotFound(c)
	}

	return ResponseOk(ts, c)
}
