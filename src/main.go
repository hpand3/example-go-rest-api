package main

import (
	"github.com/gin-gonic/gin"
	"gorestapi/src/httpd"
	"gorestapi/src/repository"
)

func main() {
	feedRepo := repository.NewItemInMemoryRepository()
	r := gin.Default()

	r.GET("/ping", httpd.Ping())

	r.GET("/feed", httpd.GetAllItems(feedRepo))
	r.POST("/feed", httpd.AddItem(feedRepo))

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
