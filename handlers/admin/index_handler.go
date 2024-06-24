package admin

import ("github.com/gin-gonic/gin"
	
	"net/http"
)

func IndexHandler(c *gin.Context) {

	// Render the template with the user data
	c.HTML(http.StatusOK, "index.html", nil)
}
