package miniprogram

import (
	"github.com/ArtisanCloud/PowerWeChat/v3/src/miniProgram/urlLink/request"
	"github.com/gin-gonic/gin"
	"net/http"
	"power-wechat-tutorial/services"
)

// 获取小程序 URL Link，适用于短信、邮件、网页、微信内等拉起小程序的业务场景
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/url-link/urllink.generate.html
func APIURLLinkGenerate(c *gin.Context) {

	path := c.DefaultQuery("path", "pages/index/index")

	rs, err := services.MiniProgramApp.URLLink.Generate(&request.URLLinkGenerate{
		EnvVersion:     "release",
		ExpireInterval: 1606737600,
		Path:           path,
		Query:          "a=1",
	})

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, rs)
}
