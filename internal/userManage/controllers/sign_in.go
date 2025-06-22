package controllers

import (
	"face-signIn/internal/userManage/logics"
	"face-signIn/internal/userManage/repositories"
	"face-signIn/internal/userManage/requests"
	"face-signIn/pkg/globals"
	"face-signIn/pkg/response"
	"fmt"
	"github.com/gin-gonic/gin"
)

// GetPendingSignInTasks godoc
// @Summary      学生获取待签到任务
// @Tags         sign_in
// @Produce      json
// @Success      200   {object}  response.AppData
// @Router       /student/sign_in/pending [get]
func GetPendingSignInTasks(c *gin.Context) {
	studentID := c.GetUint("userID")
	student, err := repositories.GetStudentByID(studentID)
	if err != nil {
		response.Failed(c, 200, response.NewAppErr(globals.StatusBadRequest, err, nil))
		return
	}
	tasks, err := logics.GetPendingSignInTasksForStudent(studentID, student.ClassID)
	if err != nil {
		response.Failed(c, 200, response.NewAppErr(globals.StatusInternalServerError, err, nil))
		return
	}
	response.Success(c, 200, response.NewAppData(globals.StatusOK, "查询成功", tasks))
}

// GetSignInRecordsByStudent godoc
// @Summary      学生查询个人签到历史
// @Tags         学生签到
// @Produce      json
// @Param        Authorization header string true "Bearer <token>"
// @Success      200   {object}  response.AppData
// @Router       /student/sign_in/history [get]
func GetSignInRecordsByStudent(c *gin.Context) {
	studentID := c.GetUint("userID")
	if studentID == 0 {
		response.Failed(c, 200, response.NewAppErr(globals.StatusUnauthorized, fmt.Errorf("用户未登录或Token无效"), nil))
		return
	}

	records, err := logics.GetSignInRecordsByStudentID(studentID)
	if err != nil {
		response.Failed(c, 200, response.NewAppErr(globals.StatusInternalServerError, err, nil))
		return
	}

	response.Success(c, 200, response.NewAppData(globals.StatusOK, "查询成功", records))
}

// StudentSignIn godoc
// @Summary      学生进行人脸签到
// @Tags         学生签到
// @Accept       json
// @Produce      json
// @Param        Authorization header string true "Bearer <token>"
// @Param        data  body  requests.StudentSignInRequest  true  "签到信息"
// @Success      200   {object}  response.AppData
// @Router       /student/sign_in/do [post]
func StudentSignIn(c *gin.Context) {
	var req requests.StudentSignInRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Failed(c, 200, response.NewAppErr(globals.StatusBadRequest, fmt.Errorf("参数错误: %v", err), nil))
		return
	}

	// 从JWT中间件获取学生ID
	studentID := c.GetUint("userID")
	if studentID == 0 {
		response.Failed(c, 200, response.NewAppErr(globals.StatusInternalServerError, fmt.Errorf("用户未登录或Token无效"), nil))
		return
	}

	err := logics.StudentSignIn(studentID, req.SignInTaskID, req.FaceImage)
	if err != nil {
		response.Failed(c, 200, response.NewAppErr(globals.StatusInternalServerError, err, nil))
		return
	}

	response.Success(c, 200, response.NewAppData(globals.StatusOK, "签到成功！", nil))
}
