package services

import (
	"mvc/domain"
	"mvc/utils"
)

type UsersServiceInterface interface {
	GetUser(userId int64) (*domain.User, *utils.ApplicationError)
}
