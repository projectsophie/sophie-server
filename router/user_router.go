package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"sophie-server/middleware/auth"
	"sophie-server/middleware/session"
	"sophie-server/model/pages"
	"sophie-server/model/users"
	"sophie-server/store"
	"sophie-server/util"
)

// CreateUser Creates users instance and adds it to database
func CreateUser(c *gin.Context) {
	var userCreate users.UserCreate
	if err := c.BindJSON(&userCreate); err != nil {
		return
	}
	c.IndentedJSON(http.StatusOK, store.CreateUser(userCreate))
}

// AuthUser is a method that generates a JWT token for users via gin's context
// or rejects the request if users data is not valid and not found in database.
func AuthUser(c *gin.Context) {
	token := auth.Login(c)
	if token != "" {
		store.AppendSession(session.GenerateSession(c, token))
		c.JSON(http.StatusOK, gin.H{"token": token})
	} else {
		c.JSON(http.StatusUnauthorized, nil)
	}
}

// GetCurrentUser is a method that returns
// current users instance via gin's context.
func GetCurrentUser(c *gin.Context) {
	userGet, success := store.GetAuthedUser(util.GetToken(c))
	if success {
		c.IndentedJSON(http.StatusOK, userGet)
	} else {
		c.JSON(http.StatusUnauthorized, nil)
	}
}

// Logout is a method that logs out users via gin's context.
func Logout(c *gin.Context) {
	store.DeleteSession(util.GetToken(c))
	c.JSON(http.StatusOK, nil)
}

// GetUserWorkspaces is a method that returns
// current users workspaces via gin's context.
func GetUserWorkspaces(c *gin.Context) {
	userGet, success := store.GetAuthedUser(util.GetToken(c))
	if success {
		c.IndentedJSON(http.StatusOK, userGet.Workspaces)
	} else {
		c.JSON(http.StatusUnauthorized, nil)
	}
}

// UpdateUser updates user's data in database
// with given data in users.UserUpdate model.
func UpdateUser(c *gin.Context) {
	var user users.UserUpdate
	if err := c.BindJSON(&user); err != nil {
		return
	}
	store.UpdateUser(&user)
}

// CreateComment creates comment via gin.Context.
// It binds pages.CommentCreate model from request
// and adds comment instance to database.
func CreateComment(c *gin.Context) {
	var comment pages.CommentCreate
	if err := c.BindJSON(&comment); err != nil {
		return
	}
	user, success := store.GetUserByToken(util.GetToken(c))
	if success {
		store.CreateComment(&comment, user.ID)
		c.Status(http.StatusCreated)
	} else {
		c.Status(http.StatusBadRequest)
	}
}

// DeleteComment deletes comment from database.
// It checks if user has access to provided page
// and user is allowed to manage provided comment.
func DeleteComment(c *gin.Context) {
	var comment pages.CommentDelete
	if err := c.BindJSON(&comment); err != nil {
		return
	}
	user, success := store.GetUserByToken(util.GetToken(c))
	if success {
		if store.DeleteCommentById(&comment, &user) {
			c.Status(http.StatusOK)
		} else {
			c.Status(http.StatusAlreadyReported)
		}
	} else {
		c.Status(http.StatusBadRequest)
	}
}
