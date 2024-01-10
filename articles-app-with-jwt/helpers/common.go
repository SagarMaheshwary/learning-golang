package helpers

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func CreateSlug(str string) string {
	return strings.ReplaceAll(strings.ToLower(str), " ", "-")
}

func AuthUserId(c *fiber.Ctx) uint {
	token := c.Locals("user").(*jwt.Token)
	claims := token.Claims.(jwt.MapClaims)

	return uint(claims["user_id"].(float64))
}
