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

// CreateRole @Summary 创建角色
// @Description 创建角色
// @Tags roles
// @Accept json
// @Produce json
// @Param role body domain.Role true "角色信息"
// @Success 201 {object} domain.Response "角色创建成功"
// @Failure 400 {object} domain.Response "请求错误"
// @Failure 500 {object} domain.Response "内部服务器错误"
// @Router /roles [post]
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

// DeleteRole @Summary 删除角色
// @Description 删除角色
// @Tags roles
// @Accept json
// @Produce json
// @Param roleId path string true "角色ID"
// @Param permissionIds query []string false "权限ID列表"
// @Success 200 {object} domain.Response "角色删除成功"
// @Failure 500 {object} domain.Response "内部服务器错误"
// @Router /roles/{roleId} [delete]
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

// UpdateRole @Summary 更新角色
// @Description 更新角色
// @Tags roles
// @Accept json
// @Produce json
// @Param roleId path string true "角色ID"
// @Param updatedRole body domain.Role true "更新后的角色信息"
// @Success 200 {object} domain.Response{data=domain.Role} "角色更新成功"
// @Failure 400 {object} domain.Response "请求错误"
// @Failure 500 {object} domain.Response "内部服务器错误"
// @Router /roles/{roleId} [put]
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

// GetRoleByID @Summary 根据ID获取角色信息
// @Description 根据ID获取角色信息
// @Tags roles
// @Accept json
// @Produce json
// @Param roleId path string true "角色ID"
// @Success 200 {object} domain.Response{data=domain.Role} "获取角色信息成功"
// @Failure 500 {object} domain.Response "内部服务器错误"
// @Router /roles/{roleId} [get]
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

// ListRoles @Summary 获取角色列表
// @Description 获取角色列表
// @Tags roles
// @Accept json
// @Produce json
// @Success 200 {object} domain.Response{data=[]domain.Role} "获取角色列表成功"
// @Failure 500 {object} domain.Response "内部服务器错误"
// @Router /roles [get]
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
