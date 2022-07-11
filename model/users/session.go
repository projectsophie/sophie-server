package users

import (
	static "sophie-server/store"
)

// SessionCreate is a struct which
// describes a session instance while creation.
type SessionCreate struct {
	IP          string `json:"ip" form:"ip"`
	AccessToken string `json:"accessToken" form:"accessToken"`
	UserAgent   string `json:"userAgent" form:"userAgent"`
}

// Session is a struct that represents a session.
type Session struct {
	ID            int    // ID is id of session.
	CreationDate  string // CreationDate is a creation date of session.
	LastUsageDate string // LastUsageDate is a last usage date of session.
	IP            string // IP is an IP-Address which was used for session creation.
	UserAgent     string // UserAgent is a user's agent which was used for session creation.
	AccessToken   string // AccessToken is an access token of session.
}

// AppendSession appends a provided session
// to user's sessions array who was provided by token.
func (session Session) AppendSession() {
	if user, success := static.GetUserByToken(session.AccessToken); success {
		static.ApplySession(session, &user)
		user.UpdateUser()
	}
}

// SessionCreateToSession converts SessionCreate to Session
func (session *SessionCreate) SessionCreateToSession() Session {
	return Session{
		IP:          session.IP,
		AccessToken: session.AccessToken,
		UserAgent:   session.UserAgent,
	}
}
