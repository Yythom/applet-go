package route

import (
	"github.com/gin-gonic/gin"
	"test/bootstrap"
	"test/controlles"
	"test/domain"
	"test/mongo"
	"test/repository"
	"test/usecase"
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
