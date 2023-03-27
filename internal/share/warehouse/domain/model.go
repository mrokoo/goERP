package domain

import (
	"github.com/mrokoo/goERP/internal/share/valueobj"
)

type Warehouse struct {
	ID      WarehouseId        `json:"id" binding:"required"`
	Name    valueobj.Name      `json:"name" binding:"required"`
	Admin   string             `json:"admin"`
	Phone   valueobj.Phone     `json:"phone"`
	Address valueobj.Address   `json:"address"`
	Note    string             `json:"note"`
	State   valueobj.StateType `json:"state"`
}

type WarehouseId = string
