package middlewares

import (
	"face-signIn/pkg/globals"
	"face-signIn/pkg/response"
	"face-signIn/pkg/utils"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// JWTMiddleware 是一个JWT认证中间件
func JWTMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 1. 从请求头获取Token
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			response.Failed(c, http.StatusUnauthorized, response.NewAppErr(globals.StatusUnauthorized, fmt.Errorf("请求未携带Token，无访问权限"), nil))
			c.Abort()
			return
		}

		// 2. 校验Token格式
		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			response.Failed(c, http.StatusUnauthorized, response.NewAppErr(globals.StatusUnauthorized, fmt.Errorf("Token格式不正确"), nil))
			c.Abort()
			return
		}

		// 3. 解析Token
		tokenString := parts[1]
		claims, err := utils.ParseToken(tokenString)
		if err != nil {
			response.Failed(c, http.StatusUnauthorized, response.NewAppErr(globals.StatusUnauthorized, fmt.Errorf("无效的Token"), nil))
			c.Abort()
			return
		}

		// 4. 将用户信息注入上下文
		c.Set("userID", claims.ID)
		c.Set("userName", claims.Name)
		c.Set("userRole", claims.Role)

		// 5. 继续处理请求
		c.Next()
	}
}
