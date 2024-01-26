package route

import (
	"applet/apis/controller"
	"applet/core/bootstrap"
	"applet/core/mongo"
	"applet/domain"
	"applet/repository"
	"applet/usecase"
	"github.com/gin-gonic/gin"
	"time"
)

func NewPermissionRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewPermissionRepository(db, domain.CollectionUser)
	lc := &controller.PermissionController{
		PermissionUseCase: usecase.NewPermissionUsecase(ur, timeout),
		Env:               env,
	}
	group.POST("/permission/create", lc.CreatePermission)
	group.DELETE("/permission/delete", lc.DeletePermission)
	group.GET("/permission", lc.GetPermission)
	group.GET("/permission/list", lc.GetPermissionList)
}
