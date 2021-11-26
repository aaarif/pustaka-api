package book

import "encoding/json"

type BookInput struct {
	Title string      `json:"title" binding:"required"`
	Price json.Number `json:"price" binding:"required,number"`
	// Subtitle string `json:"sub_title"` do this for mapping, different
	Subtitle string
	Email    string `json:"email" binding:"email"`
}
