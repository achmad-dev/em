package repository

/*
--- MIT License (c) 2024 achmad
--- See LICENSE for more details
*/

import (
	"context"

	"github.com/achmad/em/backend/internal/domain"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
)

// interface for event repo
type EventRepo interface {
	InsertEvent(ctx context.Context, event domain.Event) error
	GetEventsByUserID(ctx context.Context, userID string) ([]*domain.Event, error)
	GetEventsByVendorName(ctx context.Context, vendorName string) ([]*domain.Event, error)
	GetEventById(ctx context.Context, id string) (*domain.Event, error)
	UpdateEvent(ctx context.Context, event domain.Event) error
}

// event repo implementation
type eventRepositoryImpl struct {
	sqlx *sqlx.DB
}

// GetEventById implements EventRepo.
func (e *eventRepositoryImpl) GetEventById(ctx context.Context, id string) (*domain.Event, error) {
	event := &domain.Event{}
	query := `SELECT id, user_id, vendor_name, event_name, proposed_dates, rejected_remarks, status, confirmed_date, postal_code, location, created_at, updated_at FROM events WHERE id = $1`
	err := e.sqlx.GetContext(ctx, event, query, id)
	if err != nil {
		return nil, err
	}
	return event, nil
}

// GetEventsByVendorName implements EventRepo.
func (e *eventRepositoryImpl) GetEventsByVendorName(ctx context.Context, vendorName string) ([]*domain.Event, error) {
	events := []*domain.Event{}
	query := `SELECT id, user_id, vendor_name, event_name, proposed_dates, rejected_remarks, status, confirmed_date, postal_code, location, created_at, updated_at FROM events WHERE vendor_name = $1`
	err := e.sqlx.SelectContext(ctx, &events, query, vendorName)
	if err != nil {
		return nil, err
	}
	return events, nil
}

// UpdateEvent implements EventRepo.
func (e *eventRepositoryImpl) UpdateEvent(ctx context.Context, event domain.Event) error {
	tx, err := e.sqlx.BeginTxx(ctx, nil)
	if err != nil {
		return err
	}

	query := `UPDATE events SET rejected_remarks = $2, status = $3, confirmed_date = $4 WHERE id = $1`

	_, err = tx.ExecContext(ctx, query, event.ID, event.RejectedRemarks, event.Status, event.ConfirmedDate)
	if err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}

// GetEventsByUserID implements EventRepo.
func (e *eventRepositoryImpl) GetEventsByUserID(ctx context.Context, userID string) ([]*domain.Event, error) {
	events := []*domain.Event{}
	query := `SELECT id, user_id, vendor_name, event_name, proposed_dates, rejected_remarks, status, confirmed_date, postal_code, location, created_at, updated_at FROM events WHERE user_id = $1`
	err := e.sqlx.SelectContext(ctx, &events, query, userID)
	if err != nil {
		return nil, err
	}
	return events, nil
}

// InsertEvent implements EventRepo.
func (e *eventRepositoryImpl) InsertEvent(ctx context.Context, event domain.Event) error {
	// convert time.Time to date format
	proposedDates := make([]string, len(event.ProposedDates))
	for i, date := range event.ProposedDates {
		proposedDates[i] = date.Format("2006-01-02")
	}
	query := `INSERT INTO events (user_id, vendor_name, event_name, proposed_dates, status, postal_code, location) VALUES ($1, $2, $3, $4, $5, $6, $7)`
	_, err := e.sqlx.ExecContext(ctx, query, event.UserID, event.VendorName, event.EventName, pq.Array(proposedDates), event.Status, event.PostalCode, event.Location)
	if err != nil {
		return err
	}
	return nil
}

// NewEventRepository creates a new event repository
func NewEventRepository(sqlx *sqlx.DB) EventRepo {
	return &eventRepositoryImpl{sqlx: sqlx}
}
