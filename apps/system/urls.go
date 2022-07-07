package system

import (
	"github.com/gin-gonic/gin"
	"github.com/zhongsdjn/apps/system/role"
	"github.com/zhongsdjn/apps/system/user"
	//"github.com/zhongsdjn/apps/system/user"
)

//type RouterGroup struct {
//	user.UserRouter
//}

//func InitSystemRouter(Router *gin.RouterGroup) {
//	userRouter := Router.Group("users/")
//	//userRouter.Use()
//	{
//		userRouter.GET("list", UserGet) // jwt加入黑名单
//	}
//}

func SystemRouter(rg *gin.RouterGroup) *gin.RouterGroup {

	{
		usersGroup := rg.Group("users/")
		usersGroup.Use()
		user.UserRouter(usersGroup)
	}

	{
		rolesGroup := rg.Group("roles/")
		rolesGroup.Use()
		role.RoleRouter(rolesGroup)
	}

	return rg
}
