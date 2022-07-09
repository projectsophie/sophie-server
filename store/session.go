package store

import (
	middleware "sophie-server/middleware/session"
	"sophie-server/model/users"
)

// ApplySession is a method
// that applies a session to a users via session array in users struct.
func ApplySession(session users.Session, user *users.User) {
	sessionsArray := middleware.GetSessionsFromJson(user.Sessions)
	session.ID = len(sessionsArray)
	sessionsArray = append(sessionsArray, session)
	user.Sessions = middleware.ConvertSessionsToJson(sessionsArray)
}

// RemoveSession is a method
// that removes a session from a users via session array in users struct.
func RemoveSession(token string, user *users.User) {
	sessionsArray := middleware.GetSessionsFromJson(user.Sessions)
	for i := 0; i < len(sessionsArray); i++ {
		if sessionsArray[i].AccessToken == token {
			sessionsArray = append(sessionsArray[:i], sessionsArray[i+1:]...)
			break
		}
	}
	user.Sessions = middleware.ConvertSessionsToJson(sessionsArray)
}
