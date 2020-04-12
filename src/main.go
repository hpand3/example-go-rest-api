package main

import (
	"github.com/gin-gonic/gin"
	"gorestapi/src/handler"
	"gorestapi/src/httpd"
	"gorestapi/src/repository"
)

func main() {
	feedRepo := repository.NewItemInMemoryRepository()

	createFeedItemHandler := handler.NewCreateFeedItemHandler(feedRepo)
	getAllFeedsQueryHandler := handler.NewGetAllFeedsQueryHandler(feedRepo)

	r := gin.Default()

	r.GET("/ping", httpd.Ping())

	r.GET("/feed", httpd.GetAllItems(getAllFeedsQueryHandler))
	r.POST("/feed", httpd.AddItem(createFeedItemHandler))

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
