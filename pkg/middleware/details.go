package middleware

import (
	"github.com/gofiber/fiber/v2"
)

func GetUserDetails(c *fiber.Ctx) (*UserDataClaims, error) {
	claims, ok := c.Locals("userClaims").(*UserDataClaims)
	if !ok || claims == nil {
		return nil, fiber.NewError(fiber.StatusUnauthorized, "user claims not found")
	}

	if claims.ID == "" {
		return nil, fiber.NewError(fiber.StatusUnauthorized, "invalid user data in token")
	}

	return claims, nil
}
