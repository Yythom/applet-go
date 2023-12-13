package usecase

import (
	"applet/core/jwt"
	"applet/domain"
	"context"
	"time"
)

type registerUsecase struct {
	userRepository domain.UserRepository
	contextTimeout time.Duration
}

func NewRegisterUsecase(userRepository domain.UserRepository, timeout time.Duration) domain.RegisterUsecase {
	return &registerUsecase{
		userRepository: userRepository,
		contextTimeout: timeout,
	}
}

func (su *registerUsecase) Create(c context.Context, user *domain.User) error {
	ctx, cancel := context.WithTimeout(c, su.contextTimeout)
	defer cancel()
	return su.userRepository.Create(ctx, user)
}

func (su *registerUsecase) GetUserByName(c context.Context, name string) (domain.User, error) {
	ctx, cancel := context.WithTimeout(c, su.contextTimeout)
	defer cancel()
	return su.userRepository.GetByName(ctx, name)
}

func (su *registerUsecase) CreateAccessToken(user *domain.User, secret string, expiry int) (accessToken string, err error) {
	return jwt.CreateAccessToken(user, secret, expiry)
}

func (su *registerUsecase) CreateRefreshToken(user *domain.User, secret string, expiry int) (refreshToken string, err error) {
	return jwt.CreateRefreshToken(user, secret, expiry)
}
