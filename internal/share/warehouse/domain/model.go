package domain

import (
	"github.com/mrokoo/goERP/internal/share/valueobj/state"
)

type Warehouse struct {
	ID      string      `json:"id" binding:"required" gorm:"primaryKey;<-:create"`
	Name    string      `json:"name" binding:"required" gorm:"not null"`
	Admin   string      `json:"admin"`
	Phone   string      `json:"phone"`
	Address string      `json:"address"`
	Note    string      `json:"note"`
	State   state.State `json:"state" gorm:"default:active"`
}
