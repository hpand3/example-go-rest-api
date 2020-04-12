package httpd

import (
	"github.com/gin-gonic/gin"
	"gorestapi/src/handler"
	"net/http"
)

func AddItem(h *handler.CreateFeedItemHandler) gin.HandlerFunc {
	return func(c *gin.Context) {
		var createFeedItemCommand handler.CreateFeedItemCommand
		err := c.Bind(&createFeedItemCommand)
		if err != nil {
			return
		}

		h.Handle(createFeedItemCommand)

		c.Status(http.StatusCreated)
	}
}

func GetAllItems(h *handler.GetAllFeedsQueryHandler) gin.HandlerFunc {
	return func(c *gin.Context) {
		res := h.Handle()

		c.JSON(http.StatusOK, res)
	}
}
