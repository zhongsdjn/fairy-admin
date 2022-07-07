package user

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type User struct {
}

func UserGet(c *gin.Context) {

	var u User
	err := c.ShouldBind(&u)
	if err != nil {
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"mes":  "ok",
		"info": "pxt",
	})
}
