package middleware

import (
	"applet/core/jwt"
	"applet/domain"
	"github.com/gin-gonic/gin"
	"net/http"
)

func JwtAuthMiddleware(secret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader != "" {
			authorized, err := jwt.IsAuthorized(authHeader, secret)
			if authorized {
				userID, err := jwt.ExtractIDFromToken(authHeader, secret)
				if err != nil {
					c.JSON(http.StatusUnauthorized, domain.Response{Message: err.Error()})
					c.Abort()
					return
				}
				c.Set("x-user-id", userID)
				c.Next()
				return
			}
			c.JSON(http.StatusUnauthorized, domain.Response{Message: err.Error()})
			c.Abort()
			return
		}
		c.JSON(http.StatusUnauthorized, domain.Response{Message: "Not authorized"})
		c.Abort()
	}
}
