package handler

import (
	"github.com/gin-gonic/gin"
	"gorestapi/platform/gorestapi"
)

type createFeedCommand struct {
	Title string `json:"title"`
	Post string `json:"post"`
}

func FeedPost(feedRepo *gorestapi.Repo) gin.HandlerFunc {
	return func(c *gin.Context) {
		var requestBody createFeedCommand
		c.Bind(&requestBody)

		item := gorestapi.Item{
			Title: requestBody.Title,
			Post: requestBody.Post,
		}
		feedRepo.Add(item)

		c.JSON(201, gin.H{
			"message": "Successfully created feed",
		})
	}
}