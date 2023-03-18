package valueobj

import "errors"

type Contact string

func NewContact(contact string) (Contact, error) {
	if l := len(contact); l < 0 || l > 50 {
		return "", errors.New(" the contact name length does not meet the requirements")
	}
	return Contact(contact), nil
}
