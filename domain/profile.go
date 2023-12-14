package domain

import "context"

type ProfileRequest struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type ProfileUsecase interface {
	GetProfileByID(c context.Context, userID string) (*ProfileRequest, error)
}
