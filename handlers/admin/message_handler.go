package admin


import (
	"github.com/gin-gonic/gin"
	"net/http"
)


func MessageHandler(c *gin.Context){

	c.HTML(http.StatusOK, "admin-message.html", nil)

}


