package domain

/*
--- MIT License (c) 2024 achmad
--- See LICENSE for more details
*/

type Response struct {
	Success bool        `json:"success,omitempty"`
	Message string      `json:"message,omitempty"`
	Role    string      `json:"role,omitempty"`
	Status  int         `json:"status,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}
