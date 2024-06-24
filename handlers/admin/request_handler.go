package admin


import(
	"github.com/gin-gonic/gin"
	"net/http"
)


func RequestHandler(c *gin.Context){

	c.HTML(http.StatusOK, "request.html", nil)

}