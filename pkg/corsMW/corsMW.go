package corsMW

import (
	"github.com/gin-gonic/gin"
)

// CorsMiddleware 跨域中间件
func CorsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		origin := c.Request.Header.Get("Origin") // 获取请求来源

		// 允许的前端域名列表
		allowedOrigins := map[string]bool{
			"http://127.0.0.1:8081": true,
		}

		// 判断 origin 是否在允许列表内
		if allowedOrigins[origin] {
			c.Header("Access-Control-Allow-Origin", origin) // 仅允许特定域名
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE,PATCH")
			c.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization")
			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Cache-Control, Content-Language, Content-Type")
			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Cache-Control, Content-Language, Content-Type, X-Csrf-Token")
			c.Header("Access-Control-Allow-Credentials", "true") // 允许跨域请求携带 Cookie
		}

		// 处理预检请求（OPTIONS 请求）
		if method == "OPTIONS" {
			c.AbortWithStatus(204) // 直接返回 204 状态码，表示接受预检请求
			return
		}

		c.Next()
	}
}
