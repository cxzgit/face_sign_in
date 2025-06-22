package server

import (
	adminRouter "face-signIn/internal/adminManage/routers"
	teacherRouter "face-signIn/internal/teacherManage/routers"
	studentRouter "face-signIn/internal/userManage/routers"
	"face-signIn/pkg/corsMW"
	"face-signIn/pkg/globals"
)

// SetupRouter 启动处理函数
func SetupRouter() {
	// 跨域
	globals.Router.Use(corsMW.CorsMiddleware())

	// 空接口，不执行操作，用来 get X-CSRF-Token
	globals.Router.GET("/get_csrf_token")

	//学生管理的接口
	studentRouter.RegisterStudentRoutes(globals.Router)

	//教师管理的接口
	teacherRouter.RegisterTeacherRoutes(globals.Router)

	//管理员管理接口
	adminRouter.RegisterAdminRoutes(globals.Router)
}
