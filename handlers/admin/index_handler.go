package admin

import (
	"github.com/gin-gonic/gin"
	"lib/models"
	"log"
	"net/http"
)

func IndexHandler(c *gin.Context) {

	index, err := models.GetAllUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error retrieving messages"})
		return
	}

	if index == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error"})
		return
	}

	log.Println("Users retrived: ", index)

	// Render the template with the user data
	c.HTML(http.StatusOK, "index.html", gin.H{"index": index})

}
