package student

import (
	"github.com/gin-gonic/gin"
	"lib/models"
	"log"
	"net/http"
)

func InsertMessageHandler(c *gin.Context) {
	// Ensure the request is a POST request
	if c.Request.Method != http.MethodPost {
		c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Invalid method request"})
		return
	}

	// Parse form data
	err := c.Request.ParseForm()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error parsing data form"})
		log.Println("Error parsing form:", err)
		return
	}

	// Validate data
	name := c.PostForm("name")
	roll_no := c.PostForm("roll_no")
	message := c.PostForm("message")

	log.Printf("Retrieved data: name: %s, roll_no: %s, message: %s", name, roll_no, message)

	if name == "" || roll_no == "" || message == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "All fields are required"})
		return
	}

	newMessage := models.Message{
		Name:     name,
		Roll_no:  roll_no,
		Messages: message,
	}

	err = models.InsertMessage(newMessage)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error adding recommendation"})
		return
	}

}
