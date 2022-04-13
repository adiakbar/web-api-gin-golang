package main

import (
	"fmt"
	"log"
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

	books, err := bookRepository.FindAll()
	fmt.Printf("the object book %v", books)

	book := book.Book{
		Title:       "$100 Startup",
		Description: "Good Book",
		Price:       95000,
		Discount:    5,
		Rating:      4,
	}

	bookRepository.Create(book)

	router := gin.Default()

	v1 := router.Group("/v1")

	v1.GET("/", handler.RootHandler)
	v1.GET("/hello", handler.HelloHandler)
	v1.GET("/books/:id", handler.BooksHandler)
	v1.GET("/query", handler.QueryHandler)

	v1.POST("/books", handler.PostBooksHandler)

	router.Run(":8888")
}
