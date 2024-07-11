package middleware

import (
	"net/http"
	"strings"
	"thelastking/kingseafood/pkg/security"

	"github.com/gin-gonic/gin"
)

func JwtMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		claims, err := security.ValidateToken(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}

		// Lưu trữ userId vào context để handler sử dụng
		c.Set("user_id", claims.UserID)
		c.Next()
	}
}
