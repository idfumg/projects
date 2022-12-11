package utils

import "github.com/gin-gonic/gin"

func MakeResponse(c *gin.Context, status int, body interface{}) {
	if c.GetHeader("Accept") == "application/xml" {
		c.XML(status, body)
	} else {
		c.JSON(status, body)
	}
}