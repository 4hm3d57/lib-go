package admin

import (
	"github.com/gin-gonic/gin"
	"lib/models"
	"net/http"
)

// [LOG]: Admin message code
// -------------------------------------------------
func MessageHandler(c *gin.Context) {

	messages, err := models.GetAllMessages()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error retrieving messages"})
		return
	}

	if messages == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "message not found"})
		return
	}

	c.HTML(http.StatusOK, "admin-message.html", gin.H{"messages": messages})

}
