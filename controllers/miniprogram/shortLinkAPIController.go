package miniprogram

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"power-wechat-tutorial/services"
)

// 获取小程序 Short Link
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/short-link/shortlink.generate.html
func APIShortLinkGenerate(c *gin.Context) {

	pageUrl := "/pages/publishHomework/publishHomework?query1=q1"
	pageTitle := "Homework title"
	isPermanent := false

	rs, err := services.MiniProgramApp.ShortLink.Generate(pageUrl, pageTitle, isPermanent)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, rs)
}
