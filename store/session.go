package store

import (
	middleware "sophie-server/middleware/session"
	models "sophie-server/model/users"
)

// ApplySession is a method
// that applies a session to a users via session array in users struct.
func ApplySession(session models.Session, user *models.User) {
	sessionsArray := middleware.GetSessionsFromJson(user.Sessions)
	session.ID = len(sessionsArray)
	sessionsArray = append(sessionsArray, session)
	user.Sessions = middleware.ConvertSessionsToJson(sessionsArray)
}

// RemoveSession is a method
// that removes a session from a users via session array in users struct.
func RemoveSession(token string, user *models.User) {
	sessionsArray := middleware.GetSessionsFromJson(user.Sessions)
	for i := 0; i < len(sessionsArray); i++ {
		if sessionsArray[i].AccessToken == token {
			sessionsArray = append(sessionsArray[:i], sessionsArray[i+1:]...)
			break
		}
	}
	user.Sessions = middleware.ConvertSessionsToJson(sessionsArray)
}

// DeleteSession removes a provided session
// from user's sessions array who was provided by token.
func DeleteSession(token string) {
	if user, success := GetUserByToken(token); success {
		RemoveSession(token, &user)
		user.UpdateUser()
	}
}

// AppendSession appends a provided session
// to user's sessions array who was provided by token.
func AppendSession(session models.Session) {
	if user, success := GetUserByToken(session.AccessToken); success {
		ApplySession(session, &user)
		user.UpdateUser()
	}
}

// SessionCreateToSession converts SessionCreate to Session
func SessionCreateToSession(session models.SessionCreate) models.Session {
	return models.Session{
		IP:          session.IP,
		AccessToken: session.AccessToken,
		UserAgent:   session.UserAgent,
	}
}
