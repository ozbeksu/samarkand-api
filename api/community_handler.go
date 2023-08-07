package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ozbeksu/samarkand-api/store"
)

type CommunityHandler struct {
	store *store.Store
}

func NewCommunityHandler(s *store.Store) *CommunityHandler {
	return &CommunityHandler{store: s}
}

func (h *CommunityHandler) HandleGetCommunities(c *fiber.Ctx) error {
	ts, err := h.store.Community.Find(c.Queries())
	if err != nil {
		return NotFound(c)
	}

	return ResponseOk(ts, c)
}

func (h *CommunityHandler) HandleGetCommunity(c *fiber.Ctx) error {
	ps := map[string]string{"slug": c.Params("slug")}
	t, err := h.store.Community.FindOne(ps)
	if err != nil {
		return NotFound(c)
	}

	return ResponseOk(t, c)
}

func (h *CommunityHandler) HandleGetCommunityWithComments(c *fiber.Ctx) error {
	ps := map[string]string{"slug": c.Params("slug")}
	ts, err := h.store.Community.FindOneWithComments(ps, c.Queries())
	if err != nil {
		return NotFound(c)
	}

	return ResponseOk(ts, c)
}
