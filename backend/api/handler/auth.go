package handler

import (
	"github.com/achmad/em/backend/api/dto"
	"github.com/achmad/em/backend/internal/service"
	"github.com/achmad/em/backend/pkg/utils"
	"github.com/gofiber/fiber/v2"
)

/*
--- MIT License (c) 2024 achmad
--- See LICENSE for more details
*/

type AuthHandler interface {
	SignIn(c *fiber.Ctx) error
}

type authHandlerImpl struct {
	userService service.UserService
}

// SignIn implements AuthHandler.
func (a *authHandlerImpl) SignIn(c *fiber.Ctx) error {
	var signInRequest dto.AuthDto
	if err := c.BodyParser(&signInRequest); err != nil {
		return utils.ErrorResponse(c, fiber.StatusBadRequest, "Bad Request")
	}
	authResponse, err := a.userService.SignIn(c.Context(), signInRequest.Username, signInRequest.Password)
	if err != nil {
		return utils.ErrorResponse(c, fiber.StatusUnauthorized, "Unauthorized")
	}

	return utils.SuccessResponse(c, authResponse, "Success")
}

func NewAuthHandler(userService service.UserService) AuthHandler {
	return &authHandlerImpl{
		userService: userService,
	}
}
