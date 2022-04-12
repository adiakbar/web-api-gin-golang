package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/", rootHandler)
	router.GET("/hello", helloHandler)

	router.Run(":8888")
}

func rootHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"name": "Adi Akbar",
		"bio":  "A Software Engineer",
	})
}

func helloHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"title":    "Hello World",
		"subtitle": "Belajar Golang bareng Youtube Agung Setiawan",
	})
}
