package bot

import "sync"

type fsm struct {
	current sync.Map
}

type State string

const (
	StateDefault State = "default"
	StateAny     State = "any"
)

// Доставать из апдейта
func (f *fsm) currentState(userID int64) State {
	v, ok := f.current.Load(userID)
	if !ok {
		return StateDefault
	}

	state := v.(State)

	return state
}

func (b *Bot) StateTransition(state State, userID int64) {

	b.fsm.current.Store(userID, state)
}

func (b *Bot) StateFinish(userID int64) {

	b.fsm.current.Delete(userID)
}
