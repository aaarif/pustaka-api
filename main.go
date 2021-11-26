package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func main() {
	router := gin.Default()
	v1 := router.Group("/v1")
	v1.GET("/", rootHandler)
	v1.GET("/hello", helloHandler)
	v1.GET("/books/:id/:title", booksHandler)
	v1.GET("/query", queryHandler)
	v1.POST("/books", postBooksHandler)

	router.Run()
}

func rootHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"name": "  Riswanto",
		"bio":  "Software Engineer and visionary",
	})
}

func helloHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"content":  "Belajar GOLANG",
		"subtitle": "Bersama Arif Riswanto",
	})
}

func booksHandler(c *gin.Context) {
	// books/1
	id := c.Param("id")
	title := c.Param("title")
	c.JSON(http.StatusOK, gin.H{
		"id":    id,
		"title": title,
	})
}

func queryHandler(c *gin.Context) {
	title := c.Query("title")
	price := c.Query("price")
	c.JSON(http.StatusOK, gin.H{
		"title": title,
		"price": price,
	})
}

type BookInput struct {
	Title string      `json:"title" binding:"required"`
	Price json.Number `json:"price" binding:"required,number"`
	// Subtitle string `json:"sub_title"` do this for mapping, different
	Subtitle string
	Email    string `json:"email" binding:"email"`
}

func postBooksHandler(c *gin.Context) {
	// title and price
	var bookInput BookInput
	err := c.ShouldBindJSON(&bookInput)
	if err != nil {
		errorMessages := []string{}
		for _, e := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Error on field: %s, condition: %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)

		}
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": errorMessages,
		})
		return

		// log.Fatal(err) coused server shutdown
	}

	c.JSON(http.StatusOK, gin.H{
		"title":    bookInput.Title,
		"price":    bookInput.Price,
		"subtitle": bookInput.Subtitle,
		"email":    bookInput.Email,
	})
}
