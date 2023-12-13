package middleware

import (
	"applet/core/jwt"
	"applet/domain"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func JwtAuthMiddleware(secret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		t := strings.Split(authHeader, " ")
		if len(t) == 2 {
			authToken := t[1]
			authorized, err := jwt.IsAuthorized(authToken, secret)
			if authorized {
				userID, err := jwt.ExtractIDFromToken(authToken, secret)
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
