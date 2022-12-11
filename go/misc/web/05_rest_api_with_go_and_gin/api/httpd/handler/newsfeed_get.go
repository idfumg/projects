package handler

import (
	"net/http"
	"newsfeeder_api/platform/newsfeed"

	"github.com/gin-gonic/gin"
)

func NewsfeedGet(feed newsfeed.Getter) gin.HandlerFunc {
	return func(c *gin.Context) {
		items := feed.GetAll()
		c.JSON(http.StatusOK, items)
	}
}
