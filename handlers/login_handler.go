package handlers

import (
	"lib/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"log"
)

func LoginHandler(c *gin.Context) {
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
	password := c.PostForm("password")
	
	log.Printf("Retrieved data: name: %s, roll_no: %s, password: %s", name, roll_no, password)
	
	if name == "" || roll_no == "" || password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "All fields are required"})
		return
	}
	
	// Fetch user from the database
	user, err := models.GetUser(name, roll_no, password)
	if err != nil {
		log.Println("Error retrieving data:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error retrieving data"})
		return
	}

	if user == nil {
		log.Println("User not found")
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	
	// Redirect based on account type
	switch user.Acc_type {
	case "admin":
		c.Redirect(http.StatusSeeOther, "/admin-index")
	case "student":
		c.Redirect(http.StatusSeeOther, "/student-index")
	default:
		c.JSON(http.StatusForbidden, gin.H{"error": "Unknown account type"})
	}
}
