package admin


import (
	"lib/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

// admin recommendation code
// -----------------------------------------------------------------------
func RecommendationHandler(c *gin.Context){

	recommendations, err := models.GetAllRecommendation()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error retrieving recommendations"})
		return
	}
	
	c.HTML(http.StatusOK, "admin-recommendation.html", gin.H{"recommendations": recommendations})

}
