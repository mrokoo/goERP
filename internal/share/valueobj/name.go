package valueobj

import "errors"

type Name string

func NewName(name string) (Name, error) {
	if l := len(name); l < 0 || l > 50 {
		return "", errors.New(" the name length does not meet the requirements")
	}
	return Name(name), nil
}
