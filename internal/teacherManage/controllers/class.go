package controllers

import (
	"face-signIn/internal/teacherManage/logics"
	"face-signIn/internal/teacherManage/responses"
	"face-signIn/pkg/globals"
	"face-signIn/pkg/response"
	"fmt"
	"github.com/gin-gonic/gin"
)

// GetAllClasses godoc
// @Summary      获取所有班级列表
// @Tags         class
// @Produce      json
// @Success      200   {object}  response.AppData
// @Router       /teacher/class/list [get]
func GetAllClasses(c *gin.Context) {
	classes, err := logics.GetAllClasses()
	if err != nil {
		response.Failed(c, 200, response.NewAppErr(globals.StatusInternalServerError, fmt.Errorf("查询失败: %v", err), nil))
		return
	}
	var respList []responses.ClassInfo
	for _, class := range classes {
		respList = append(respList, *responses.NewClassInfo(class.ID, class.Name))
	}
	response.Success(c, 200, response.NewAppData(globals.StatusOK, "查询成功", respList))
}
