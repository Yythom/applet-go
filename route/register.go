package route

import (
	"applet/controlles"
	"applet/core/bootstrap"
	"applet/core/mongo"
	"applet/domain"
	"applet/repository"
	"applet/usecase"
	"github.com/gin-gonic/gin"
	"time"
)

func NewRegisterRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewUserRepository(db, domain.CollectionUser)
	sc := controlles.RegisterController{
		RegisterUsecase: usecase.NewRegisterUsecase(ur, timeout),
		Env:             env,
	}
	group.POST("/register", sc.Register)
}
