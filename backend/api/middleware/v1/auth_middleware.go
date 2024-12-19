package v1

import (
	"strings"

	"github.com/achmad/em/backend/internal/service"
	"github.com/achmad/em/backend/pkg/utils"
	"github.com/gofiber/fiber/v2"
)

/*
--- MIT License (c) 2024 achmad
--- See LICENSE for more details
*/

func AuthMiddleware(secret string, userService service.UserService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return utils.ErrorResponse(c, fiber.StatusUnauthorized, "Unauthorized")
		}
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenString == "" {
			return utils.ErrorResponse(c, fiber.StatusUnauthorized, "Unauthorized")
		}
		claims, err := utils.ValidateToken(tokenString, secret)
		if err != nil {
			return utils.ErrorResponse(c, fiber.StatusUnauthorized, "Unauthorized")
		}

		user, err := userService.GetUserByID(c.Context(), claims.UserId)
		if err != nil {
			return utils.ErrorResponse(c, fiber.StatusUnauthorized, "Unauthorized")
		}
		c.Locals("user_id", string(user.ID))
		c.Locals("user_role", claims.Role)
		return c.Next()
	}
}
