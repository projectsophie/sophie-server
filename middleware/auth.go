package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"sophie-server/service"
)

const BEARER_SCHEMA = "Bearer " // BEARER_SCHEMA is a default schema for auth via bearer tokens.

func AuthorizeJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		tokenRaw := authHeader[len(BEARER_SCHEMA):]
		token, _ := service.ValidateToken(tokenRaw)
		if !token.Valid {
			c.AbortWithStatus(http.StatusUnauthorized)
		}
	}
}
