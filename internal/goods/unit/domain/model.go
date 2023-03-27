package domain

import "github.com/google/uuid"

type Unit struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
	Note string    `json:"note"`
}

func (u *Unit) ChangeName(name string) {
	u.Name = name
}

func (u *Unit) ChangeNote(note string) {
	u.Note = note
}
