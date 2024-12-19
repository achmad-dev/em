package service

/*
--- MIT License (c) 2024 achmad
--- See LICENSE for more details
*/

import (
	"context"
	"errors"

	"github.com/achmad/em/backend/internal/domain"
	"github.com/achmad/em/backend/internal/repository"
	"github.com/sirupsen/logrus"
)

// event service interface
type EventService interface {
	InsertEvent(ctx context.Context, event domain.Event) error
	GetEventsByVendorName(ctx context.Context, vendorName string) ([]*domain.Event, error)
	GetEventById(ctx context.Context, id string) (*domain.Event, error)
	GetEventsByUserID(ctx context.Context, userID string) ([]*domain.Event, error)
	UpdateEvent(ctx context.Context, event domain.Event, userId string) error
}

// event service impl
type eventService struct {
	eventRepo repository.EventRepo
	userRepo  repository.UserRepo
	log       *logrus.Logger
}

// GetEventById implements EventService.
func (e *eventService) GetEventById(ctx context.Context, id string) (*domain.Event, error) {
	event, err := e.eventRepo.GetEventById(ctx, id)
	if err != nil {
		e.log.Errorf("Failed to get event with id %s: %v", id, err)
		return nil, errors.New("something went wrong")
	}
	return event, nil
}

// GetEventsByVendorName implements EventService.
func (e *eventService) GetEventsByVendorName(ctx context.Context, vendorName string) ([]*domain.Event, error) {
	events, err := e.eventRepo.GetEventsByVendorName(ctx, vendorName)
	if err != nil {
		e.log.Errorf("Failed to get events for vendor %s: %v", vendorName, err)
		return nil, errors.New("something went wrong")
	}
	return events, nil
}

// GetEventsByUserID implements EventService.
func (e *eventService) GetEventsByUserID(ctx context.Context, userID string) ([]*domain.Event, error) {
	events, err := e.eventRepo.GetEventsByUserID(ctx, userID)
	if err != nil {
		e.log.Errorf("Failed to get events for user %s: %v", userID, err)
		return nil, errors.New("something went wrong")
	}
	return events, nil
}

// InsertEvent implements EventService.
func (e *eventService) InsertEvent(ctx context.Context, event domain.Event) error {
	if err := e.eventRepo.InsertEvent(ctx, event); err != nil {
		e.log.Errorf("Failed to insert event: %v", err)
		return errors.New("something went wrong")
	}
	return nil
}

// UpdateEvent implements EventService.
func (e *eventService) UpdateEvent(ctx context.Context, event domain.Event, userId string) error {
	user, err := e.userRepo.GetUserByID(ctx, userId)
	if err != nil {
		e.log.Errorf("Failed to get user with id %s: %v", userId, err)
		return errors.New("something went wrong")
	}
	eventR, err := e.eventRepo.GetEventById(ctx, event.ID)
	if err != nil {
		e.log.Errorf("Failed to get event with id %s: %v", event.ID, err)
		return errors.New("something went wrong")
	}

	if user.Role != "vendor" || eventR.VendorName != user.CompanyName {
		return errors.New("unauthorized")
	}

	if err := e.eventRepo.UpdateEvent(ctx, event); err != nil {
		e.log.Errorf("Failed to update event: %v", err)
		return errors.New("something went wrong")
	}
	return nil
}

// new event service
func NewEventService(eventRepo repository.EventRepo, userRepo repository.UserRepo, log *logrus.Logger) EventService {
	return &eventService{
		eventRepo: eventRepo,
		userRepo:  userRepo,
		log:       log,
	}
}
