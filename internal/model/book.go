package model

import "encoding/json"

type Book struct {
	Title    string      `json:"title" binding:"required"`
	Price    json.Number `json:"price" binding:"required,number"`
	SubTitle string      `json:"sub_title"`
	Email    string      `json:"email" binding:"required,email"`
}

type TotalRows struct {
	CountRows int `json:"totalRows" db:"totalRows"`
}
