package store

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"sophie-server/database"
	"sophie-server/model"
	"sophie-server/service"
	"sophie-server/util"
	"strings"
	"time"
)

// CreateUser Creates user instance and adds it to database
func CreateUser(model model.UserCreate) gin.H {
	if len(model.Firstname) < 2 || len(model.Lastname) < 2 || len(model.Password) < 8 || len(model.Nickname) < 3 {
		return gin.H{"error": "invalid data"}
	}
	statement, _ := database.GetUsersDB().Prepare(database.CREATE_USER)
	password := util.GenerateHash([]byte(model.Password))
	_, err := statement.Exec(model.Firstname, model.Lastname, model.Nickname, password, model.Email, time.Now().Format("2006-01-02 15:04:05"))
	if err != nil {
		array := strings.Split(err.Error(), ".")
		return gin.H{"error": fmt.Sprintf("an user with this %s already exist", array[len(array)-1])}
	}
	GetUserByNickname(model.Nickname)
	return gin.H{"message": "a new user was successfully created"}
}

// GetUserByNickname returns user instance by provided nickname
// if it was found in database.
func GetUserByNickname(nickname string) (model.User, bool) {
	statement, _ := database.GetUsersDB().Prepare(database.GET_USER_BY_NICKNAME)
	rows, _ := statement.Query(nickname)
	var user model.User
	for rows.Next() {
		err := rows.Scan(&user.ID, &user.Firstname, &user.Lastname, &user.Nickname, &user.Password, &user.Email, &user.EmailVerified, &user.Workspaces)
		if err != nil {
			return model.User{}, false
		}
	}
	return user, true
}

// GetAuthedUser returns user instance (model.UserGet) by provided token.
func GetAuthedUser(token string) (model.UserGet, bool) {
	claims, err := service.ParseToken(token)
	if err != nil {
		return model.UserGet{}, false
	}
	user, success := GetUserByNickname(claims["iss"].(string))
	if success {
		return model.UserToUserGet(user), true
	}
	return model.UserGet{}, false
}
