package controllers

import (
	"face-signIn/internal/adminManage/logics"
	"face-signIn/internal/adminManage/requests"
	"face-signIn/internal/adminManage/responses"
	"face-signIn/pkg/globals"
	"face-signIn/pkg/response"
	"face-signIn/pkg/utils"
	"github.com/gin-gonic/gin"
)

// AdminLogin godoc
// @Summary      管理员登录
// @Description  管理员通过账号登录
// @Tags         admin
// @Accept       json
// @Produce      json
// @Param        data  body  requests.AdminLoginRequest  true  "登录参数"
// @Success      200   {object}  response.AppData
// @Failure      200   {object}  response.AppErr
// @Router       /admin/admin/login [post]
func AdminLogin(c *gin.Context) {
	var req requests.AdminLoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Failed(c, 200, response.NewAppErr(globals.StatusBadRequest, err, nil))
		return
	}
	admin, err := logics.LoginAdmin(req.AdminID, req.Password)
	if err != nil {
		response.Failed(c, 200, response.NewAppErr(globals.StatusBadRequest, err, nil))
		return
	}
	resp := responses.NewAdminInfo(admin.AdminID, admin.Name)
	token, _ := utils.GenerateToken(admin.ID, admin.Name, "admin")
	response.Success(c, 200, response.NewAppData(globals.StatusOK, "登录成功", gin.H{
		"token": token,
		"admin": resp,
	}))
}

// AdminLogout godoc
// @Summary      管理员登出
// @Description  管理员登出接口
// @Tags         admin
// @Accept       json
// @Produce      json
// @Success      200   {object}  response.AppData
// @Router       /admin/admin/logout [post]
func AdminLogout(c *gin.Context) {
	response.Success(c, 200, response.NewAppData(globals.StatusOK, "登出成功", nil))
}
