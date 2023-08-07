package middleware

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/ozbeksu/samarkand-api/store"
	"os"
	"time"
)

func JWTAuthenticate(us store.UserStore) fiber.Handler {
	return func(c *fiber.Ctx) error {
		token, ok := c.GetReqHeaders()["X-Api-Token"]
		if !ok {
			token = c.Cookies("token")
			if len(token) < 1 {
				return fmt.Errorf("unauthorized")
			}
		}

		claims, err := validateToken(token)
		if err != nil {
			return err
		}

		expires := claims["expires"].(float64)
		if float64(time.Now().Unix()) > expires {
			return fmt.Errorf("token expired")
		}

		id := claims["id"].(float64)
		u, err := us.FindOne(map[string]string{"id": fmt.Sprintf("%d", int(id))})
		if err != nil {
			fmt.Println("user not found")
			return fmt.Errorf("unauthorized")
		}

		c.Context().SetUserValue("user", u)
		return c.Next()
	}
}

func validateToken(t string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(t, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			fmt.Printf("invalid signing method: %v", token.Header["alg"])
			return nil, fmt.Errorf("unauthorized")
		}

		sec := os.Getenv("JWT_SECRET")
		return []byte(sec), nil
	})

	if err != nil {
		fmt.Println("failed to parse token: ", err)
		return nil, fmt.Errorf("unauthorized")
	}

	if !token.Valid {
		fmt.Println("invalid token")
		return nil, fmt.Errorf("unauthorized")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, fmt.Errorf("unauthorized")
	}

	return claims, nil
}
