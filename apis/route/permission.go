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
	ur := repository.NewPermissionRepository(db, domain.CollectionPermission)
	lc := &controller.PermissionController{
		PermissionUseCase: usecase.NewPermissionUsecase(ur, timeout),
		Env:               env,
	}
	group.GET("/permissions/:permissionId", lc.GetPermission)
	group.PUT("/permissions/:permissionId", lc.UpdatePermission)
	group.DELETE("/permissions/:permissionId", lc.DeletePermission)
	group.POST("/permissions", lc.CreatePermission)
	group.GET("/permissions", lc.GetPermissionList)
}
