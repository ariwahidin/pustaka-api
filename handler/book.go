package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"pustaka-api/book"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

type bookHandler struct {
	bookService book.Service
}

func NewBookHandler(bookService book.Service) *bookHandler {
	return &bookHandler{bookService}
}

func (h *bookHandler) GetAllBooks(ctx *gin.Context) {
	books, err := h.bookService.FindAll()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	var booksResponse []book.BookResponse

	for _, b := range books {
		booksResponse = append(booksResponse, convertToResponse(b))
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": booksResponse,
	})
}

func (h *bookHandler) GetBook(ctx *gin.Context) {

	idString := ctx.Param("id")
	id, _ := strconv.Atoi(idString)

	books, err := h.bookService.FindByID(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	bookResponse := convertToResponse(books)

	ctx.JSON(http.StatusOK, gin.H{
		"data": bookResponse,
	})
}

func (h *bookHandler) CreateBook(ctx *gin.Context) {
	var bookInput book.BookRequest
	err := ctx.ShouldBindJSON(&bookInput)
	if err != nil {

		errorMessages := []string{}
		for _, e := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Error on field %s, condition %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)
		}

		ctx.JSON(http.StatusBadRequest, gin.H{
			"errors": errorMessages,
		})
		return
	}

	book, err := h.bookService.Create(bookInput)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": convertToResponse(book),
	})
}

func (h *bookHandler) DeleteBook(ctx *gin.Context) {

	var bookRequest book.BookRequest
	err := ctx.ShouldBindJSON(&bookRequest)
	if err != nil {

		errorMessages := []string{}
		for _, e := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Error on field %s, condition %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)
		}

		ctx.JSON(http.StatusBadRequest, gin.H{
			"errors": errorMessages,
		})
		return
	}

	idString := ctx.Param("id")
	id, _ := strconv.Atoi(idString)
	book, err := h.bookService.Delete(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": convertToResponse(book),
	})
}

func (h *bookHandler) UpdateBook(ctx *gin.Context) {

	var bookRequest book.BookRequest
	err := ctx.ShouldBindJSON(&bookRequest)
	if err != nil {

		errorMessages := []string{}
		for _, e := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Error on field %s, condition %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)
		}

		ctx.JSON(http.StatusBadRequest, gin.H{
			"errors": errorMessages,
		})
		return
	}

	idString := ctx.Param("id")
	id, _ := strconv.Atoi(idString)
	book, err := h.bookService.Update(id, bookRequest)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": convertToResponse(book),
	})
}

func (h *bookHandler) QueryHandler(ctx *gin.Context) {
	title := ctx.Query("title")
	category := ctx.Query("category")
	ctx.JSON(http.StatusOK, gin.H{
		"title":    title,
		"category": category,
	})
}

func (h *bookHandler) BooksHandler(ctx *gin.Context) {
	id := ctx.Param("id")
	title := ctx.Param("title")
	ctx.JSON(http.StatusOK, gin.H{
		"id":    id,
		"title": title,
	})
}

func (h *bookHandler) RootHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"name": "Ari Wahidin",
		"bio":  "A IT Programmer and Software Engineer",
	})
}

func (h *bookHandler) HelloHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"title":    "Hello World",
		"subtitle": "Belajar Go lang demi masa depan cerah",
	})
}

func convertToResponse(b book.Book) book.BookResponse {
	bookResponse := book.BookResponse{
		ID:       b.ID,
		Title:    b.Title,
		SubTitle: b.Description,
		Price:    b.Price,
		Rating:   b.Rating,
		Discount: b.Discount,
	}
	return bookResponse
}
