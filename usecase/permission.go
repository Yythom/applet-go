package usecase

import (
	"applet/domain"
	"context"
	"fmt"
	"time"
)

type permissionUsecase struct {
	permissionRepository domain.PermissionRepository
	contextTimeout       time.Duration
}

func (p permissionUsecase) GetList(c context.Context) ([]domain.Permission, error) {
	// 获取权限列表的业务逻辑
	permissions, err := p.permissionRepository.Fetch(c)
	if err != nil {
		return nil, err
	}

	return permissions, nil
}

func (p permissionUsecase) Delete(c context.Context, id string) error {
	// 调用 Repository 的删除方法
	err := p.permissionRepository.DeleteByID(c, id)
	if err != nil {
		return err
	}

	return nil
}

func (p permissionUsecase) Create(c context.Context, permission domain.CreatePermissionRequest) error {
	ctx, cancel := context.WithTimeout(c, p.contextTimeout)
	defer cancel()

	// 设置默认状态
	Status := "init"

	// 创建包含状态的请求
	withStatusRequest := domain.CreatePermissionForRepository{
		CreatePermissionRequest: permission,
		Status:                  Status,
	}

	return p.permissionRepository.Create(ctx, withStatusRequest)
}

func (p permissionUsecase) GetByName(c context.Context, name string) (*domain.Permission, error) {
	ctx, cancel := context.WithTimeout(c, p.contextTimeout)
	defer cancel()

	// 调用 Repository 的方法获取权限
	permission, err := p.permissionRepository.GetByName(ctx, name)
	if err != nil {
		// 根据实际情况处理错误
		return nil, fmt.Errorf("failed to get permission by name: %w", err)
	}

	return permission, nil
}

func (p permissionUsecase) Fetch(c context.Context) ([]domain.Permission, error) {
	ctx, cancel := context.WithTimeout(c, p.contextTimeout)
	defer cancel()
	return p.permissionRepository.Fetch(ctx)
}

func (p permissionUsecase) GetByID(c context.Context, id string) (*domain.Permission, error) {
	ctx, cancel := context.WithTimeout(c, p.contextTimeout)
	defer cancel()
	return p.permissionRepository.GetByID(ctx, id)
}

func (p permissionUsecase) GetByIDs(c context.Context, ids []string) ([]domain.Permission, error) {
	ctx, cancel := context.WithTimeout(c, p.contextTimeout)
	defer cancel()

	permissions, err := p.permissionRepository.GetByIDs(ctx, ids)
	if err != nil {
		// 处理错误
		return nil, err
	}

	return permissions, nil
}

func (p permissionUsecase) Update(c context.Context, permission domain.UpdatePermissionRequest) (*domain.Permission, error) {
	ctx, cancel := context.WithTimeout(c, p.contextTimeout)
	defer cancel()
	return p.permissionRepository.Update(ctx, permission)
}

func NewPermissionUsecase(permissionRepository domain.PermissionRepository, timeout time.Duration) domain.PermissionUseCase {
	return &permissionUsecase{
		permissionRepository: permissionRepository,
		contextTimeout:       timeout,
	}
}
