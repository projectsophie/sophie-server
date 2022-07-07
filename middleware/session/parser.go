package session

import (
	parser "encoding/json"
	"sophie-server/model"
)

// GetSessionsFromJson is a method
// that parses a JSON string and returns a slice of sessions ([]model.Session).
func GetSessionsFromJson(json string) []model.Session {
	var sessions []model.Session
	err := parser.Unmarshal([]byte(json), &sessions)
	if err != nil {
		return nil
	}
	return sessions
}

// ConvertSessionsToJson is a method
// that converts a slice of sessions ([]model.Session) to a JSON string.
func ConvertSessionsToJson(sessions []model.Session) string {
	json, err := parser.Marshal(sessions)
	if err != nil {
		return ""
	}
	return string(json)
}
