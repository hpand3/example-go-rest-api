package httpd

import (
	"github.com/gin-gonic/gin"
	"gorestapi/src/handler"
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

func GetAllItems(h *handler.GetAllFeedsQueryHandler) gin.HandlerFunc {
	return func(c *gin.Context) {
		var getAllFeedsQuery handler.GetAllFeedsQuery

		res := h.Handle(getAllFeedsQuery)

		c.JSON(200, res)
	}
}
