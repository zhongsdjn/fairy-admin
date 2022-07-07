package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	err := r.Run("127.0.0.1:8080")
	if err != nil {
		return
	}
}
