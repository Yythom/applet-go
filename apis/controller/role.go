package controller

import (
	"applet/core/bootstrap"
	"applet/domain"
	"github.com/gin-gonic/gin"
	"net/http"
)

type RoleController struct {
	RoleUseCase domain.RoleUseCase
	Env         *bootstrap.Env
}

// CreateRole 创建角色
func (rc *RoleController) CreateRole(c *gin.Context) {
	var role domain.Role
	if err := c.ShouldBindJSON(&role); err != nil {
		c.JSON(http.StatusBadRequest, domain.Response{Message: err.Error()})
		return
	}

	if err := rc.RoleUseCase.Create(c, role); err != nil {
		c.JSON(http.StatusInternalServerError, domain.Response{Message: err.Error()})
		return
	}

	c.JSON(http.StatusCreated, domain.Response{Message: "Role created successfully"})
}

// DeleteRole 删除角色
func (rc *RoleController) DeleteRole(c *gin.Context) {
	//  从 URL 参数或请求体中获取角色 ID
	roleId := c.Param("roleId")
	permissionIds := c.QueryArray("permissionIds")

	userId := c.GetString("x-user-id")

	if err := rc.RoleUseCase.RemovePermissionIds(c, roleId, permissionIds); err != nil {
		c.JSON(http.StatusInternalServerError, domain.Response{Message: err.Error()})
		return
	}

	if err := rc.RoleUseCase.RemoveRolesFromUser(c, userId, roleId); err != nil {
		c.JSON(http.StatusInternalServerError, domain.Response{Message: err.Error()})
		return
	}

	if err := rc.RoleUseCase.Delete(c, roleId); err != nil {
		c.JSON(http.StatusInternalServerError, domain.Response{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, domain.Response{Message: "Role deleted successfully"})
}

// UpdateRole 更新角色
func (rc *RoleController) UpdateRole(c *gin.Context) {
	//  从 URL 参数或请求体中获取角色 ID
	roleId := c.Param("roleId")

	var updatedRole domain.Role
	if err := c.ShouldBindJSON(&updatedRole); err != nil {
		c.JSON(http.StatusBadRequest, domain.Response{Message: err.Error()})
		return
	}

	roles, err := rc.RoleUseCase.Update(c, roleId, updatedRole)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.Response{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, domain.Response{Message: "Role updated successfully", Data: roles})
}

// GetRoleByID 根据ID获取角色信息
func (rc *RoleController) GetRoleByID(c *gin.Context) {
	//  从 URL 参数中获取角色 ID
	roleId := c.Param("roleId")

	//  调用 usecase 方法获取角色信息
	role, err := rc.RoleUseCase.GetByID(c, roleId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.Response{Message: err.Error()})
		return
	}

	//  返回角色信息
	c.JSON(http.StatusOK, domain.Response{Message: "SUCCESS", Data: role})
}

// ListRoles 获取角色列表
func (rc *RoleController) ListRoles(c *gin.Context) {
	//  调用 usecase 方法获取角色列表
	roles, err := rc.RoleUseCase.Fetch(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	//  返回角色列表
	c.JSON(http.StatusOK, domain.Response{Message: "SUCCESS", Data: roles})
}
