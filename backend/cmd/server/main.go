package main

import v1 "github.com/achmad/em/backend/api/route/v1"

/*
--- MIT License (c) 2024 achmad
--- See LICENSE for more details
*/

func main() {
	// generatePassword := "12345678"
	// hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(generatePassword), bcrypt.DefaultCost)
	// fmt.Println(string(hashedPassword))
	envPath := "../../.env"
	v1.ServeRoute(envPath)
}
