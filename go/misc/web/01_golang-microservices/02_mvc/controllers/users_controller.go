package controllers

import (
	"net/http"
	"strconv"

	"mvc/domain"
	"mvc/services"
	"mvc/utils"

	"github.com/gin-gonic/gin"
)

func getUser(usersServiceInterface services.UsersServiceInterface, userId int64) (*domain.User, *utils.ApplicationError) {
	return usersServiceInterface.GetUser(userId)
}

func GetUser(c *gin.Context) {
	usersService := &services.UsersService{
		UsersDaoInterface: &domain.UsersDao{},
	}

	if userId, err := strconv.ParseInt(c.Param("user_id"), 10, 64); err != nil {
		apiErr := &utils.ApplicationError{
			Message:    "user_id must be a number",
			StatusCode: http.StatusBadRequest,
			Code:       "bad_request",
		}
		utils.MakeResponse(c, apiErr.StatusCode, apiErr)
	} else if user, apiErr := getUser(usersService, userId); apiErr != nil {
		utils.MakeResponse(c, apiErr.StatusCode, apiErr)
	} else {
		utils.MakeResponse(c, http.StatusOK, user)
	}
}
