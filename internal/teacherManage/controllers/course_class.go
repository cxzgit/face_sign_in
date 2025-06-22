package controllers

import (
	"face-signIn/internal/teacherManage/logics"
	"face-signIn/internal/teacherManage/requests"
	"face-signIn/pkg/globals"
	"face-signIn/pkg/response"
	"fmt"
	"github.com/gin-gonic/gin"
)

// 绑定课程和班级
func BindCourseClass(c *gin.Context) {
	var req requests.BindCourseClassRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Failed(c, 200, response.NewAppErr(globals.StatusBadRequest, fmt.Errorf("参数错误"), nil))
		return
	}
	if err := logics.BindCourseClass(req.CourseID, req.ClassIDs); err != nil {
		response.Failed(c, 200, response.NewAppErr(globals.StatusInternalServerError, fmt.Errorf("绑定失败: %v", err), nil))
		return
	}
	response.Success(c, 200, response.NewAppData(globals.StatusOK, "绑定成功", nil))
}

// 解绑课程和班级
func UnbindCourseClass(c *gin.Context) {
	var req requests.UnbindCourseClassRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Failed(c, 200, response.NewAppErr(globals.StatusBadRequest, fmt.Errorf("参数错误"), nil))
		return
	}
	if err := logics.UnbindCourseClass(req.CourseID, req.ClassID); err != nil {
		response.Failed(c, 200, response.NewAppErr(globals.StatusInternalServerError, fmt.Errorf("解绑失败: %v", err), nil))
		return
	}
	response.Success(c, 200, response.NewAppData(globals.StatusOK, "解绑成功", nil))
}
