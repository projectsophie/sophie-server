package service

import (
	"sophie-server/store"
	"sophie-server/util"
)

// LoginVerify verifies user's credentials and returns if user is valid.
func LoginVerify(username string, password string) bool {
	user, success := store.GetUserByNickname(username)
	return success && util.VerifyPassword(user.Password, []byte(password))
}
