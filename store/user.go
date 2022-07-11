package store

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"sophie-server/database"
	"sophie-server/model/users"
	"sophie-server/util"
	"strings"
	"time"
)

// GetUserByNickname returns users instance by provided nickname
// if it was found in database.
func GetUserByNickname(nickname string) (users.User, bool) {
	statement, _ := database.GetUsersDB().Prepare(database.GET_USER_BY_NICKNAME)
	rows, _ := statement.Query(nickname)
	var user users.User
	for rows.Next() {
		err := rows.Scan(&user.ID, &user.Firstname, &user.Lastname, &user.Nickname, &user.Password, &user.Email, &user.RegisterDate, &user.EmailVerified, &user.Sessions, &user.Workspaces)
		if err != nil {
			fmt.Println(err)
			return users.User{}, false
		}
	}
	return user, true
}

// GetUserByID returns users instance by provided id
// if it was found in database.
func GetUserByID(id int) (users.User, bool) {
	statement, _ := database.GetUsersDB().Prepare(database.GET_USER_BY_ID)
	rows, _ := statement.Query(id)
	var user users.User
	for rows.Next() {
		err := rows.Scan(&user.ID, &user.Firstname, &user.Lastname, &user.Nickname, &user.Password, &user.Email, &user.EmailVerified, &user.Workspaces)
		if err != nil {
			return users.User{}, false
		}
	}
	return user, true
}

// GetAuthedUser returns users instance (model.UserGet) by provided token.
func GetAuthedUser(token string) (users.UserGet, bool) {
	if user, success := GetUserByToken(token); success {
		return user.UserToUserGet(), true
	}
	return users.UserGet{}, false
}

// GetUserByToken returns users instance (model.User) by provided token.
func GetUserByToken(token string) (users.User, bool) {
	claims, err := util.ParseToken(token)
	if err != nil {
		return users.User{}, false
	}
	if user, success := GetUserByNickname(claims["iss"].(string)); success {
		return user, true
	}
	return users.User{}, false
}

// CreateUser Creates users instance and adds it to database
func CreateUser(model users.UserCreate) gin.H {
	if len(model.Firstname) < 2 || len(model.Lastname) < 2 || len(model.Password) < 8 || len(model.Nickname) < 3 {
		return gin.H{"error": "invalid data"}
	}
	statement, _ := database.GetUsersDB().Prepare(database.CREATE_USER)
	password := util.GenerateHash([]byte(model.Password))
	_, err := statement.Exec(model.Firstname, model.Lastname, model.Nickname, password, model.Email, time.Now().Format("2006-01-02 15:04:05"))
	if err != nil {
		array := strings.Split(err.Error(), ".")
		return gin.H{"error": fmt.Sprintf("an users with this %s already exist", array[len(array)-1])}
	}
	return gin.H{"message": "a new users was successfully created"}
}
