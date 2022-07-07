package fairy_admin

import (
	"github.com/gin-gonic/gin"
	"github.com/zhongsdjn/apps"
)

func Routers() *gin.Engine {
	r := gin.Default()

	AppsRouters := apps.AppsRouters
	AppsRouters(r.Group(""))

	return r
}
