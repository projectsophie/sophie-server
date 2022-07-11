package controls

// Control is a basic interface for any controls.
type Control interface {
	IsValidJson(string) bool // IsValidJson checks if requested JSON string for control is valid.
}
