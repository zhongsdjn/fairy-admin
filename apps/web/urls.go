package web

import (
	"github.com/gin-gonic/gin"
	"github.com/zhongsdjn/apps/web/example"
)

func WebRouter(rg *gin.RouterGroup) *gin.RouterGroup {

	{
		exampleGroup := rg.Group("example/")
		exampleGroup.Use()
		example.ExampleRouter(exampleGroup)
	}

	return rg
}
