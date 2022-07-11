package util

import "encoding/json"

// ToJson converts a struct to json string.
func ToJson(meta interface{}) string {
	if metadata, err := json.Marshal(meta); err == nil {
		return string(metadata)
	}
	return ""
}
