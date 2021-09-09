package miniprogram

import (
	"github.com/gin-gonic/gin"
	"power-wechat-tutorial/services"
)

func APIInternetGetUserEncryptKey(c *gin.Context) {

	openID, exist := c.GetQuery("openID")
	if !exist {
		panic("parameter open id expected")
	}

	rs, err := services.AppMiniProgram.Internet.GetUserEncryptKey(openID,"","hmac_sha256")

	if err != nil {
		panic(err)
	}

	c.JSON(200, rs)
}
