package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ozbeksu/samarkand-api/store"
)

type TagHandler struct {
	store *store.Store
}

func NewTagHandler(s *store.Store) *TagHandler {
	return &TagHandler{store: s}
}

func (h *TagHandler) HandleGetTags(c *fiber.Ctx) error {
	ts, err := h.store.Tag.Find(c.Queries())
	if err != nil {
		return NotFound(c)
	}

	return ResponseOk(ts, c)
}

func (h *TagHandler) HandleGetTag(c *fiber.Ctx) error {
	ps := map[string]string{"slug": c.Params("slug")}
	t, err := h.store.Tag.FindOne(ps)
	if err != nil {
		return NotFound(c)
	}

	return ResponseOk(t, c)
}

func (h *TagHandler) HandleGetTagWithComments(c *fiber.Ctx) error {
	ps := map[string]string{"slug": c.Params("slug")}
	t, err := h.store.Tag.FindOneWithComments(ps, c.Queries())
	if err != nil {
		return NotFound(c)
	}

	return ResponseOk(t, c)
}
