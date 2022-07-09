package model

import "sophie-server/model/users"

// UserToUserGet converts User to UserGet
// via getting fields from User model.
func UserToUserGet(user users.User) users.UserGet {
	return users.UserGet{
		Firstname:  user.Firstname,
		Lastname:   user.Lastname,
		Nickname:   user.Nickname,
		Email:      user.Email,
		Workspaces: user.Workspaces,
	}
}

// SessionCreateToSession converts SessionCreate to Session
func SessionCreateToSession(session users.SessionCreate) users.Session {
	return users.Session{
		IP:          session.IP,
		AccessToken: session.AccessToken,
		UserAgent:   session.UserAgent,
	}
}
