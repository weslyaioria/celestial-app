package model

type Book struct {
	Oid  string `json:"oid" binding:"required"`
	Name string `json:"name"`
}

type TotalRows struct {
	CountRows int `json:"totalRows" db:"totalRows"`
}
