package apps

import (
	"github.com/gin-gonic/gin"
	"github.com/zhongsdjn/apps/system"
	"github.com/zhongsdjn/apps/web"
)

func AppsRouters(rg *gin.RouterGroup) *gin.RouterGroup {

	{
		systemGroup := rg.Group("system/")
		systemGroup.Use()
		system.SystemRouter(systemGroup)

	}

	{
		WebGroup := rg.Group("web/")
		WebGroup.Use()
		web.WebRouter(WebGroup)
	}

	return rg
}
