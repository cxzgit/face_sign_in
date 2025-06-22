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

// CreateCourse godoc
// @Summary      创建课程
// @Tags         course
// @Accept       json
// @Produce      json
// @Param        data  body  requests.CreateCourseRequest  true  "课程信息"
// @Success      200   {object}  response.AppData
// @Router       /teacher/course/create [post]
func CreateCourse(c *gin.Context) {
	var req requests.CreateCourseRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Failed(c, 200, response.NewAppErr(globals.StatusBadRequest, fmt.Errorf("参数错误"), nil))
		return
	}
	teacherID := c.GetUint("userID")
	err := logics.CreateCourse(teacherID, req.Name, req.ClassName, req.Description)
	if err != nil {
		response.Failed(c, 200, response.NewAppErr(globals.StatusInternalServerError, fmt.Errorf("创建失败: %v", err), nil))
		return
	}
	response.Success(c, 200, response.NewAppData(globals.StatusOK, "创建成功", nil))
}

// GetMyCourses godoc
// @Summary      查询我的课程
// @Tags         course
// @Produce      json
// @Success      200   {object}  response.AppData
// @Router       /teacher/course/list [get]
func GetMyCourses(c *gin.Context) {
	teacherID := c.GetUint("userID")
	courses, err := logics.GetCoursesByTeacherID(teacherID)
	if err != nil {
		response.Failed(c, 200, response.NewAppErr(globals.StatusInternalServerError, fmt.Errorf("查询失败: %v", err), nil))
		return
	}
	var respList []responses.CourseInfo
	for _, course := range courses {
		respList = append(respList, *responses.NewCourseInfo(course.ID, course.Name, course.TeacherID, course.ClassName, course.Description))
	}
	response.Success(c, 200, response.NewAppData(globals.StatusOK, "查询成功", respList))
}

// UpdateCourse godoc
// @Summary      更新课程
// @Tags         course
// @Accept       json
// @Produce      json
// @Param        data  body  requests.UpdateCourseRequest  true  "课程信息"
// @Success      200   {object}  response.AppData
// @Router       /teacher/course/update [post]
func UpdateCourse(c *gin.Context) {
	var req requests.UpdateCourseRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Failed(c, 200, response.NewAppErr(globals.StatusBadRequest, fmt.Errorf("参数错误"), nil))
		return
	}
	teacherID := c.GetUint("userID")
	err := logics.UpdateCourse(teacherID, req.ID, req.Name, req.ClassName, req.Description)
	if err != nil {
		response.Failed(c, 200, response.NewAppErr(globals.StatusInternalServerError, fmt.Errorf("更新失败: %v", err), nil))
		return
	}
	response.Success(c, 200, response.NewAppData(globals.StatusOK, "更新成功", nil))
}

// DeleteCourse godoc
// @Summary      删除课程
// @Tags         course
// @Accept       json
// @Produce      json
// @Param        id  query  int  true  "课程ID"
// @Success      200   {object}  response.AppData
// @Router       /teacher/course/delete [post]
func DeleteCourse(c *gin.Context) {
	id := c.Query("id")
	teacherID := c.GetUint("userID")
	var courseID uint
	parsedID, err := strconv.ParseUint(id, 10, 0) // 返回 uint64 类型
	if err != nil {
		response.Failed(c, 200, response.NewAppErr(globals.StatusBadRequest, fmt.Errorf("参数错误"), nil))
		return
	}
	courseID = uint(parsedID)
	err = logics.DeleteCourse(teacherID, courseID)
	if err != nil {
		response.Failed(c, 200, response.NewAppErr(globals.StatusInternalServerError, fmt.Errorf("删除失败: %v", err), nil))
		return
	}
	response.Success(c, 200, response.NewAppData(globals.StatusOK, "删除成功", nil))
}
