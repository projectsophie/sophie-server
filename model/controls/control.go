package controls

type Control interface {
	IsValidJson(string) bool
}
