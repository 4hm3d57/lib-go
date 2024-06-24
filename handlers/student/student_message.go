package student


import(
	"lib/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"log"
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
	roll_no := c.PostForm("roll_no")
	message := c.PostForm("message")

	
	log.Printf("Retrieved data: roll_no: %s, message: %s", roll_no, message)
	
	if roll_no == "" || message == ""  {
		c.JSON(http.StatusBadRequest, gin.H{"error": "All fields are required"})
		return
	}

	newMessage := models.Message{
		Roll_no: roll_no,
		Messages: message,
	}

	err = models.InsertMessage(newMessage)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error adding recommendation"})
		return
	}
	
}





// [LOG]: Admin message code
// -------------------------------------------------
// func StudentMessageHandler(c *gin.Context) {
	
// 	messages, err := models.GetAllMessage()
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "error retrieving messages"})
// 		return
// 	}
	
// 	if messages == nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "message not found"})
// 		return
// 	}
	
// 	c.HTML(http.StatusOK, "message.html", gin.H{"messages": messages})
	
// }
