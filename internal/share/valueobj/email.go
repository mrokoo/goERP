package valueobj

import (
	"errors"
	"regexp"
)

type Email string

func NewEmail(email string) (Email, error) {
	regRuler := `^[0-9a-z][_.0-9a-z-]{0,31}@([0-9a-z][0-9a-z-]{0,30}[0-9a-z]\.){1,4}[a-z]{2,4}$`
	reg := regexp.MustCompile(regRuler)
	if !reg.MatchString(email) {
		return "", errors.New("not a email")
	}
	return Email(email), nil
}
