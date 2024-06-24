package student


import(
	"lib/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"log"
)


func InsertRecommHandler(c *gin.Context) {
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
	title := c.PostForm("title")
	description := c.PostForm("description")

	
	log.Printf("Retrieved data: title: %s, description: %s", title, description)
	
	if title == "" || description == ""  {
		c.JSON(http.StatusBadRequest, gin.H{"error": "All fields are required"})
		return
	}

	newRecomm := models.Recommendation{
		Title: title,
		Description: description,
	}

	err = models.InsertRecommendation(newRecomm)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error adding recommendation"})
		return
	}
	
}


// admin recommendation code
// -----------------------------------------------------------------------
// func StudentRecommendationHandler(c *gin.Context){

// 	recommendations, err := models.GetAllRecommendation()
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "error retrieving recommendations"})
// 		return
// 	}
	
// 	if messages == nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "recommendation not found"})
// 		return
// 	}
	
// 	c.HTML(http.StatusOK, "recommendation.html", gin.H{"recommendations": recommendations})

// }
