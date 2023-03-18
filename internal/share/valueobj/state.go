package valueobj

import "errors"

type StateType int

const (
	STATE_INVAILD StateType = iota
	STATE_ACTIVE
	STATE_FREEZE
)

func NewState(state int) (StateType, error) {
	if state < 1 || state > 2 {
		return STATE_INVAILD, errors.New("the state is invaild")
	}
	return StateType(state), nil
}
