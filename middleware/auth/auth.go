package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mrokoo/goERP/pkg/jwta"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("Authorization")
		if token == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}

		_, err := jwta.ParseToken(token)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}
		// 可以进一步验证，通过查询是否存在该用户。
		c.Next()
	}
}
