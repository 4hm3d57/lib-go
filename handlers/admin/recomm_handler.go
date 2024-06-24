package admin


import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func RecommendationHandler(c *gin.Context){

	c.HTML(http.StatusOK, "admin-recommendation.html", nil)

}


