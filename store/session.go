package store

import (
	middleware "sophie-server/middleware/session"
	"sophie-server/model"
)

// ApplySession is a method
// that applies a session to a user via session array in user struct.
func ApplySession(session model.Session, user *model.User) {
	sessionsArray := middleware.GetSessionsFromJson(user.Sessions)
	sessionsArray[len(sessionsArray)] = session
	user.Sessions = middleware.ConvertSessionsToJson(sessionsArray)
}

// RemoveSession is a method
// that removes a session from a user via session array in user struct.
func RemoveSession(token string, user *model.User) {
	sessionsArray := middleware.GetSessionsFromJson(user.Sessions)
	for i := 0; i < len(sessionsArray); i++ {
		if sessionsArray[i].AccessToken == token {
			sessionsArray = append(sessionsArray[:i], sessionsArray[i+1:]...)
			break
		}
	}
	user.Sessions = middleware.ConvertSessionsToJson(sessionsArray)
}
