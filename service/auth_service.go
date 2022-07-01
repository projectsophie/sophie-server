package service

import (
	"github.com/gin-gonic/gin"
	schemes "sophie-server/model"
)

func Login(ctx *gin.Context) string {
	var model schemes.UserAuth
	err := ctx.ShouldBind(&model)
	if err != nil {
		return ""
	}
	isAuthed := LoginVerify(model.Username, model.Password)
	if isAuthed {
		return GenerateToken(model.Username)
	}
	return ""
}
