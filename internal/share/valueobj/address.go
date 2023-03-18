package valueobj

import "errors"

type Address string

func NewAddress(addr string) (Address, error) {
	if len(addr) > 50 {
		return "", errors.New("the address length is too length")
	}
	return Address(addr), nil
}
