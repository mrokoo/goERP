package state

type State string

const (
	STATE_ACTIVE State = "active"
	STATE_FREEZE State = "freeze"
)

// func NewState(state int) (State, error) {
// 	if state < 1 || state > 2 {
// 		return STATE_INVAILD, errors.New("the state is invaild")
// 	}
// 	return State(state), nil
// }

func (s *State) String() string {
	switch *s {
	case STATE_ACTIVE:
		return "active"
	default:
		return "freeze"
	}
}
