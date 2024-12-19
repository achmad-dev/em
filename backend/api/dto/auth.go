package dto

/*
--- MIT License (c) 2024 achmad
--- See LICENSE for more details
*/

type AuthDto struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type AuthResponseDto struct {
	Token string `json:"token"`
	Role  string `json:"role"`
}
