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
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, domain.Response{Message: "Not authorized"})
			c.Abort()
			return
		}

		userID, err := extractUserID(authHeader, secret)
		if err != nil {
			c.JSON(http.StatusUnauthorized, domain.Response{Message: "Not found userID is" + userID})
			c.Abort()
			return
		}

		c.Set("x-user-id", userID)
		c.Next()
	}
}

func extractUserID(authHeader string, secret string) (string, error) {
	authorized, err := jwt.IsAuthorized(authHeader, secret)
	if err != nil {
		return "", err
	}
	if !authorized {
		return "", err
	}

	userID, err := jwt.ExtractIDFromToken(authHeader, secret)
	if err != nil {
		return "", err
	}

	return userID, nil
}
