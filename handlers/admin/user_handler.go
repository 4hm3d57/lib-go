package admin

import (
	"lib/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"log"
)

func RenderUserHandler(c *gin.Context) {
	users, err := models.GetAllUsers()
	if err != nil {
		c.String(http.StatusInternalServerError, "Error retrieving users")
		log.Printf("Error retrieving users: %v", err)
		return
	}

	// Render the template with the user data
	c.HTML(http.StatusOK, "users.html", gin.H{
		"users": users,
	})
}
