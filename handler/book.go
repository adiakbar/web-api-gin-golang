package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"web-api-gin-golang/book"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type bookHandler struct {
	bookService book.Service
}

func NewBookHandler(bookService book.Service) *bookHandler {
	return &bookHandler{bookService}
}

func (h *bookHandler) GetAllBooks(c *gin.Context) {
	books, err := h.bookService.FindAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	var booksResponse []book.BookResponse

	for _, bookObj := range books {
		bookResponse := convertToBookResponse(bookObj)

		booksResponse = append(booksResponse, bookResponse)
	}

	c.JSON(http.StatusOK, gin.H{
		"data": booksResponse,
	})
}

func (h *bookHandler) GetBookById(c *gin.Context) {
	idString := c.Param("id")
	id, _ := strconv.Atoi(idString)

	bookObj, err := h.bookService.FindByID(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	bookResponse := convertToBookResponse(bookObj)

	c.JSON(http.StatusOK, gin.H{"data": bookResponse})
}

func (h *bookHandler) CreateBook(c *gin.Context) {
	var BookRequest book.BookRequest

	err := c.ShouldBindJSON(&BookRequest)
	if err != nil {
		errorMessages := []string{}
		for _, e := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Error on field %s, condition: %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)

		}
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": errorMessages,
		})
		return
	}

	bookObj, err := h.bookService.Create(BookRequest)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": convertToBookResponse(bookObj),
	})
}

func (h *bookHandler) UpdateBook(c *gin.Context) {
	var BookRequest book.BookRequest

	err := c.ShouldBindJSON(&BookRequest)
	if err != nil {
		errorMessages := []string{}
		for _, e := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Error on field %s, condition: %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)

		}
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": errorMessages,
		})
		return
	}

	idString := c.Param("id")
	id, _ := strconv.Atoi(idString)

	bookObj, err := h.bookService.Update(id, BookRequest)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": convertToBookResponse(bookObj),
	})
}

func convertToBookResponse(bookObj book.Book) book.BookResponse {
	return book.BookResponse{
		ID:          bookObj.ID,
		Title:       bookObj.Title,
		Price:       bookObj.Price,
		Description: bookObj.Description,
		Rating:      bookObj.Rating,
		Discount:    bookObj.Discount,
	}
}
