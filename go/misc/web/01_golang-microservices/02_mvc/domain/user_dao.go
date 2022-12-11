package domain

import (
	"fmt"
	"net/http"

	"mvc/utils"
)

var (
	users = map[int64]*User{
		123: {Id: 123, FirstName: "Fede", LastName: "Leon", Email: "myemail@gmail.com"},
	}
)

type UsersDao struct {

}

func (ud *UsersDao) GetUser(userId int64) (*User, *utils.ApplicationError) {
	if user := users[userId]; user != nil {
		return user, nil
	}
	return nil, &utils.ApplicationError{
		Message: fmt.Sprintf("user %v was not found", userId),
		StatusCode: http.StatusNotFound,
		Code: "not found",
	}
}
