package domain

/*
--- MIT License (c) 2024 achmad
--- See LICENSE for more details
*/

type RequestLog struct {
	UserID  string `db:"user_id" json:"user_id,omitempty"`
	Message string `db:"message" json:"message,omitempty"`
	Status  string `db:"status" json:"status,omitempty"`
}
