package domain

import (
	"github.com/google/uuid"
)

type Category struct {
	ID   uuid.UUID `json:"id" bson:"id"`
	Name string    `json:"name" bson:"name"`
	Note string    `json:"note" bson:"note"`
}

func (c *Category) ChangeName(name string) {
	// 可以封装一些检验逻辑，如：name长度
	c.Name = name
}

func (c *Category) ChangeNote(note string) {
	// 可以封装一些检验逻辑，如：note长度
	c.Note = note
}
