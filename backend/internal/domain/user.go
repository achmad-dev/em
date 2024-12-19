package domain

/*
--- MIT License (c) 2024 achmad
--- See LICENSE for more details
*/

import (
	"time"
)

type User struct {
	ID          string    `db:"id" json:"id,omitempty"`
	Username    string    `db:"username" json:"username,omitempty"`
	Password    string    `db:"password" json:"password,omitempty"`
	Role        string    `db:"role" json:"role,omitempty"`
	CompanyName string    `db:"company_name" json:"company_name,omitempty"`
	CreatedAt   time.Time `db:"created_at" json:"created_at,omitempty"`
	UpdatedAt   time.Time `db:"updated_at" json:"updated_at,omitempty"`
}

type UserCompanyResponse struct {
	CompanyName string `db:"company_name" json:"company_name"`
}
