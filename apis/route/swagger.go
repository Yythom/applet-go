package route

import (
	_ "applet/apis/docs"
	"applet/core/bootstrap"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func NewSwaggerRouter(env *bootstrap.Env, group *gin.RouterGroup) {
	group.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
