package utils

import (
	"github.com/achmad/em/backend/internal/domain"
	"github.com/gofiber/fiber/v2"
)

func SendResponse(c *fiber.Ctx, res *domain.Response) error {
	response := fiber.Map{
		"success": res.Success,
		"message": res.Message,
		// "type":    res.Role,
		"status": res.Status,
		// "data":   res.Data,
	}
	if res.Data != nil {
		response["data"] = res.Data
	}
	return c.Status(res.Status).JSON(response)
}

func SuccessResponse(c *fiber.Ctx, data interface{}, message string) error {
	res := &domain.Response{
		Success: true,
		Message: message,
		Status:  fiber.StatusOK,
		Data:    data,
	}
	return SendResponse(c, res)
}

func SuccessResponseWithRole(c *fiber.Ctx, data interface{}, message, resRole string) error {
	res := &domain.Response{
		Success: true,
		Message: message,
		Role:    resRole,
		Status:  fiber.StatusOK,
		Data:    data,
	}
	return SendResponse(c, res)
}

func SuccessResponseWithStatus(c *fiber.Ctx, data interface{}, message string, status int) error {
	res := &domain.Response{
		Success: true,
		Message: message,
		Status:  status,
		Data:    data,
	}
	return SendResponse(c, res)
}

func ErrorResponse(c *fiber.Ctx, status int, message string) error {
	res := &domain.Response{
		Success: false,
		Message: message,
		Status:  status,
	}
	return SendResponse(c, res)
}
