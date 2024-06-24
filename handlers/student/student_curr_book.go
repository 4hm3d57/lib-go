package student


import(
	"github.com/gin-gonic/gin"
	"net/http"
)


func StudentCurrHandler(c *gin.Context){

	c.HTML(http.StatusOK, "curr-books.html", nil)
	
}
