package admin

import (
	"lib/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"log"
)


func AddBookHandler(c *gin.Context){

	if c.Request.Method != http.MethodPost {
		c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "invalid method request"})
		return
	}

	err := c.Request.ParseForm()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error parsing form data:"})
		return
	}

	title := c.PostForm("title")
	author := c.PostForm("author")
	publisher := c.PostForm("publisher")
	year := c.PostForm("year")
	copies := c.PostForm("copies")

	log.Printf("Added data: title: %s, author: %s, publisher: %s, year: %s, copies: %s", title, author, publisher, year, copies)

	if title == "" || author == "" || publisher == "" || year == "" || copies == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "All fields required"})
		return
	}

	newBook := models.Book{
		Title: title,
		Author: author,
		Publisher: publisher,
		Year: year,
		Copies: copies,
	}
	

	err = models.InsertBook(newBook)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error adding book"})
		log.Println("error adding book: ", err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "book added"})
	
}
