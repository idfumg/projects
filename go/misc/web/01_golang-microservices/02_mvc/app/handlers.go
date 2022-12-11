package app

import (
	"mvc/controllers"
)

func mapHandlers() {
	router.GET("/users/:user_id", controllers.GetUser)
}
