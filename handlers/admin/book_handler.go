package admin

import (
	"github.com/gin-gonic/gin"
	"lib/models"
	"net/http"
)

func BookHandler(c *gin.Context) {

	books, err := models.GetAllBooks()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error retrieving books"})
		return
	}

	if books == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "books not found"})
		return
	}

	c.HTML(http.StatusOK, "all-books.html", gin.H{"books": books})

}
