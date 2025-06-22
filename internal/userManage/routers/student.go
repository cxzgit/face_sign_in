package routers

import (
	"face-signIn/internal/middlewares"
	"face-signIn/internal/userManage/controllers"
	"github.com/gin-gonic/gin"
)

// RegisterStudentRoutes 注册学生端所有相关接口
func RegisterStudentRoutes(e *gin.Engine) {
	// 学生接口根路由
	studentAPI := e.Group("/student")

	// --- 开放接口 (无需JWT认证) ---
	studentAPI.POST("/register", controllers.StudentRegister) // 学生注册
	studentAPI.POST("/login", controllers.StudentLogin)       // 学生登录

	// --- 认证接口 (需要JWT认证) ---
	authGroup := studentAPI.Group("/")
	authGroup.Use(middlewares.JWTMiddleware())
	{
		// 账号相关
		authGroup.POST("/logout", controllers.StudentLogout) // 学生登出

		// 签到相关
		signInGroup := authGroup.Group("/sign_in")
		{
			signInGroup.POST("/do", controllers.StudentSignIn)                 // 执行人脸签到
			signInGroup.GET("/pending", controllers.GetPendingSignInTasks)     // 获取待签到任务列表
			signInGroup.GET("/history", controllers.GetSignInRecordsByStudent) // 获取个人签到历史
		}
	}
}
