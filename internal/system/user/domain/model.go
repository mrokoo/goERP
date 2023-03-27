package domain

import (
	"github.com/google/uuid"
)

type User struct {
	ID       uuid.UUID
	UserName string
	Name     string
	Phone    string
	Email    string
	Gender   string
	State    int
	Role     string
}
