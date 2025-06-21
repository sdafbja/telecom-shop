package middlewares

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/sdafbja/telecom-shop/config"
	"github.com/sdafbja/telecom-shop/utils"
)

// JWTAuthMiddleware kiểm tra token và blacklist
func JWTAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		auth := c.GetHeader("Authorization")
		if !strings.HasPrefix(auth, "Bearer ") {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Thiếu hoặc sai định dạng Authorization"})
			return
		}
		tokenStr := strings.TrimPrefix(auth, "Bearer ")

		// kiểm tra blacklist
		if val, _ := config.RedisClient.Get(config.Ctx, "blacklist:"+tokenStr).Result(); val == "true" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Token đã bị thu hồi"})
			return
		}

		// parse token
		claims, err := utils.ParseJWT(tokenStr)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Token không hợp lệ"})
			return
		}

		// lưu thông tin vào context
		c.Set("user_id", claims.UserID)
		c.Set("role", claims.Role)
		c.Next()
	}
}
func RequireAdmin() gin.HandlerFunc {
    return func(c *gin.Context) {
        role, exists := c.Get("role")
        if !exists || role != "admin" {
            c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "Chỉ admin mới được phép truy cập"})
            return
        }
        c.Next()
    }
}
