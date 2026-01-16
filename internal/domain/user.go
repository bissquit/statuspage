package domain

import "time"

type Role string

const (
	RoleUser     Role = "user"
	RoleOperator Role = "operator"
	RoleAdmin    Role = "admin"
)

type User struct {
	ID           string
	Email        string
	PasswordHash string
	FirstName    string
	LastName     string
	Role         Role
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
