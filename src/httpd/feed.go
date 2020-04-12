package httpd

import (
	"github.com/gin-gonic/gin"
	"gorestapi/src/handler"
	"gorestapi/src/repository"
)

func AddItem(h *handler.CreateFeedItemHandler) gin.HandlerFunc {
	return func(c *gin.Context) {
		var createFeedItemCommand handler.CreateFeedItemCommand
		c.Bind(&createFeedItemCommand)

		h.Handle(createFeedItemCommand)

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
