package routers

import (
	"face-signIn/internal/middlewares"
	"face-signIn/internal/teacherManage/controllers"
	"github.com/gin-gonic/gin"
)

// RegisterTeacherRoutes 注册教师端所有相关接口路由
func RegisterTeacherRoutes(e *gin.Engine) {
	// 教师接口根路由
	teacherAPI := e.Group("/teacher")

	// --- 开放接口 (无需JWT认证) ---
	teacherAPI.POST("/register", controllers.TeacherRegister) // 教师注册
	teacherAPI.POST("/login", controllers.TeacherLogin)       // 教师登录

	// --- 认证接口 (需要JWT认证) ---
	authGroup := teacherAPI.Group("/")
	authGroup.Use(middlewares.JWTMiddleware())
	{
		// 账号相关
		authGroup.POST("/logout", controllers.TeacherLogout) // 教师登出

		// 课程管理
		courseGroup := authGroup.Group("/course")
		{
			courseGroup.POST("/create", controllers.CreateCourse) // 创建课程
			courseGroup.GET("/list", controllers.GetMyCourses)    // 查询我的课程
			courseGroup.POST("/update", controllers.UpdateCourse) // 更新课程
			courseGroup.POST("/delete", controllers.DeleteCourse) // 删除课程
		}

		// 班级管理
		classGroup := authGroup.Group("/class")
		{
			classGroup.GET("/list", controllers.GetAllClasses) // 获取所有班级列表 (用于下拉选择)
		}

		// 课程-班级绑定
		courseClassGroup := authGroup.Group("/course_class")
		{
			courseClassGroup.POST("/bind", controllers.BindCourseClass)     // 课程批量绑定班级
			courseClassGroup.POST("/unbind", controllers.UnbindCourseClass) // 课程解绑班级
		}

		// 学生管理
		studentGroup := authGroup.Group("/student")
		{
			studentGroup.POST("/create", controllers.CreateStudent)   // 添加学生
			studentGroup.POST("/update", controllers.UpdateStudent)   // 编辑学生
			studentGroup.POST("/delete", controllers.DeleteStudent)   // 删除学生
			studentGroup.GET("/list", controllers.GetStudentsByClass) // 查询班级下的学生列表
		}

		// 签到管理
		signInGroup := authGroup.Group("/sign_in")
		{
			signInGroup.POST("/create", controllers.CreateSignInTask)    // 发起签到 (可指定多个班级)
			signInGroup.GET("/records", controllers.GetSignInRecords)    // 查询指定签到任务的签到结果
			signInGroup.POST("/manual", controllers.TeacherManualSignIn) // 手动为学生代签到
		}
	}
}
