package routers

import (
	"face-signIn/internal/adminManage/controllers"
	"github.com/gin-gonic/gin"
)

// RegisterAdminRoutes 注册管理员账号相关接口路由，包括登录、登出
func RegisterAdminRoutes(e *gin.Engine) {
	// 管理员账号相关接口分组
	adminGroup := e.Group("/admin")
	// 管理员登录
	adminGroup.POST("/admin/login", controllers.AdminLogin)
	// 管理员登出
	adminGroup.POST("/admin/logout", controllers.AdminLogout)
}
