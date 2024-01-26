package domain

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Permission struct {
	Id          primitive.ObjectID `bson:"_id"`
	Name        string             `bson:"permission_name"`
	Description string             `bson:"permission_description"`
	Status      string             `bson:"permission_status"`
}

// CreatePermissionRequest 创建权限请求
type CreatePermissionRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type CreatePermissionForRepository struct {
	CreatePermissionRequest
	Status string `json:"status"`
}

// CreatePermissionResponse 创建权限响应
type CreatePermissionResponse struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Status      string `json:"status"`
}

// UpdatePermissionRequest 更新权限请求
type UpdatePermissionRequest struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

// UpdatePermissionResponse 更新权限响应
type UpdatePermissionResponse struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

// GetPermissionRequest 获取权限请求
type GetPermissionRequest struct {
	ID string `json:"id"`
}

// GetPermissionResponse 获取权限响应
type GetPermissionResponse struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// DeletePermissionRequest 删除权限请求
type DeletePermissionRequest struct {
	ID string `json:"id"`
}

// DeletePermissionResponse 删除权限响应
type DeletePermissionResponse struct {
	ID string `json:"id"`
}

type PermissionRepository interface {
	Create(c context.Context, permission CreatePermissionForRepository) error
	Fetch(c context.Context) ([]Permission, error)
	GetByID(c context.Context, id string) (*Permission, error)
	GetByIDs(c context.Context, ids []string) ([]Permission, error)
	Update(c context.Context, permission UpdatePermissionRequest) (*Permission, error)
	DeleteByID(c context.Context, id string) error
	GetByName(c context.Context, name string) (*Permission, error)
}

type PermissionUseCase interface {
	Create(c context.Context, permission CreatePermissionRequest) error
	Fetch(c context.Context) ([]Permission, error)
	GetByID(c context.Context, id string) (*Permission, error)
	GetByIDs(c context.Context, ids []string) ([]Permission, error)
	GetByName(c context.Context, name string) (*Permission, error)
	GetList(c context.Context) ([]Permission, error)
	Update(c context.Context, permission UpdatePermissionRequest) (*Permission, error)
	Delete(c context.Context, id string) error
}
