package domain

import (
	"github.com/google/uuid"
)

type User struct {
	ID       uuid.UUID
	Name     string
	Phone    string
	Email    string
	Gender   string
	Password string
}
