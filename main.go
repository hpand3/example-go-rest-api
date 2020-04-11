package main

import (
	"github.com/gin-gonic/gin"
	"gorestapi/httpd/handler"
	"gorestapi/platform/gorestapi"
)

func main() {
	feedRepo := gorestapi.New()
	r := gin.Default()

	r.GET("/ping", handler.PingGet())
	r.GET("/feed", handler.FeedGet(feedRepo))
	r.POST("/feed", handler.FeedPost(feedRepo))

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}