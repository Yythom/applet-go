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

func NewRoleRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	rr := repository.NewRoleRepository(db, domain.CollectionRole)
	ur := repository.NewUserRepository(db, domain.CollectionUser)
	pr := repository.NewPermissionRepository(db, domain.CollectionPermission)
	pc := usecase.NewPermissionUsecase(pr, timeout)
	lc := &controller.RoleController{
		RoleUseCase: usecase.NewRoleUsecase(ur, rr, pc, timeout),
		Env:         env,
	}
	group.POST("/roles", lc.CreateRole)
	group.DELETE("/roles/:roleId", lc.DeleteRole)
	group.GET("/roles/:roleId", lc.GetRoleByID)
	group.PUT("/roles/:roleId", lc.UpdateRole)
	group.GET("/roles", lc.ListRoles)
}
