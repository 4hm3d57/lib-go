package handlers

import (
	"lib/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"log"
)

func SignupHandler(c *gin.Context) {
	if c.Request.Method != http.MethodPost {
		c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "invalid method request"})
		return
	}
	
	err := c.Request.ParseForm()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error parsing data form"})
		log.Println("Error parsing form:", err)
		return
	}
	
	name := c.PostForm("name")
	roll_no := c.PostForm("roll_no")
	password := c.PostForm("password")
	acc_type := c.PostForm("acc_type")
	
	log.Printf("Added data => name: %s, roll_no: %s, password: %s, acc_type: %s", name, roll_no, password, acc_type)
	
	if name == "" || roll_no == "" || password == "" || acc_type == "" {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "all fields are required"})
		return
	}
	
	newUser := models.User{
		Name:     name,
		Roll_no:  roll_no,
		Password: password,
		Acc_type: acc_type,
	}
	
	err = models.InsertUser(newUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error adding user"})
		log.Println("Error inserting user:", err)
		return
	}
	

	// c.JSON(http.StatusOK, gin.H{"message": "user added successfully"})

	switch newUser.Acc_type {
	case "admin":
		c.Redirect(http.StatusSeeOther, "/admin-index")
	// case "student":
	// 	c.Redirect(http.StatusSeeOther, "/index")
	default:
		c.JSON(http.StatusForbidden, gin.H{"error": "unknown account type"})
	}
	
}
