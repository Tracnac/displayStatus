package main

import (
	"embed"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	static "github.com/soulteary/gin-static"
)

//go:embed html
var EmbedFS embed.FS

var sampleJson = `[
	{
		"id": 1,
	"name": "John Doe",
	"age": 25,
	"email": "",
	"phone": "123-456-7890"
},
{
	"id": 2,
	"name": "Jane Smith",
	"age": 30,
	"email": "",
	"phone": "123-456-7890"
}
]`

func main() {
	r := gin.Default()

	r.Use(func(c *gin.Context) {
		c.Header("X-Content-Type-Options", "nosniff")
		c.Next()
	})

	r.GET("/api/json", func(c *gin.Context) {
		get := c.Query("get")
		if get == "sample" {
			c.String(200, sampleJson)
		} else {
			c.String(200, "{"+get+": \"not found\"}")
		}
	})

	r.Use(static.ServeEmbed("html", EmbedFS))

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("Listening on port %s", port)
	r.Run(":" + port)
}
