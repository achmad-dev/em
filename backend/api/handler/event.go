package handler

import (
	"github.com/achmad/em/backend/api/dto"
	"github.com/achmad/em/backend/internal/domain"
	"github.com/achmad/em/backend/internal/service"
	"github.com/achmad/em/backend/pkg/utils"
	"github.com/gofiber/fiber/v2"
)

/*
--- MIT License (c) 2024 achmad
--- See LICENSE for more details
*/

type EventHandler interface {
	InsertEvent(c *fiber.Ctx) error
	GetEvents(c *fiber.Ctx) error
	UpdateEvent(c *fiber.Ctx) error
	GetCompanies(c *fiber.Ctx) error
}

type eventHandlerImpl struct {
	eventService service.EventService
	userService  service.UserService
	reqService   service.RequestLogService
}

// GetCompanies implements EventHandler.
func (e *eventHandlerImpl) GetCompanies(c *fiber.Ctx) error {
	companies, err := e.userService.GetUsersCompany(c.Context(), "vendor")
	if err != nil {
		return utils.ErrorResponse(c, fiber.StatusInternalServerError, "Internal Server Error")
	}
	return utils.SuccessResponse(c, companies, "Success")
}

// GetEvents implements EventHandler.
func (e *eventHandlerImpl) GetEvents(c *fiber.Ctx) error {
	if c.Locals("user_role") == "hr" {
		events, err := e.eventService.GetEventsByUserID(c.Context(), c.Locals("user_id").(string))
		if err != nil {
			return utils.ErrorResponse(c, fiber.StatusInternalServerError, "Internal Server Error")
		}
		return utils.SuccessResponseWithRole(c, events, "Success", "hr")
	}
	if c.Locals("user_role") == "vendor" {
		user, err := e.userService.GetUserByID(c.Context(), c.Locals("user_id").(string))
		if err != nil {
			return utils.ErrorResponse(c, fiber.StatusUnauthorized, "Unauthorized")
		}
		events, err := e.eventService.GetEventsByVendorName(c.Context(), user.CompanyName)
		if err != nil {
			return utils.ErrorResponse(c, fiber.StatusInternalServerError, "Internal Server Error")
		}
		return utils.SuccessResponseWithRole(c, events, "Success", "vendor")
	}
	return utils.ErrorResponse(c, fiber.StatusUnauthorized, "Unauthorized")
}

// UpdateEvent implements EventHandler.
func (e *eventHandlerImpl) UpdateEvent(c *fiber.Ctx) error {
	if c.Locals("user_role") != "vendor" {
		return utils.ErrorResponse(c, fiber.StatusUnauthorized, "Unauthorized")
	}

	eventId := c.Query("eventId")

	var updateEventRequest dto.EventVendorOption
	if err := c.BodyParser(&updateEventRequest); err != nil {
		return utils.ErrorResponse(c, fiber.StatusBadRequest, "Bad Request")
	}
	event := domain.Event{
		ID:              eventId,
		RejectedRemarks: updateEventRequest.Remarks,
		Status:          updateEventRequest.Status,
		ConfirmedDate:   updateEventRequest.ConfirmedDate,
	}
	if event.ConfirmedDate == nil || event.Status == "rejected" {
		event.ConfirmedDate = nil
		event.Status = "rejected"
	}
	if event.Status == "confirmed" && event.ConfirmedDate != nil {
		event.RejectedRemarks = nil
	}
	if err := e.eventService.UpdateEvent(c.Context(), event, c.Locals("user_id").(string)); err != nil {
		e.reqService.InsertRequestLog(c.Context(), c.Locals("user_id").(string), "update event", err.Error())
		return utils.ErrorResponse(c, fiber.StatusInternalServerError, "Internal Server Error")
	}
	e.reqService.InsertRequestLog(c.Context(), c.Locals("user_id").(string), "update event", "success")
	return utils.SuccessResponse(c, nil, "Success")
}

// InsertEvent implements EventHandler.
func (e *eventHandlerImpl) InsertEvent(c *fiber.Ctx) error {
	var createEventRequest dto.CreateEventRequest
	if err := c.BodyParser(&createEventRequest); err != nil {
		return utils.ErrorResponse(c, fiber.StatusBadRequest, "Bad Request")
	}
	// only hr can create event
	userId := c.Locals("user_id").(string)
	if c.Locals("user_role") != "hr" {
		return utils.ErrorResponse(c, fiber.StatusUnauthorized, "Unauthorized")
	}
	event := domain.Event{
		UserID:        userId,
		EventName:     createEventRequest.Name,
		VendorName:    createEventRequest.VendorName,
		PostalCode:    createEventRequest.PostalCode,
		Location:      createEventRequest.Location,
		ProposedDates: createEventRequest.ProposedDates,
		Status:        "pending",
	}
	if err := e.eventService.InsertEvent(c.Context(), event); err != nil {
		e.reqService.InsertRequestLog(c.Context(), c.Locals("user_id").(string), "create event", err.Error())
		return utils.ErrorResponse(c, fiber.StatusInternalServerError, "Internal Server Error")
	}
	e.reqService.InsertRequestLog(c.Context(), c.Locals("user_id").(string), "create event", "success")
	return utils.SuccessResponse(c, nil, "Success")
}

func NewEventHandler(eventService service.EventService, userService service.UserService, reqService service.RequestLogService) EventHandler {
	return &eventHandlerImpl{
		eventService: eventService,
		userService:  userService,
		reqService:   reqService,
	}
}
