package main

import (
	"fmt"
	handler "url_shortner/pkg/handler"
	store "url_shortner/pkg/store"

	"github.com/gin-gonic/gin"
)

func hello(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hey Go URL Shortener !",
	})
}

func main() {
	r := gin.Default()
	r.GET("/", hello)

	r.POST("/create-short-url", func(c *gin.Context) {
		handler.CreateShortUrl(c)
	})

	r.GET("/:shortUrl", func(c *gin.Context) {
		handler.HandleShortUrlRedirect(c)
	})

	// Note that store initialization happens here
	store.InitializeStore()

	err := r.Run(":9808")
	if err != nil {
		panic(fmt.Sprintf("Failed to start the web server - Error: %v", err))
	}

}
