package valueobj

import (
	"errors"
	"regexp"
)

type PhoneNumber string

func NewPhoneNumber(number string) (PhoneNumber, error) {
	regRuler := "^((13[0-9])|(14[5,7])|(15[0-3,5-9])|(17[0,3,5-8])|(18[0-9])|166|198|199|(147))\\d{8}$"
	reg := regexp.MustCompile(regRuler)
	if !reg.MatchString(number) {
		return "", errors.New("not a phone number")
	}
	return PhoneNumber(number), nil
}
