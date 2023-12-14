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

func NewProfileRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewUserRepository(db, domain.CollectionUser)
	pc := &controller.ProfileController{
		ProfileUsecase: usecase.NewProfileUsecase(ur, timeout),
	}
	group.GET("/profile", pc.Fetch)
}
