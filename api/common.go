package api

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/ozbeksu/samarkand-api/ent"
	"net/http"
)

type Response struct {
	Data  any `json:"data"`
	Error any `json:"error"`
}

type AuthResponse struct {
	User  *ent.User `json:"user"`
	Token string    `json:"token"`
}

func NotFound(c *fiber.Ctx) error {
	return c.Status(http.StatusNotFound).JSON(Response{
		Data:  nil,
		Error: "not found",
	})
}

func InvalidCredentials(c *fiber.Ctx) error {
	return c.Status(http.StatusBadRequest).JSON(Response{
		Data:  nil,
		Error: "invalid credentials",
	})
}

func InvalidParameters(c *fiber.Ctx) error {
	return c.Status(http.StatusBadRequest).JSON(Response{
		Data:  nil,
		Error: "invalid parameters",
	})
}

func UnauthorizedAccess(c *fiber.Ctx) error {
	return c.Status(http.StatusUnauthorized).JSON(Response{
		Data:  nil,
		Error: "request unauthorized",
	})
}

func ResponseOk(d any, c *fiber.Ctx) error {
	return c.JSON(Response{
		Data:  d,
		Error: nil,
	})
}

func HandleErrors(c *fiber.Ctx, err error) error {
	var e *fiber.Error

	code := fiber.StatusInternalServerError
	if errors.As(err, &e) {
		code = e.Code
	}

	return c.Status(code).JSON(Response{
		Data:  nil,
		Error: err.Error(),
	})
}
