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
	// book.Title = "Influencer Economy"
	// book.Price = 100000
	// book.Discount = 5
	// book.Rating = 4
	// book.Description = "Academy publisher"

	// err = db.Create(&book).Error
	// if err != nil {
	// 	fmt.Println("==============================")
	// 	fmt.Println("====Error creating book=======")
	// 	fmt.Println("==============================")
	// }

	var book2 book.Book
	err = db.Where("title LIKE ? ", "Influencer%").First(&book2).Error
	if err != nil {
		fmt.Println("==============================")
		fmt.Println("====Error Retrieving book for update =======")
		fmt.Println("==============================")
	}
	fmt.Println("Book to update: ", book2.Title)
	book2.Title = "Influencer Economy 2nd Edition"
	err = db.Save(&book2).Error
	if err != nil {
		fmt.Println("==============================")
		fmt.Println("====Error updating book=======")
		fmt.Println("==============================")
	}

	// var books []book.Book
	// err = db.Debug().Where("title LIKE ? ", "Influencer%").Find(&books).Error
	// if err != nil {
	// 	fmt.Println("==============================")
	// 	fmt.Println("====Error Retrieving all books with title Influencer=======")
	// 	fmt.Println("==============================")
	// }

	// for _, b := range books {
	// 	fmt.Println("Book Title like Influencer: ", b.Title)
	// 	fmt.Println("Book object %v", b)
	// }
	router := gin.Default()
	v1 := router.Group("/v1")
	v1.GET("/", handler.RootHandler)
	v1.GET("/hello", handler.HelloHandler)
	v1.GET("/books/:id/:title", handler.BooksHandler)
	v1.GET("/query", handler.QueryHandler)
	v1.POST("/books", handler.PostBooksHandler)

	router.Run()
}
