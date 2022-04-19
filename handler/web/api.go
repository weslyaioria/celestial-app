package web

import (
	"fmt"
	"net/http"

	"celestial-app/internal/model"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func RootHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"name": "wesly aioria",
		"bio":  "A Software enginner & contect creator",
	})
}

func HelloHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"name": "wesly aioria",
		"bio":  "A Software enginner & contect creator",
	})
}

func BookIdHandler(c *gin.Context) {
	id := c.Param("id")
	title := c.Param("title")

	c.JSON(http.StatusOK, gin.H{
		"Id":    id,
		"title": title,
	})
}

func QueryHandler(c *gin.Context) {
	id := c.Query("id")
	title := c.Query("title")

	c.JSON(http.StatusOK, gin.H{
		"Id":    id,
		"title": title,
	})
}

func PostBookHandler(c *gin.Context) {
	var book model.Book
	err := c.ShouldBindJSON(&book)

	if err != nil {
		errorMessages := []string{}
		for _, e := range err.(validator.ValidationErrors) {
			fmt.Println(err)
			errorMessage := fmt.Sprintf("Error on filed %s, condition: %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)

		}
		c.JSON(http.StatusBadRequest, errorMessages)
		return

	}
	c.JSON(http.StatusOK, gin.H{
		"title":     book.Title,
		"price":     book.Price,
		"sub_title": book.SubTitle,
		"email":     book.Email,
	})
}
