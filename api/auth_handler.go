package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/ozbeksu/samarkand-api/ent"
	"github.com/ozbeksu/samarkand-api/store"
	"github.com/ozbeksu/samarkand-api/types"
	"github.com/ozbeksu/samarkand-api/utils"
	log "github.com/sirupsen/logrus"
	"os"
	"time"
)

type AuthHandler struct {
	store *store.Store
}

func NewAuthHandler(s *store.Store) *AuthHandler {
	return &AuthHandler{store: s}
}

func (h *AuthHandler) HandleAuthenticate(c *fiber.Ctx) error {
	var p types.AuthParams
	if err := c.BodyParser(&p); err != nil {
		log.Errorf("user post parameter parsing failed: %v", err)
		return err
	}

	ps := map[string]string{"email": p.Email}
	u, err := h.store.User.FindOne(ps)
	if err != nil {
		return InvalidCredentials(c)
	}

	ok := utils.IsValidPassword(u.Password, p.Password)
	if !ok {
		return InvalidCredentials(c)
	}

	return ResponseOk(AuthResponse{User: u, Token: createTokenFromUser(u)}, c)
}

func createTokenFromUser(u *ent.User) string {
	expires := time.Now().Add(time.Hour * 24 * 7).Unix()
	claims := jwt.MapClaims{"id": u.ID, "email": u.Email, "expires": expires}

	secret := os.Getenv("JWT_SECRET")
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signed, err := token.SignedString([]byte(secret))
	if err != nil {
		log.Warning("failed to sign token with secret: ", err)
	}

	return signed
}
