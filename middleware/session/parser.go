package session

import (
	parser "encoding/json"
	"github.com/gin-gonic/gin"
	"sophie-server/model/users"
	"sophie-server/util"
	"time"
)

// GetSessionsFromJson is a method
// that parses a JSON string and returns a slice of sessions ([]model.Session).
func GetSessionsFromJson(json string) []users.Session {
	var sessions []users.Session
	err := parser.Unmarshal([]byte(json), &sessions)
	if err != nil {
		return nil
	}
	return sessions
}

// ConvertSessionsToJson is a method
// that converts a slice of sessions ([]model.Session) to a JSON string.
func ConvertSessionsToJson(sessions []users.Session) string {
	json, err := parser.Marshal(sessions)
	if err != nil {
		return ""
	}
	return string(json)
}

// ValidateSession is a method
// that validates a session and returns if it is valid.
func ValidateSession(session users.Session) bool {
	return session.AccessToken != "" && util.IsCorrectIP(session.IP) && session.UserAgent != ""
}

// GenerateSession is a method
// that generates a session for a users.
func GenerateSession(c *gin.Context, token string) users.Session {
	return users.Session{
		IP:            c.ClientIP(),
		AccessToken:   token,
		CreationDate:  time.Now().Format("2006-01-02 15:04:05"),
		LastUsageDate: time.Now().Format("2006-01-02 15:04:05"),
		UserAgent:     c.GetHeader("User-Agent"),
	}
}
