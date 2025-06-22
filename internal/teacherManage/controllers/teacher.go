package controllers

import (
	"face-signIn/internal/teacherManage/logics"
	"face-signIn/internal/teacherManage/requests"
	"face-signIn/internal/teacherManage/responses"
	"face-signIn/pkg/globals"
	"face-signIn/pkg/response"
	"face-signIn/pkg/utils"
	"fmt"
	"github.com/gin-gonic/gin"
)

// TeacherRegister godoc
// @Summary      教师注册
// @Description  教师通过教师号注册
// @Tags         teacher
// @Accept       json
// @Produce      json
// @Param        data  body  requests.TeacherRegisterRequest  true  "注册参数"
// @Success      200   {object}  response.AppData
// @Failure      200   {object}  response.AppErr
// @Router       /teacher/teacher/register [post]
func TeacherRegister(c *gin.Context) {
	var req requests.TeacherRegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Failed(c, 200, response.NewAppErr(globals.StatusBadRequest, fmt.Errorf("参数错误"), nil))
		return
	}
	if err := logics.RegisterTeacher(req.TeacherID, req.Name, req.Password); err != nil {
		response.Failed(c, 200, response.NewAppErr(globals.StatusBadRequest, fmt.Errorf("注册失败: %v", err), nil))
		return
	}
	response.Success(c, 200, response.NewAppData(globals.StatusOK, "注册成功", nil))
}

// TeacherLogin godoc
// @Summary      教师登录
// @Description  教师通过教师号登录
// @Tags         teacher
// @Accept       json
// @Produce      json
// @Param        data  body  requests.TeacherLoginRequest  true  "登录参数"
// @Success      200   {object}  response.AppData
// @Failure      200   {object}  response.AppErr
// @Router       /teacher/teacher/login [post]
func TeacherLogin(c *gin.Context) {
	var req requests.TeacherLoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Failed(c, 200, response.NewAppErr(globals.StatusBadRequest, fmt.Errorf("参数错误"), nil))
		return
	}
	teacher, err := logics.LoginTeacher(req.TeacherID, req.Password)
	if err != nil {
		response.Failed(c, 200, response.NewAppErr(globals.StatusBadRequest, fmt.Errorf("登录失败: %v", err), nil))
		return
	}
	resp := responses.NewTeacherInfo(teacher.TeacherID, teacher.Name)
	token, _ := utils.GenerateToken(teacher.ID, teacher.Name, "teacher")
	response.Success(c, 200, response.NewAppData(globals.StatusOK, "登录成功", gin.H{
		"token":   token,
		"teacher": resp,
	}))
}

// TeacherLogout godoc
// @Summary      教师登出
// @Description  教师登出接口
// @Tags         teacher
// @Accept       json
// @Produce      json
// @Success      200   {object}  response.AppData
// @Router       /teacher/teacher/logout [post]
func TeacherLogout(c *gin.Context) {
	response.Success(c, 200, response.NewAppData(globals.StatusOK, "登出成功", nil))
}
