package miniprogram

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"power-wechat-tutorial/services"
)

func GetUserPhoneNumber(c *gin.Context) {

	code, exist := c.GetQuery("code")
	if !exist {
		panic("parameter code expected")
	}

	rs, err := services.MiniProgramApp.PhoneNumber.GetUserPhoneNumber(c.Request.Context(), code)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, rs)
}

func GetUserPhoneNumberByAES(c *gin.Context) {
	encryptData := c.Query("encryptData")
	sessionKey := c.Query("sessionKey")
	iv := c.Query("iv")

	data, err := services.MiniProgramApp.Encryptor.DecryptData(encryptData, sessionKey, iv)
	if err != nil {
		panic(err)
	}

	c.String(200, string(data))
}
