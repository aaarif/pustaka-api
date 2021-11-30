package main

import (
	"fmt"
	"log"
	"pustaka-api/book"
	"pustaka-api/handler"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {

	dsn := "root:root@tcp(127.0.0.1:3306)/pustaka-api?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("DB Connection Error")
	}
	fmt.Println("Database connected.")
	db.AutoMigrate(&book.Book{})
	fmt.Println("Automigration succeeded.")

	// book := book.Book{}
	// book.Title = "The Kite Runner"
	// book.Price = 90000
	// book.Discount = 7
	// book.Rating = 5
	// book.Description = "Written by Khalid Hussaini"

	// err = db.Create(&book).Error
	// if err != nil {
	// 	fmt.Println("==============================")
	// 	fmt.Println("====Error creating book=======")
	// 	fmt.Println("==============================")
	// }

	var book book.Book
	err = db.First(&book).Error
	if err != nil {
		fmt.Println("==============================")
		fmt.Println("====Error Retrieving book=======")
		fmt.Println("==============================")
	}
	fmt.Println("Title: ", book.Title)

	router := gin.Default()
	v1 := router.Group("/v1")
	v1.GET("/", handler.RootHandler)
	v1.GET("/hello", handler.HelloHandler)
	v1.GET("/books/:id/:title", handler.BooksHandler)
	v1.GET("/query", handler.QueryHandler)
	v1.POST("/books", handler.PostBooksHandler)

	router.Run()
}
