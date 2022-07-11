package types

// Checkbox is a struct which
// describes a checkbox type for database.
type Checkbox struct {
	State bool // State is a state of checkbox.
}

// ReverseState reverses the state of checkbox.
func (control *Checkbox) ReverseState() {
	control.State = !control.State
}

// SetState sets the state of checkbox (state variable)
func (control *Checkbox) SetState(state bool) {
	control.State = state
}

// IsValidJson is an implementation of Control interface.
func (control *Checkbox) IsValidJson(json string) bool {
	return true
}
