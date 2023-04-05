package domain

import "github.com/google/uuid"

type Unit struct {
	ID   uuid.UUID `gorm:"primaryKey;<-:create"`
	Name string    `gorm:"not null"`
	Note string
}
