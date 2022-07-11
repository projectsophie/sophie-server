package types

import "sophie-server/model/users"

// Mention is a struct which
// describes a mention type.
type Mention struct {
	User *users.User `json:"user"` // User is a user who was mentioned.
}

// GetID is a method
// that returns a user's id who was mentioned.
func (mention *Mention) GetID() int {
	return mention.User.ID
}

// GetNickname is a method
// that returns a user's nickname who was mentioned.
func (mention *Mention) GetNickname() string {
	return mention.User.Nickname
}

// IsValidJson is an implementation of Control interface.
func (mention *Mention) IsValidJson(json string) bool {
	return true
}
