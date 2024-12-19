package v1

/*
--- MIT License (c) 2024 achmad
--- See LICENSE for more details
*/

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"

	"github.com/gofiber/fiber/v2"
)

func HMACMiddleware(secretKey string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		xSignature := c.Get("X-Signature")
		xTimestamp := c.Get("X-Timestamp")
		if xSignature == "" {
			return c.Status(fiber.StatusUnauthorized).SendString("Missing XSignature header")
		}
		if xTimestamp == "" {
			return c.Status(fiber.StatusUnauthorized).SendString("Missing Xtimestamp header")
		}

		mac := hmac.New(sha256.New, []byte(secretKey))
		mac.Write([]byte(xTimestamp))
		expectedMAC := mac.Sum(nil)
		expectedSignature := hex.EncodeToString(expectedMAC)

		if !hmac.Equal([]byte(xSignature), []byte(expectedSignature)) {
			return c.Status(fiber.StatusUnauthorized).SendString("Invalid XSignature")
		}

		return c.Next()
	}
}
