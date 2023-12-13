package usecase

import (
	"applet/core/jwt"
	"applet/domain"
	"context"
	"time"
)

type loginUsecase struct {
	userRepository domain.UserRepository
	contextTimeout time.Duration
}

func NewLoginUsecase(userRepository domain.UserRepository, timeout time.Duration) domain.LoginUsecase {
	return &loginUsecase{
		userRepository: userRepository,
		contextTimeout: timeout,
	}
}

func (lu *loginUsecase) GetUserByName(c context.Context, name string) (domain.User, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.userRepository.GetByName(ctx, name)
}

func (lu *loginUsecase) CreateAccessToken(user *domain.User, secret string, expiry int) (accessToken string, err error) {
	return jwt.CreateAccessToken(user, secret, expiry)
}

func (lu *loginUsecase) CreateRefreshToken(user *domain.User, secret string, expiry int) (refreshToken string, err error) {
	return jwt.CreateRefreshToken(user, secret, expiry)
}
