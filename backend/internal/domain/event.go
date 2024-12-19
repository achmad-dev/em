package domain

/*
--- MIT License (c) 2024 achmad
--- See LICENSE for more details
*/

import (
	"database/sql/driver"
	"fmt"
	"strings"
	"time"

	"github.com/lib/pq"
)

type Event struct {
	ID              string     `db:"id" json:"id,omitempty"`
	UserID          string     `db:"user_id" json:"user_id,omitempty"`
	VendorName      string     `db:"vendor_name" json:"vendor_name,omitempty"`
	EventName       string     `db:"event_name" json:"event_name,omitempty"`
	ProposedDates   TimeSlice  `db:"proposed_dates" json:"proposed_dates,omitempty"`
	RejectedRemarks *string    `db:"rejected_remarks" json:"rejected_remarks,omitempty"`
	Status          string     `db:"status" json:"status,omitempty"`
	ConfirmedDate   *time.Time `db:"confirmed_date" json:"confirmed_date,omitempty"`
	PostalCode      string     `db:"postal_code" json:"postal_code,omitempty"`
	Location        string     `db:"location" json:"location,omitempty"`
	CreatedAt       time.Time  `db:"created_at" json:"created_at,omitempty"`
	UpdatedAt       time.Time  `db:"updated_at" json:"updated_at,omitempty"`
}

type TimeSlice []time.Time

func (ts *TimeSlice) Scan(value interface{}) error {
	if value == nil {
		*ts = nil
		return nil
	}

	switch v := value.(type) {
	case string:
		// PostgreSQL may return timestamp[] as a string
		times, err := parseTimestampArray(v)
		if err != nil {
			return err
		}
		*ts = times
	case []byte:
		// Or as a byte slice
		s := string(v)
		times, err := parseTimestampArray(s)
		if err != nil {
			return err
		}
		*ts = times
	default:
		return fmt.Errorf("unsupported type for TimeSlice: %T", v)
	}

	return nil
}

func (ts TimeSlice) Value() (driver.Value, error) {
	return pq.Array([]time.Time(ts)).Value()
}

func parseTimestampArray(s string) ([]time.Time, error) {
	s = strings.Trim(s, "{}")
	if s == "" {
		return nil, nil
	}
	elements := strings.Split(s, ",")
	times := make([]time.Time, len(elements))
	for i, elem := range elements {
		elem = strings.Trim(elem, "\"") // Remove any surrounding quotes
		t, err := time.Parse("2006-01-02 15:04:05", elem)
		if err != nil {
			return nil, fmt.Errorf("failed to parse time '%s': %v", elem, err)
		}
		times[i] = t
	}
	return times, nil
}
