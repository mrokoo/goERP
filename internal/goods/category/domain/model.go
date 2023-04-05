package domain

import (
	"github.com/google/uuid"
)

type Category struct {
	ID   uuid.UUID `gorm:"primaryKey;<-:create"`
	Name string    `gorm:"not null"`
	Note string   
}

func (c *Category) ChangeName(name string) {
	// 可以封装一些检验逻辑，如：name长度
	c.Name = name
}

func (c *Category) ChangeNote(note string) {
	// 可以封装一些检验逻辑，如：note长度
	c.Note = note
}
