package domain

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Role struct {
	Id            primitive.ObjectID `bson:"_id"`
	Name          string             `bson:"role_name"`
	Description   string             `bson:"role_description"`
	Status        string             `bson:"role_status"`
	PermissionIds []string           `bson:"permission_ids"`
}

type RoleWithPermission struct {
	Role
	Permissions []Permission `bson:"permissions"`
}

type RoleRepository interface {
	Create(c context.Context, role Role) error
	Fetch(c context.Context) ([]Role, error)
	GetByID(c context.Context, id string) (Role, error)
	Update(c context.Context, id string, role Role) (Role, error)
	DeleteByID(ctx context.Context, id string) error
}

type RoleUseCase interface {
	Create(c context.Context, role Role) error
	Fetch(c context.Context) ([]RoleWithPermission, error)
	GetByID(c context.Context, id string) (RoleWithPermission, error)
	Update(c context.Context, id string, role Role) (RoleWithPermission, error)
	Delete(c context.Context, id string) error
	AddPermissionIds(c context.Context, roleId string, permissionIds []string) error
	RemovePermissionIds(c context.Context, roleId string, permissionIds []string) error
	AssignRolesToUser(c context.Context, userId string, roleId string) error
	RemoveRolesFromUser(c context.Context, userId string, roleId string) error
}
