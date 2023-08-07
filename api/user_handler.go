package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ozbeksu/samarkand-api/store"
	"github.com/ozbeksu/samarkand-api/types"
	log "github.com/sirupsen/logrus"
)

type UserHandler struct {
	store *store.Store
}

func NewUserHandler(s *store.Store) *UserHandler {
	return &UserHandler{store: s}
}

func (h *UserHandler) HandleGetUsers(c *fiber.Ctx) error {
	us, err := h.store.User.Find(c.Queries())
	if err != nil {
		return NotFound(c)
	}

	return ResponseOk(us, c)
}

func (h *UserHandler) HandlePostUsers(c *fiber.Ctx) error {
	var p types.AuthParams
	if err := c.BodyParser(&p); err != nil {
		log.Errorf("user post parameter parsing failed: %v", err)
		return InvalidParameters(c)
	}

	m := p.Validate()
	if len(m) > 0 {
		return c.JSON(m)
	}

	u, err := h.store.User.Create(&p)
	if err != nil {
		log.Errorf("user create failed: %v", err)
		return InvalidParameters(c)
	}

	return ResponseOk(AuthResponse{User: u, Token: createTokenFromUser(u)}, c)
}

func (h *UserHandler) HandleGetUser(c *fiber.Ctx) error {
	ps := map[string]string{"username": c.Params("username")}
	u, err := h.store.User.FindOne(ps)
	if err != nil {
		return NotFound(c)
	}

	return ResponseOk(u, c)
}

func (h *UserHandler) HandleGetUserWithPosts(c *fiber.Ctx) error {
	ps := map[string]string{"username": c.Params("username")}
	u, err := h.store.User.FindOneWithPosts(ps, c.Queries())
	if err != nil {
		return NotFound(c)
	}

	return ResponseOk(u, c)
}

func (h *UserHandler) HandleGetUserWithTreads(c *fiber.Ctx) error {
	ps := map[string]string{"username": c.Params("username")}
	u, err := h.store.User.FindOneWithTreads(ps, c.Queries())
	if err != nil {
		return NotFound(c)
	}

	return ResponseOk(u, c)
}

func (h *UserHandler) HandleGetUserWithComments(c *fiber.Ctx) error {
	ps := map[string]string{"username": c.Params("username")}
	u, err := h.store.User.FindOneWithComments(ps, c.Queries())
	if err != nil {
		return NotFound(c)
	}

	return ResponseOk(u, c)
}

func (h *UserHandler) HandleGetUserWithMedia(c *fiber.Ctx) error {
	ps := map[string]string{"username": c.Params("username")}
	u, err := h.store.User.FindOneWithMedia(ps, c.Queries())
	if err != nil {
		return NotFound(c)
	}

	return ResponseOk(u, c)
}

func (h *UserHandler) HandleGetUserWithBookmark(c *fiber.Ctx) error {
	ps := map[string]string{"username": c.Params("username")}
	u, err := h.store.User.FindOneWithBookmark(ps, c.Queries())
	if err != nil {
		return NotFound(c)
	}

	return ResponseOk(u, c)
}

func (h *UserHandler) HandleGetUserWithUpVoted(c *fiber.Ctx) error {
	ps := map[string]string{"username": c.Params("username")}
	u, err := h.store.User.FindOneWithUpVoted(ps, c.Queries())
	if err != nil {
		return NotFound(c)
	}

	return ResponseOk(u, c)
}

func (h *UserHandler) HandleGetUserWithDownVoted(c *fiber.Ctx) error {
	ps := map[string]string{"username": c.Params("username")}
	u, err := h.store.User.FindOneWithDownVoted(ps, c.Queries())
	if err != nil {
		return NotFound(c)
	}

	return ResponseOk(u, c)
}
