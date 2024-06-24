package student

import(
	"github.com/gin-gonic/gin"
	"net/http"
)


func StudentIndexHandler(c *gin.Context){

	c.HTML(http.StatusOK, "student_index.html", nil)
	
}
