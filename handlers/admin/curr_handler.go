package admin


import(
	"github.com/gin-gonic/gin"
	"net/http"
)


func CurrHandler(c *gin.Context){

	c.HTML(http.StatusOK, "admin-curr-books.html", nil)

}