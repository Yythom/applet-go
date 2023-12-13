package controlles

import (
	"applet/core/bootstrap"
	"applet/domain"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

type RegisterController struct {
	RegisterUsecase domain.RegisterUsecase
	Env             *bootstrap.Env
}

func (sc *RegisterController) Register(c *gin.Context) {
	var request domain.RegisterRequest

	err := c.ShouldBind(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.Response{Message: err.Error()})
		return
	}

	_, err = sc.RegisterUsecase.GetUserByName(c, request.Name)
	if err == nil {
		c.JSON(http.StatusConflict, domain.Response{Message: "User already exists with the given email"})
		return
	}

	encryptedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(request.Password),
		bcrypt.DefaultCost,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.Response{Message: err.Error()})
		return
	}

	request.Password = string(encryptedPassword)

	user := domain.User{
		ID:       primitive.NewObjectID(),
		Name:     request.Name,
		Email:    request.Email,
		Password: request.Password,
	}

	err = sc.RegisterUsecase.Create(c, &user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.Response{Message: err.Error()})
		return
	}

	accessToken, err := sc.RegisterUsecase.CreateAccessToken(&user, sc.Env.AccessTokenSecret, sc.Env.AccessTokenExpiryHour)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.Response{Message: err.Error()})
		return
	}

	refreshToken, err := sc.RegisterUsecase.CreateRefreshToken(&user, sc.Env.RefreshTokenSecret, sc.Env.RefreshTokenExpiryHour)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.Response{Message: err.Error()})
		return
	}

	data := domain.RegisterResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	signupResponse := domain.Response{
		Message: "register success",
		Data:    data,
		Code:    "SUCCESS",
	}

	c.JSON(http.StatusOK, signupResponse)
}
