package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ozbeksu/samarkand-api/store"
	log "github.com/sirupsen/logrus"
)

type CommentHandler struct {
	store *store.Store
}

func NewCommentHandler(s *store.Store) *CommentHandler {
	return &CommentHandler{store: s}
}

func (h *CommentHandler) HandleGetComments(c *fiber.Ctx) error {
	ts, err := h.store.Comment.Find(c.Queries())
	if err != nil {
		return NotFound(c)
	}

	return ResponseOk(ts, c)
}

func (h *CommentHandler) HandlePostComments(c *fiber.Ctx) error {
	return ResponseOk(nil, c)
}

func (h *CommentHandler) HandleGetComment(c *fiber.Ctx) error {
	ps := map[string]string{"slug": c.Params("slug")}
	t, err := h.store.Comment.FindOne(ps)
	if err != nil {
		log.Printf("comment find one querry err |> %s\n", err)
		return NotFound(c)
	}

	return ResponseOk(t, c)
}

func (h *CommentHandler) HandleGetCommentWithComments(c *fiber.Ctx) error {
	ps := map[string]string{"slug": c.Params("slug")}
	ts, err := h.store.Comment.FindOneWithComments(ps, c.Queries())
	if err != nil {
		return NotFound(c)
	}

	return ResponseOk(ts, c)
}

func (h *CommentHandler) HandlePostCommentUpVote(c *fiber.Ctx) error {
	ps := map[string]string{"slug": c.Params("slug"), "username": c.Params("username")}
	t, err := h.store.Comment.IncrementVote(ps)
	if err != nil {
		return err
	}

	return ResponseOk(t, c)
}

func (h *CommentHandler) HandlePostCommentDownVote(c *fiber.Ctx) error {
	ps := map[string]string{"slug": c.Params("slug"), "username": c.Params("username")}
	t, err := h.store.Comment.DecrementVote(ps)
	if err != nil {
		return err
	}

	return ResponseOk(t, c)
}

func (h *CommentHandler) HandlePostCommentBookmark(c *fiber.Ctx) error {
	ps := map[string]string{"slug": c.Params("slug"), "username": c.Params("username")}
	t, err := h.store.Comment.UpdateBookmark(ps)
	if err != nil {
		return err
	}

	return ResponseOk(t, c)
}
