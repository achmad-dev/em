package dto

import "time"

/*
--- MIT License (c) 2024 achmad
--- See LICENSE for more details
*/

type CreateEventRequest struct {
	Name string `json:"name" validate:"required"`
	// vendor name for the event
	VendorName string `json:"vendor_name" validate:"required"`
	PostalCode string `json:"postal_code" validate:"required"`
	Location   string `json:"location" validate:"required"`
	// proposed dates for the event
	ProposedDates []time.Time `json:"proposed_dates" validate:"required"`
}

type EventVendorOption struct {
	EventID string `json:"event_id,omitempty" validate:"required"`
	// remarks for the vendor
	Remarks *string `json:"remarks,omitempty" validate:"required"`
	// confirmed date for the event
	ConfirmedDate *time.Time `json:"confirmed_date,omitempty"`
	// status of the vendor option eg approved, rejected
	Status string `json:"status,omitempty" validate:"required"`
}
