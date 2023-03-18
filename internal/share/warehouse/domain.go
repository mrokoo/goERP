package warehouse

import (
	"errors"
	"regexp"

	"github.com/mrokoo/goERP/internal/share/valueobj"
)

type Warehouse struct {
	ID          WarehouseId          `json:"id"`
	Name        valueobj.Name        `json:"name"`
	Admin       string               `json:"admin"`
	PhoneNumber valueobj.PhoneNumber `json:"phoneNumber"`
	Address     valueobj.Address     `json:"address"`
	Note        string               `json:"note"`
	State       valueobj.StateType   `json:"state"`
}

type WarehouseCMD struct {
	ID          string
	Name        string
	Admin       string
	PhoneNumber string
	Address     string
	Note        string
	State       int
}

func NewWarehouse(cmd WarehouseCMD) (Warehouse, error) {
	var w Warehouse
	var err error
	w.ID, err = NewWarehouseID(cmd.ID)
	if err != nil {
		return Warehouse{}, err
	}
	w.Name, err = valueobj.NewName(cmd.Name)
	if err != nil {
		return Warehouse{}, err
	}
	w.Admin = cmd.Admin

	w.PhoneNumber, err = valueobj.NewPhoneNumber(cmd.PhoneNumber)
	if err != nil {
		return Warehouse{}, err
	}

	w.Address, err = valueobj.NewAddress(cmd.Address)
	if err != nil {
		return Warehouse{}, err
	}

	w.Note = cmd.Note

	w.State, err = valueobj.NewState(cmd.State)
	if err != nil {
		return Warehouse{}, err
	}
	return w, nil
}

type WarehouseId string

func (s *WarehouseId) CheckWarehouseID() error {
	regRuler := "^W[0-9]+$"
	reg := regexp.MustCompile(regRuler)
	if !reg.MatchString(string(*s)) {
		return errors.New("the WarehouseID is wrong")
	}
	return nil
}

func NewWarehouseID(id string) (WarehouseId, error) {
	wid := WarehouseId(id)
	if err := wid.CheckWarehouseID(); err != nil {
		return "", err
	}
	return wid, nil
}
