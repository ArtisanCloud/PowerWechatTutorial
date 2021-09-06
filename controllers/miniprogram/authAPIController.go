package miniprogram

import (
	"github.com/gin-gonic/gin"
	"power-wechat-tutorial/services"
)

func APISNSSession(c *gin.Context)  {


	rs , err:=services.AppMiniProgram.Auth.Session("123")

	if err != nil {
		panic(err)
	}

	c.JSON(200, rs)

}