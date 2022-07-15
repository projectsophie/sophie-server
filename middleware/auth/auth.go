package auth

import (
	"github.com/gin-gonic/gin"
	"net/http"
	schemes "sophie-server/model/users"
	"sophie-server/store"
	"sophie-server/util"
)

// AuthorizeJWT is a middleware that checks if users is authorized via JWT token.
func AuthorizeJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		token, _ := util.ValidateToken(util.GetToken(c))
		if token == nil || !token.Valid {
			c.AbortWithStatus(http.StatusUnauthorized)
		}
	}
}

// LoginVerify verifies user's credentials and returns if users is valid.
func LoginVerify(username string, password string) bool {
	user, success := store.GetUserByNickname(username)
	return success && util.VerifyPassword(user.Password, []byte(password))
}

// Login is a method
// that generates a JWT token for users via gin's context.
func Login(ctx *gin.Context) string {
	var model schemes.UserAuth
	err := ctx.ShouldBind(&model)
	if err != nil {
		return ""
	}
	isAuthed := LoginVerify(model.Username, model.Password)
	if isAuthed {
		return util.GenerateToken(model.Username)
	}
	return ""
}
