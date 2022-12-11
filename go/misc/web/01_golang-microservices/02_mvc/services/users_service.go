package services

import (
	"mvc/domain"
	"mvc/utils"
)

type UsersService struct {
	UsersDaoInterface domain.UsersDaoInterface
}

func (us *UsersService) GetUser(userId int64) (*domain.User, *utils.ApplicationError) {
	if user, err := us.UsersDaoInterface.GetUser(userId); err == nil {
		return user, nil
	} else {
		return nil, err
	}
}
