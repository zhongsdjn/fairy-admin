package example

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Eaxmple(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{
		"mes":  "ok",
		"info": "web",
	})
}
