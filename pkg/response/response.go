package response

import (
	"github.com/gin-gonic/gin"
)

func Success(c *gin.Context, status int, data *AppData) {
	//status 指的是 http.StatusOK等之类的状态码
	c.JSON(status, data)
}

func Failed(c *gin.Context, status int, err *AppErr) {
	c.JSON(status, err)
}
