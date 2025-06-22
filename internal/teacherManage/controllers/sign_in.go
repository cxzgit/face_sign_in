package controllers

import (
	"face-signIn/internal/teacherManage/logics"
	"face-signIn/internal/teacherManage/requests"
	"face-signIn/internal/teacherManage/responses"
	"face-signIn/pkg/globals"
	"face-signIn/pkg/response"
	"fmt"
	"github.com/gin-gonic/gin"
)

// CreateSignInTask godoc
// @Summary      发起签到
// @Tags         sign_in
// @Accept       json
// @Produce      json
// @Param        data  body  requests.CreateSignInTaskRequest  true  "签到任务信息"
// @Success      200   {object}  response.AppData
// @Router       /teacher/sign_in/create [post]
func CreateSignInTask(c *gin.Context) {
	var req requests.CreateSignInTaskRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Failed(c, 200, response.NewAppErr(globals.StatusBadRequest, fmt.Errorf("参数错误"), nil))
		return
	}
	teacherID := c.GetUint("userID")
	err := logics.CreateSignInTask(teacherID, req.CourseID, req.ClassIDs, req.StartTime, req.EndTime, req.Description)
	if err != nil {
		response.Failed(c, 200, response.NewAppErr(globals.StatusInternalServerError, fmt.Errorf("发起签到失败: %v", err), nil))
		return
	}
	response.Success(c, 200, response.NewAppData(globals.StatusOK, "发起签到成功", nil))
}

// StudentSignIn godoc
// @Summary      学生签到
// @Tags         sign_in
// @Accept       json
// @Produce      json
// @Param        data  body  requests.StudentSignInRequest  true  "签到信息"
// @Success      200   {object}  response.AppData
// @Router       /student/sign_in/do [post]
func StudentSignIn(c *gin.Context) {
	var req requests.StudentSignInRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Failed(c, 200, response.NewAppErr(globals.StatusBadRequest, fmt.Errorf("参数错误"), nil))
		return
	}
	studentID := c.GetUint("userID")
	err := logics.StudentSignIn(studentID, req.SignInTaskID, req.FaceImage)
	if err != nil {
		response.Failed(c, 200, response.NewAppErr(globals.StatusInternalServerError, fmt.Errorf("签到失败: %v", err), nil))
		return
	}
	response.Success(c, 200, response.NewAppData(globals.StatusOK, "签到成功", nil))
}

// GetSignInRecords godoc
// @Summary      查询签到结果
// @Tags         sign_in
// @Produce      json
// @Param        sign_in_task_id  query  int  true  "签到任务ID"
// @Success      200   {object}  response.AppData
// @Router       /teacher/sign_in/records [get]
func GetSignInRecords(c *gin.Context) {
	taskID := c.Query("sign_in_task_id")
	var id uint
	if _, err := fmt.Sscan(taskID, &id); err != nil {
		response.Failed(c, 200, response.NewAppErr(globals.StatusBadRequest, fmt.Errorf("参数错误"), nil))
		return
	}
	records, err := logics.GetSignInRecordsByTaskID(id)
	if err != nil {
		response.Failed(c, 200, response.NewAppErr(globals.StatusInternalServerError, fmt.Errorf("查询失败: %v", err), nil))
		return
	}
	var respList []responses.SignInRecordInfo
	for _, record := range records {
		respList = append(respList, *responses.NewSignInRecordInfo(record.ID, record.SignInTaskID, record.StudentID, record.SignInTime, record.FaceImage, record.Status))
	}
	response.Success(c, 200, response.NewAppData(globals.StatusOK, "查询成功", respList))
}
