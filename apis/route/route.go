package route

import (
	"applet/apis/middleware"
	"applet/core/bootstrap"
	"applet/core/mongo"
	"github.com/gin-gonic/gin"
	"time"
)

func Setup(env *bootstrap.Env, timeout time.Duration, db mongo.Database, gin *gin.Engine) {
	publicRouter := gin.Group("")
	// All Public APIs
	NewRegisterRouter(env, timeout, db, publicRouter)
	NewLoginRouter(env, timeout, db, publicRouter)
	NewRefreshTokenRouter(env, timeout, db, publicRouter)

	protectedRouter := gin.Group("")
	// Middleware to verify AccessToken
	protectedRouter.Use(middleware.JwtAuthMiddleware(env.AccessTokenSecret))
	// All Private APIs
	//NewProfileRouter(env, timeout, db, protectedRouter)
	//NewTaskRouter(env, timeout, db, protectedRouter)
}
