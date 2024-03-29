package miniprogram

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"power-wechat-tutorial/services"
)

// 获取小程序 Short Link
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/short-link/shortlink.generate.html
func APIShortLinkGenerate(c *gin.Context) {

	pageUrl := c.DefaultQuery("pageUrl", "/pages/index/index?query1=q1")
	pageTitle := "Homework title"
	isPermanent := false

	rs, err := services.MiniProgramApp.ShortLink.Generate(c.Request.Context(), pageUrl, pageTitle, isPermanent)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, rs)
}
