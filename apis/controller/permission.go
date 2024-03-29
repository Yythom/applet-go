package controller

import (
	"applet/core/bootstrap"
	"applet/domain"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type PermissionController struct {
	PermissionUseCase domain.PermissionUseCase
	Env               *bootstrap.Env
}

// CreatePermission @Summary 创建权限
// @Description 创建权限
// @Tags permissions
// @Accept json
// @Produce json
// @Param request body domain.CreatePermissionRequest true "权限信息"
// @Success 200 {object} domain.Response "成功"
// @Failure 400 {object} domain.Response "请求错误"
// @Failure 404 {object} domain.Response "未找到资源"
// @Failure 409 {object} domain.Response "资源冲突"
// @Failure 500 {object} domain.Response "内部服务器错误"
// @Router /permissions [post]
func (pc *PermissionController) CreatePermission(c *gin.Context) {
	var request domain.CreatePermissionRequest

	err := c.ShouldBind(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.Response{Message: err.Error()})
		return
	}

	log.Print(request.Name)

	// 检查改权限是否已经存在
	existingPermission, err := pc.PermissionUseCase.GetByName(c, request.Name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.Response{Message: "Internal Server Error", Code: "ERROR"})
		return
	}
	if existingPermission != nil {
		c.JSON(http.StatusConflict, domain.Response{Message: "Permission already exists", Code: "CONFLICT"})
		return
	}

	// 创建该权限
	err = pc.PermissionUseCase.Create(c, request)
	if err != nil {
		return
	}

	permissionResponse := domain.Response{
		Message: "login success",
		Data:    domain.CreatePermissionResponse{},
		Code:    "SUCCESS",
	}

	c.JSON(http.StatusOK, permissionResponse)
}

// UpdatePermission @Summary 修改权限
// @Description 修改权限
// @Tags permissions
// @Accept json
// @Produce json
// @Param permissionId path string true "权限ID"
// @Param request body domain.UpdatePermissionRequest true "权限信息"
// @Success 200 {object} domain.Response "成功"
// @Failure 400 {object} domain.Response "请求错误"
// @Failure 404 {object} domain.Response "未找到资源"
// @Failure 500 {object} domain.Response "内部服务器错误"
// @Router /permissions/{permissionId} [put]
func (pc *PermissionController) UpdatePermission(c *gin.Context) {
	var request domain.UpdatePermissionRequest

	permissionId := c.Param("permissionId")

	err := c.ShouldBind(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.Response{Message: err.Error()})
		return
	}

	// 根据权限ID查询权限
	permission, err := pc.PermissionUseCase.GetByID(c, permissionId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.Response{Message: "Internal Server Error", Code: "ERROR"})
		return
	}
	if permission == nil {
		c.JSON(http.StatusNotFound, domain.Response{Message: "Permission not found", Code: "NOT_FOUND"})
		return
	}

	// 更新权限
	permission, err = pc.PermissionUseCase.Update(c, request)
	if err != nil {
		// 处理错误
		return
	}

	c.JSON(http.StatusOK, domain.Response{Message: "Permission updated successfully", Code: "SUCCESS", Data: permission})
}

// DeletePermission @Summary 删除权限
// @Description 删除权限
// @Tags permissions
// @Accept json
// @Produce json
// @Param permissionId path string true "权限ID"
// @Success 200 {object} domain.Response "成功"
// @Failure 404 {object} domain.Response "未找到资源"
// @Failure 500 {object} domain.Response "内部服务器错误"
// @Router /permissions/{permissionId} [delete]
func (pc *PermissionController) DeletePermission(c *gin.Context) {
	permissionID := c.Param("permissionId")

	// 根据权限ID查询权限
	permission, err := pc.PermissionUseCase.GetByID(c, permissionID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.Response{Message: "Internal Server Error", Code: "ERROR"})
		return
	}
	if permission == nil {
		c.JSON(http.StatusNotFound, domain.Response{Message: "Permission not found", Code: "NOT_FOUND"})
		return
	}

	// 删除权限
	err = pc.PermissionUseCase.Delete(c, permissionID)
	if err != nil {
		// 处理错误
		return
	}

	c.JSON(http.StatusOK, domain.Response{Message: "Permission deleted successfully", Code: "SUCCESS"})
}

// GetPermission @Summary 查询权限
// @Description 查询权限
// @Tags permissions
// @Accept json
// @Produce json
// @Param permissionId path string true "权限ID"
// @Success 200 {object} domain.Response "成功"
// @Failure 404 {object} domain.Response "未找到资源"
// @Failure 500 {object} domain.Response "内部服务器错误"
// @Router /permissions/{permissionId} [get]
func (pc *PermissionController) GetPermission(c *gin.Context) {
	permissionID := c.Param("permissionId")

	// 根据权限ID查询权限
	permission, err := pc.PermissionUseCase.GetByID(c, permissionID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.Response{Message: "Internal Server Error", Code: "ERROR"})
		return
	}
	if permission == nil {
		c.JSON(http.StatusNotFound, domain.Response{Message: "Permission not found", Code: "NOT_FOUND"})
		return
	}

	c.JSON(http.StatusOK, domain.Response{Message: "Permission found", Data: permission, Code: "SUCCESS"})
}

// GetPermissionList @Summary 查询权限列表
// @Description 查询权限列表
// @Tags permissions
// @Accept json
// @Produce json
// @Success 200 {object} domain.Response "成功"
// @Failure 500 {object} domain.Response "内部服务器错误"
// @Router /permissions [get]
func (pc *PermissionController) GetPermissionList(c *gin.Context) {
	// 查询权限列表
	permissions, err := pc.PermissionUseCase.GetList(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.Response{Message: "Internal Server Error", Code: "ERROR"})
		return
	}

	c.JSON(http.StatusOK, domain.Response{Message: "Permission list found", Data: permissions, Code: "SUCCESS"})
}
