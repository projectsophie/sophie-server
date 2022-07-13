package util

import "time"

// GetCurrentServerTime returns current date in format "2006-01-02 15:04:05".
func GetCurrentServerTime() string {
	return time.Now().Format("2006-01-02 15:04:05")
}
