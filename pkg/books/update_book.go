package books

import (
	"example/postgres-connect/pkg/common/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UpdateBookRequestBody struct {
	Title       string `json:"title"`
	Author      string `json:"author"`
	Description string `json:"description"`
}

func (h handler) UpdateBook(c *gin.Context) {
	id := c.Param("id")
	body := UpdateBookRequestBody{}

	// getting request's body
	if err := c.BindJSON(&body); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var book models.Book

	if result := h.DB.First(&book, id); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	book.Title, book.Author, book.Description = body.Title, body.Author, body.Description

	h.DB.Updates(&book)

	//h.DB.Save(&book)

	c.JSON(http.StatusOK, &book)
}
