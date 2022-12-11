package services

import (
	"mvc/domain"
	"mvc/utils"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

type UsersDaoMock struct {
	getUserMock func(userId int64) (*domain.User, *utils.ApplicationError)
}

func (ud *UsersDaoMock) GetUser(userId int64) (*domain.User, *utils.ApplicationError) {
	return ud.getUserMock(userId)
}

func TestGetUserNotFoundInDatabase(t *testing.T) {
	getUserWhenNotExists := func(userId int64) (*domain.User, *utils.ApplicationError) {
		return nil, &utils.ApplicationError{
			StatusCode: http.StatusNotFound,
			Message: "user 0 was not found",
		}
	}
	usersService := UsersService{&UsersDaoMock{getUserWhenNotExists}}
	user, err := usersService.GetUser(0)
	assert.Nil(t, user)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusNotFound, err.StatusCode)
	assert.EqualValues(t, "user 0 was not found", err.Message)
}

func TestGetUserFoundInDatabase(t *testing.T) {
	getUserWhenExists := func(userId int64) (*domain.User, *utils.ApplicationError) {
		return &domain.User{
			Id: 123,
		}, nil
	}
	usersService := UsersService{&UsersDaoMock{getUserWhenExists}}
	user, err := usersService.GetUser(0)
	assert.NotNil(t, user)
	assert.Nil(t, err)
	assert.EqualValues(t, 123, user.Id)
}