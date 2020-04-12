package httpd

import (
	"github.com/gin-gonic/gin"
	"gorestapi/src/domain"
	"gorestapi/src/repository"
)

type createFeedCommand struct {
	Title string `json:"title"`
	Post  string `json:"post"`
}

func AddItem(feedRepo repository.ItemRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		var requestBody createFeedCommand
		c.Bind(&requestBody)

		item := domain.Item{
			Title: requestBody.Title,
			Post:  requestBody.Post,
		}
		feedRepo.Add(item)

		c.JSON(201, gin.H{
			"message": "Successfully created feed",
		})
	}
}

func GetAllItems(feedRepo repository.ItemRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, feedRepo.GetAll())
	}
}
