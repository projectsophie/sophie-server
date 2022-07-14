package types

// SophieLink is a struct which
// describes a Sophie link type.
type SophieLink struct {
	PageID int    `json:"pageId"`
	Text   string `json:"text"`
}

// IsValidJson is an implementation of Control interface.
func (sophieLink *SophieLink) IsValidJson(json string) bool {
	return true
}
