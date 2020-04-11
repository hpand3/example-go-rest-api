package handler

import (
	"github.com/gin-gonic/gin"
	"gorestapi/platform/gorestapi"
)

func FeedGet(feedRepo *gorestapi.Repo) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, feedRepo.GetAll())
	}
}