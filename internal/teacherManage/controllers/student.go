package controllers

import (
	"face-signIn/internal/teacherManage/logics"
	"face-signIn/internal/teacherManage/requests"
	"face-signIn/internal/teacherManage/responses"
	"face-signIn/pkg/globals"
	"face-signIn/pkg/response"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)

// CreateStudent godoc
// @Summary      教师添加学生
// @Tags         student
// @Accept       json
// @Produce      json
// @Param        data  body  requests.CreateStudentRequest  true  "学生信息"
// @Success      200   {object}  response.AppData
// @Router       /teacher/student/create [post]
func CreateStudent(c *gin.Context) {
	var req requests.CreateStudentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Failed(c, 200, response.NewAppErr(globals.StatusBadRequest, fmt.Errorf("参数错误"), nil))
		return
	}
	if err := logics.CreateStudent(req.StudentID, req.Name, req.Password, req.ClassID); err != nil {
		response.Failed(c, 200, response.NewAppErr(globals.StatusInternalServerError, fmt.Errorf("添加失败: %v", err), nil))
		return
	}
	response.Success(c, 200, response.NewAppData(globals.StatusOK, "添加成功", nil))
}

// UpdateStudent godoc
// @Summary      教师编辑学生
// @Tags         student
// @Accept       json
// @Produce      json
// @Param        data  body  requests.UpdateStudentRequest  true  "学生信息"
// @Success      200   {object}  response.AppData
// @Router       /teacher/student/update [post]
func UpdateStudent(c *gin.Context) {
	var req requests.UpdateStudentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Failed(c, 200, response.NewAppErr(globals.StatusBadRequest, fmt.Errorf("参数错误"), nil))
		return
	}
	if err := logics.UpdateStudent(req.ID, req.Name, req.Password, req.ClassID); err != nil {
		response.Failed(c, 200, response.NewAppErr(globals.StatusInternalServerError, fmt.Errorf("编辑失败: %v", err), nil))
		return
	}
	response.Success(c, 200, response.NewAppData(globals.StatusOK, "编辑成功", nil))
}

// DeleteStudent godoc
// @Summary      教师删除学生
// @Tags         student
// @Accept       json
// @Produce      json
// @Param        id  query  int  true  "学生ID"
// @Success      200   {object}  response.AppData
// @Router       /teacher/student/delete [post]
func DeleteStudent(c *gin.Context) {
	id := c.Query("id")
	var studentID uint
	if _, err := fmt.Sscan(id, &studentID); err != nil {
		response.Failed(c, 200, response.NewAppErr(globals.StatusBadRequest, fmt.Errorf("参数错误"), nil))
		return
	}
	if err := logics.DeleteStudentByID(studentID); err != nil {
		response.Failed(c, 200, response.NewAppErr(globals.StatusInternalServerError, fmt.Errorf("删除失败: %v", err), nil))
		return
	}
	response.Success(c, 200, response.NewAppData(globals.StatusOK, "删除成功", nil))
}

// GetStudentsByClass godoc
// @Summary      教师查询班级学生列表
// @Tags         student
// @Produce      json
// @Param        class_id  query  int  true  "班级ID"
// @Success      200   {object}  response.AppData
// @Router       /teacher/student/list [get]
func GetStudentsByClass(c *gin.Context) {
	classIDStr := c.Query("class_id")
	classID, err := strconv.ParseUint(classIDStr, 10, 64)
	if err != nil {
		response.Failed(c, 200, response.NewAppErr(globals.StatusBadRequest, fmt.Errorf("参数错误"), nil))
		return
	}
	students, err := logics.GetStudentsByClass(uint(classID))
	if err != nil {
		response.Failed(c, 200, response.NewAppErr(globals.StatusInternalServerError, fmt.Errorf("查询失败: %v", err), nil))
		return
	}
	var respList []responses.StudentInfo
	for _, stu := range students {
		respList = append(respList, *responses.NewStudentInfo(stu.ID, stu.StudentID, stu.Name, stu.ClassID))
	}
	response.Success(c, 200, response.NewAppData(globals.StatusOK, "查询成功", respList))
}
