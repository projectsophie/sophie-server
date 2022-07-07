package service

import (
	"github.com/gin-gonic/gin"
	schemes "sophie-server/model"
	"sophie-server/util"
)

// Login is a method
// that generates a JWT token for user via gin's context.
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
