package domain

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	CollectionUser       = "users"
	CollectionRole       = "roles"
	CollectionPermission = "permissions"
)

type User struct {
	ID       primitive.ObjectID `bson:"_id"`
	Name     string             `bson:"name"`
	Email    string             `bson:"email"`
	Password string             `bson:"password"`
	RoleIds  []string           `bson:"role_ids"`
}

type UserRepository interface {
	Create(c context.Context, user *User) error
	Fetch(c context.Context) ([]User, error)
	GetByName(c context.Context, name string) (User, error)
	GetByID(c context.Context, id string) (User, error)
	Update(c context.Context, id string, user *User) (User, error)
}
