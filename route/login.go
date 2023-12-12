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

func NewLoginRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewUserRepository(db, domain.CollectionUser)
	lc := &controlles.LoginController{
		LoginUsecase: usecase.NewLoginUsecase(ur, timeout),
		Env:          env,
	}
	group.POST("/login", lc.Login)
}