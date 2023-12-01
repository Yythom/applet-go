package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"test/controlles/user"
)

func setupRouter() *gin.Engine {
	// Disable Console Color
	// gin.DisableConsoleColor()
	r := gin.Default()

	defer func() {
		if err := recover(); err != nil {
			fmt.Println("error:", err)
		}
	}()

	userGroup := r.Group("/user")
	{
		userGroup.GET("/info", user.GetUserInfo)
		userGroup.POST("/register", user.Register)
		userGroup.POST("/bind", user.BindMobile)
		userGroup.POST("/login", user.Login)
	}

	defer fmt.Println(1)

	return r
}

func main() {
	r := setupRouter()
	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}
