package models

import "github.com/google/uuid"

type User struct {
	ID        uuid.UUID `json:"id" db:"id"`
	Email     string    `json:"email" db:"email"`
	Password  []byte    `json:"password" db:"password"`
	Name      string    `json:"name" db:"name"`
	Surname   string    `json:"surname" db:"surname"`
	ShortName string    `json:"short_name" db:"short_name"`
	IsActive  bool      `json:"is_active" db:"is_active"`
}
