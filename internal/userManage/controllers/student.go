package controllers

import (
	"face-signIn/internal/userManage/logics"
	"face-signIn/internal/userManage/requests"
	"face-signIn/internal/userManage/responses"
	"face-signIn/pkg/globals"
	"face-signIn/pkg/response"
	"face-signIn/pkg/utils"
	"github.com/gin-gonic/gin"
)

// StudentRegister godoc
// @Summary      学生注册
// @Description  学生通过学号注册
// @Tags         student
// @Accept       json
// @Produce      json
// @Param        data  body  requests.StudentRegisterRequest  true  "注册参数"
// @Success      200   {object}  response.AppData
// @Failure      200   {object}  response.AppErr
// @Router       /student/student/register [post]
func StudentRegister(c *gin.Context) {
	var req requests.StudentRegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Failed(c, 200, response.NewAppErr(globals.StatusBadRequest, err, nil))
		return
	}
	if err := logics.RegisterStudent(req.StudentID, req.Name, req.Password); err != nil {
		response.Failed(c, 200, response.NewAppErr(globals.StatusBadRequest, err, nil))
		return
	}
	response.Success(c, 200, response.NewAppData(globals.StatusOK, "注册成功", nil))
}

// StudentLogin godoc
// @Summary      学生登录
// @Description  学生通过学号登录
// @Tags         student
// @Accept       json
// @Produce      json
// @Param        data  body  requests.StudentLoginRequest  true  "登录参数"
// @Success      200   {object}  response.AppData
// @Failure      200   {object}  response.AppErr
// @Router       /student/student/login [post]
func StudentLogin(c *gin.Context) {
	var req requests.StudentLoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Failed(c, 200, response.NewAppErr(globals.StatusBadRequest, err, nil))
		return
	}
	student, err := logics.LoginStudent(req.StudentID, req.Password)
	if err != nil {
		response.Failed(c, 200, response.NewAppErr(globals.StatusBadRequest, err, nil))
		return
	}
	resp := responses.NewStudentInfo(student.StudentID, student.Name)
	token, _ := utils.GenerateToken(student.ID, student.Name, "student")
	response.Success(c, 200, response.NewAppData(globals.StatusOK, "登录成功", gin.H{
		"token":   token,
		"student": resp,
	}))
}

// StudentLogout godoc
// @Summary      学生登出
// @Description  学生登出接口
// @Tags         student
// @Accept       json
// @Produce      json
// @Success      200   {object}  response.AppData
// @Router       /student/student/logout [post]
func StudentLogout(c *gin.Context) {
	// 这里只做演示，实际项目可清除 session/token
	response.Success(c, 200, response.NewAppData(globals.StatusOK, "登出成功", nil))
}
