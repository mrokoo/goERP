package warehouse

import "github.com/mrokoo/goERP/internal/share/valueobj"

type WareHouse struct {
	ID          WareHouseId          `json:"id"`
	Name        valueobj.Name        `json:"name"`
	Admin       string               `json:"admin"`
	PhoneNumber valueobj.PhoneNumber `json:"phoneNumber"`
}

type WareHouseId string
