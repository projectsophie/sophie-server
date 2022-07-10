package types

type Checkbox struct {
	State bool
}

func ReverseState(control *Checkbox) {
	control.State = !control.State
}

func SetState(control *Checkbox, state bool) {
	control.State = state
}

func (checkbox *Checkbox) IsValidJson(json string) bool {
	return true
}
