package main

import (
	"log"
	"net/http"
	"web-api-gin-golang/book"
	"web-api-gin-golang/handler"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:@tcp(127.0.0.1:3306)/web_api_gin_golang?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("DB connection error")
	}

	db.AutoMigrate(&book.Book{})

	bookRepository := book.NewRepository(db)
	bookService := book.NewService(bookRepository)
	bookHandler := handler.NewBookHandler(bookService)

	router := gin.Default()

	v1 := router.Group("/v1")

	v1.GET("/", rootHandler)
	v1.GET("/query", queryHandler)

	v1.GET("/books", bookHandler.GetAllBooks)
	v1.GET("/books/:id", bookHandler.GetBookById)
	v1.POST("/books", bookHandler.CreateBooks)

	router.Run(":8888")
}

func rootHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"name": "Adi Akbar",
		"bio":  "A Software Engineer",
	})
}

func queryHandler(c *gin.Context) {
	title := c.Query("title")

	c.JSON(http.StatusOK, gin.H{"title": title})
}
