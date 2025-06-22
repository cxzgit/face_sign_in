package controllers

import (
	"face-signIn/internal/teacherManage/logics"
	"face-signIn/internal/teacherManage/requests"
	"face-signIn/pkg/globals"
	"face-signIn/pkg/response"
	"fmt"
	"github.com/gin-gonic/gin"
)

// TeacherManualSignIn godoc
// @Summary      老师手动为学生代签到
// @Tags         sign_in
// @Accept       json
// @Produce      json
// @Param        data  body  requests.TeacherManualSignInRequest  true  "代签到信息"
// @Success      200   {object}  response.AppData
// @Router       /teacher/sign_in/manual [post]
func TeacherManualSignIn(c *gin.Context) {
	var req requests.TeacherManualSignInRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Failed(c, 200, response.NewAppErr(globals.StatusBadRequest, fmt.Errorf("参数错误"), nil))
		return
	}
	teacherID := c.GetUint("userID")
	err := logics.TeacherManualSignIn(teacherID, req.SignInTaskID, req.StudentID, req.Remark)
	if err != nil {
		response.Failed(c, 200, response.NewAppErr(globals.StatusInternalServerError, fmt.Errorf("代签到失败: %v", err), nil))
		return
	}
	response.Success(c, 200, response.NewAppData(globals.StatusOK, "代签到成功", nil))
}
