package routers

import (
	"bill_serve/global"

	"github.com/gin-gonic/gin"
)

type RouterGroup struct {
	*gin.RouterGroup
}

func InitRouter() *gin.Engine {
	gin.SetMode(global.Config.System.Env)
	router := gin.Default()
	apiRouterGroup := router.Group("api")

	RouterGroupApp := RouterGroup{apiRouterGroup}
	// 系统配置api
	RouterGroupApp.SettingsRouter()
	RouterGroupApp.ImagesRouter()
	return router
}
