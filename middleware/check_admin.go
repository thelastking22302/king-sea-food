package middleware

import (
	"net/http"
	"thelastking/kingseafood/model"

	"github.com/gin-gonic/gin"
)

func IsAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		var data model.Users
		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error":   err.Error(),
				"command": "middleware error",
			})
			return
		}
		if data.Email != "admin@gmail.com" {
			c.JSON(http.StatusBadRequest, gin.H{
				"command": "you not api",
			})
			return
		}
		c.Next()
	}
}
