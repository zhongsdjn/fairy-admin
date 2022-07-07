package user

import (
	"github.com/gin-gonic/gin"
)

func UserRouter(Router *gin.RouterGroup) {
	{
		Router.GET("get", UserGet)
		Router.POST("add", UserGet)
		Router.PUT("change", UserGet)
		Router.DELETE("del", UserGet)
		Router.GET("list", UserGet)
	}
}
