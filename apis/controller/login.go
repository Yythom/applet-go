package controller

import (
	"applet/core/bootstrap"
	"applet/domain"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
)

type LoginController struct {
	LoginUsecase domain.LoginUsecase
	Env          *bootstrap.Env
}

// Login @Summary User Login
// @Description Logs in a user and returns access and refresh tokens
// @Tags auth
// @Accept json
// @Produce json
// @Param name formData string true "Username"
// @Param password formData string true "Password"
// @Success 200 {object} domain.Response{data=domain.LoginResponse} "Successful login"
// @Failure 401 {object} domain.Response "Invalid credentials"
// @Failure 404 {object} domain.Response "User not found with the given name"
// @Failure 500 {object} domain.Response "Internal server error"
// @Router /login [post]
func (lc *LoginController) Login(c *gin.Context) {
	var request domain.LoginRequest

	err := c.ShouldBind(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.Response{Message: err.Error()})
		return
	}

	log.Print(request.Name)
	user, err := lc.LoginUsecase.GetUserByName(c, request.Name)
	if err != nil {
		c.JSON(http.StatusNotFound, domain.Response{Message: "User not found with the given name"})
		return
	}

	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password)) != nil {
		c.JSON(http.StatusUnauthorized, domain.Response{Message: "Invalid credentials"})
		return
	}

	accessToken, err := lc.LoginUsecase.CreateAccessToken(&user, lc.Env.AccessTokenSecret, lc.Env.AccessTokenExpiryHour)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.Response{Message: err.Error()})
		return
	}

	refreshToken, err := lc.LoginUsecase.CreateRefreshToken(&user, lc.Env.RefreshTokenSecret, lc.Env.RefreshTokenExpiryHour)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.Response{Message: err.Error()})
		return
	}

	data := domain.LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	loginResponse := domain.Response{
		Message: "login success",
		Data:    data,
		Code:    "SUCCESS",
	}

	c.JSON(http.StatusOK, loginResponse)
}
