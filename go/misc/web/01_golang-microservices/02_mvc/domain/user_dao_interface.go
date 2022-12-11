package domain

import "mvc/utils"

type UsersDaoInterface interface {
	GetUser(userId int64) (*User, *utils.ApplicationError)
}
