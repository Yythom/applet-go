package middleware

import (
	"applet/domain"
	"github.com/gin-gonic/gin"
	"net/http"
)

// PermissionMiddleware 权限检查中间件
func PermissionMiddleware(requiredRoles []string) gin.HandlerFunc {
	return func(c *gin.Context) {
		if !checkPermission(c, requiredRoles) {
			c.JSON(http.StatusForbidden, domain.Response{Message: "Forbidden", Code: "FORBIDDEN"})
			c.Abort()
			return
		}
		c.Next()
	}
}

// checkPermission 根据用户角色检查是否有权限访问路由
func checkPermission(c *gin.Context, requiredRoles []string) bool {
	userRoles, exists := c.Get("userRoles")
	if !exists {
		// 如果无法获取用户角色信息，拒绝访问
		return false
	}

	// 转换为字符串切片
	roles, ok := userRoles.([]string)
	if !ok {
		// 如果无法转换为字符串切片，拒绝访问
		return false
	}

	// 检查用户角色是否包含所需的角色
	for _, requiredRole := range requiredRoles {
		found := false
		for _, userRole := range roles {
			if userRole == requiredRole {
				found = true
				break
			}
		}
		if !found {
			// 如果用户角色不包含所需角色，拒绝访问
			return false
		}
	}

	// 用户角色包含所需角色，允许访问
	return true
}
