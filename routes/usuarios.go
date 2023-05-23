package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUsuarios(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{"message": "usuario"})
}
