package model

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserInfo struct {
	UserId           int    `json:"user_id"`
	Username         string `json:"username"`
	Password         string `json:"password"`
	Gender           string `json:"gender"`
	Birthday         string `json:"birthday"`
	Address          string `json:"address"`
	LastLogin        string `json:"last_login"`
	AccountLocked    bool   `json:"account_locked"`
	RegistrationTime string `json:"registration_time"`
	UserType         string `json:"user_type"`
}

type User struct {
	ID       primitive.ObjectID `bson:"_id"`
	Name     string             `bson:"name"`
	Email    string             `bson:"email"`
	Password string             `bson:"password"`
}

type UserRepository interface {
	Create(c context.Context, user *User) error
	Fetch(c context.Context) ([]User, error)
	GetByEmail(c context.Context, email string) (User, error)
	GetByID(c context.Context, id string) (User, error)
}
