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
	rr := repository.NewRoleRepository(db, domain.CollectionUser)
	ur := repository.NewUserRepository(db, domain.CollectionUser)
	pr := repository.NewPermissionRepository(db, domain.CollectionUser)
	pc := usecase.NewPermissionUsecase(pr, timeout)
	lc := &controller.RoleController{
		RoleUseCase: usecase.NewRoleUsecase(ur, rr, pc, timeout),
		Env:         env,
	}
	group.POST("/role/create", lc.CreateRole)
	group.DELETE("/role/delete", lc.DeleteRole)
	group.GET("/role", lc.GetRoleByID)
	group.GET("/role/list", lc.ListRoles)
}
