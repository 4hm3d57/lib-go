package student


import(
	"lib/models"
	"github.com/gin-gonic/gin"
	"net/http"
)


func StudentBookHandler(c *gin.Context){
	books, err := models.GetAllBooks()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error retrieving users"})
		return
	}

	if books == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Books not found"})
		return
	}

	c.HTML(http.StatusOK, "student-all-books.html", gin.H{"books": books})
	
}
