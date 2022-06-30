package store

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"sophie-server/database"
	"sophie-server/model"
	"sophie-server/util"
	"strings"
)

// CreateUser Creates user instance and adds it to database
func CreateUser(model model.UserCreate) gin.H {
	statement, _ := database.GetUsersDB().Prepare(database.CREATE_USER)
	password := util.GenerateHash([]byte(model.Password))
	_, err := statement.Exec(model.Firstname, model.Lastname, model.Nickname, password, model.Email)
	if err != nil {
		array := strings.Split(err.Error(), ".")
		return gin.H{"error": fmt.Sprintf("an user with this %s already exist", array[len(array)-1])}
	}
	return gin.H{"message": "a new user was successfully created"}
}
