package usecase

import (
	"applet/domain"
	"context"
	"errors"
	"time"
)

type roleUsecase struct {
	roleRepository    domain.RoleRepository
	userRepository    domain.UserRepository
	permissionUsecase domain.PermissionUseCase
	contextTimeout    time.Duration
}

func (r roleUsecase) AssignRolesToUser(c context.Context, userId string, roleId string) error {
	ctx, cancel := context.WithTimeout(c, r.contextTimeout)
	defer cancel()

	// 查询用户和角色
	user, err := r.userRepository.GetByID(ctx, userId)
	if err != nil {
		return err
	}

	role, err := r.roleRepository.GetByID(ctx, roleId)
	if err != nil {
		return err
	}

	// 检查角色是否已分配给用户
	for _, assignedRoleID := range user.RoleIds {
		if assignedRoleID == role.Id.Hex() {
			return errors.New("role already assigned to user")
		}
	}

	// 将角色ID添加到用户的RoleIds中
	user.RoleIds = append(user.RoleIds, roleId)

	// 更新用户信息
	_, err = r.userRepository.Update(ctx, userId, &user)
	if err != nil {
		return err
	}

	return nil
}

func (r roleUsecase) RemoveRolesFromUser(c context.Context, userId string, roleId string) error {
	ctx, cancel := context.WithTimeout(c, r.contextTimeout)
	defer cancel()

	// 查询用户和角色
	user, err := r.userRepository.GetByID(ctx, userId)
	if err != nil {
		return err
	}

	role, err := r.roleRepository.GetByID(ctx, roleId)
	if err != nil {
		return err
	}

	// 检查角色是否已分配给用户
	roleAssigned := false
	for i, assignedRoleID := range user.RoleIds {
		if assignedRoleID == role.Id.Hex() {
			// 从用户的RoleIds中移除角色ID
			user.RoleIds = append(user.RoleIds[:i], user.RoleIds[i+1:]...)
			roleAssigned = true
			break
		}
	}

	// 如果角色未分配给用户，则返回错误
	if !roleAssigned {
		return errors.New("role not assigned to user")
	}

	// 更新用户信息
	_, err = r.userRepository.Update(ctx, userId, &user)
	if err != nil {
		return err
	}

	return nil
}

func (r roleUsecase) Create(c context.Context, role domain.Role) error {
	ctx, cancel := context.WithTimeout(c, r.contextTimeout)
	defer cancel()
	return r.roleRepository.Create(ctx, role)
}

func (r roleUsecase) Fetch(c context.Context) ([]domain.RoleWithPermission, error) {
	ctx, cancel := context.WithTimeout(c, r.contextTimeout)
	defer cancel()

	// 查询所有角色
	roles, err := r.roleRepository.Fetch(ctx)
	if err != nil {
		return nil, err
	}

	// 保存角色和对应的权限
	rolesWithPermissions := make([]domain.RoleWithPermission, 0)

	// 遍历每个角色
	for _, role := range roles {
		// 获取角色的PermissionIds
		permissionIds := role.PermissionIds

		// 查询对应的Permissions
		permissions, err := r.permissionUsecase.GetByIDs(ctx, permissionIds)
		if err != nil {
			// 处理错误
			return nil, err
		}

		// 构建RoleWithPermission结构
		roleWithPermission := domain.RoleWithPermission{
			Role:        role,
			Permissions: permissions,
		}

		// 添加到结果切片
		rolesWithPermissions = append(rolesWithPermissions, roleWithPermission)
	}

	return rolesWithPermissions, nil
}

func (r roleUsecase) GetByID(c context.Context, id string) (domain.RoleWithPermission, error) {
	ctx, cancel := context.WithTimeout(c, r.contextTimeout)
	defer cancel()

	role, err := r.roleRepository.GetByID(ctx, id)
	if err != nil {
		return domain.RoleWithPermission{}, err
	}

	// 查询对应的Permissions
	permissions, err := r.permissionUsecase.GetByIDs(ctx, role.PermissionIds)
	if err != nil {
		// 处理错误
		return domain.RoleWithPermission{}, err
	}

	// 构建RoleWithPermission结构
	roleWithPermission := domain.RoleWithPermission{
		Role:        role,
		Permissions: permissions,
	}

	return roleWithPermission, nil

}

func (r roleUsecase) Update(c context.Context, id string, role domain.Role) (domain.RoleWithPermission, error) {
	ctx, cancel := context.WithTimeout(c, r.contextTimeout)
	defer cancel()

	// 先查询要更新的角色
	oldRole, err := r.roleRepository.GetByID(ctx, id)
	if err != nil {
		return domain.RoleWithPermission{}, err
	}

	// 更新角色的字段
	oldRole.Name = role.Name
	oldRole.Description = role.Description
	oldRole.Status = role.Status

	// 保存更新后的角色
	updatedRole, err := r.roleRepository.Update(ctx, id, oldRole)
	if err != nil {
		return domain.RoleWithPermission{}, err
	}

	// 获取更新后的角色的权限
	permissions, err := r.permissionUsecase.GetByIDs(ctx, updatedRole.PermissionIds)
	if err != nil {
		return domain.RoleWithPermission{}, err
	}

	// 构建RoleWithPermission结构
	roleWithPermission := domain.RoleWithPermission{
		Role:        updatedRole,
		Permissions: permissions,
	}

	return roleWithPermission, nil
}

func (r roleUsecase) AddPermissionIds(c context.Context, roleId string, permissionIds []string) error {
	ctx, cancel := context.WithTimeout(c, r.contextTimeout)
	defer cancel()

	// 先查询要添加权限的角色
	role, err := r.roleRepository.GetByID(ctx, roleId)
	if err != nil {
		return err
	}

	// 合并新的权限ID到角色的PermissionIds中
	role.PermissionIds = append(role.PermissionIds, permissionIds...)

	// 更新角色的权限
	_, err = r.roleRepository.Update(ctx, roleId, role)
	if err != nil {
		return err
	}

	return nil
}

func removePermissionIDs(source []string, toRemove []string) []string {
	result := make([]string, 0, len(source))
	// 创建一个 map 以便检查元素是否需要移除
	toRemoveMap := make(map[string]struct{})
	for _, id := range toRemove {
		toRemoveMap[id] = struct{}{}
	}
	// 遍历原始切片，仅将不在 toRemoveMap 中的元素添加到结果切片
	for _, id := range source {
		if _, exists := toRemoveMap[id]; !exists {
			result = append(result, id)
		}
	}
	return result
}

func (r roleUsecase) RemovePermissionIds(c context.Context, roleId string, permissionIds []string) error {
	ctx, cancel := context.WithTimeout(c, r.contextTimeout)
	defer cancel()

	// 先查询要移除权限的角色
	role, err := r.roleRepository.GetByID(ctx, roleId)
	if err != nil {
		return err
	}

	// 从角色的PermissionIds中移除指定的权限ID
	role.PermissionIds = removePermissionIDs(role.PermissionIds, permissionIds)

	// 更新角色的权限
	_, err = r.roleRepository.Update(ctx, roleId, role)
	if err != nil {
		return err
	}

	return nil
}

func (r roleUsecase) Delete(c context.Context, id string) error {
	ctx, cancel := context.WithTimeout(c, r.contextTimeout)
	defer cancel()

	err := r.roleRepository.DeleteByID(ctx, id)
	if err != nil {
		return err
	}

	return nil
}

func NewRoleUsecase(userRepository domain.UserRepository, roleRepository domain.RoleRepository, permissionUsecase domain.PermissionUseCase, timeout time.Duration) domain.RoleUseCase {
	return &roleUsecase{
		userRepository:    userRepository,
		roleRepository:    roleRepository,
		permissionUsecase: permissionUsecase,
		contextTimeout:    timeout,
	}
}
