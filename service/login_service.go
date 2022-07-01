package service

import (
	"sophie-server/store"
	"sophie-server/util"
)

func LoginVerify(username string, password string) bool {
	user, success := store.GetUserByNickname(username)
	return success && util.VerifyPassword(user.Password, []byte(password))
}
