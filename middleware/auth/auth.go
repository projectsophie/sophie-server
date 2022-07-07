package auth

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"sophie-server/util"
)

const BEARER_SCHEMA = "Bearer " // BEARER_SCHEMA is a default schema for auth via bearer tokens.

// AuthorizeJWT is a middleware that checks if user is authorized via JWT token.
func AuthorizeJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		tokenRaw := authHeader[len(BEARER_SCHEMA):]
		token, _ := util.ValidateToken(tokenRaw)
		if !token.Valid {
			c.AbortWithStatus(http.StatusUnauthorized)
		}
	}
}
