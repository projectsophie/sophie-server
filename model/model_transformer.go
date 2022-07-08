package model

// UserToUserGet converts User to UserGet
// via getting fields from User model.
func UserToUserGet(user User) UserGet {
	return UserGet{
		Firstname:  user.Firstname,
		Lastname:   user.Lastname,
		Nickname:   user.Nickname,
		Email:      user.Email,
		Workspaces: user.Workspaces,
	}
}

// SessionCreateToSession converts SessionCreate to Session
func SessionCreateToSession(session SessionCreate) Session {
	return Session{
		IP:          session.IP,
		AccessToken: session.AccessToken,
		UserAgent:   session.UserAgent,
	}
}
