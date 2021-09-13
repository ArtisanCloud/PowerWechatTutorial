package miniprogram

import (
	"github.com/ArtisanCloud/power-wechat/src/miniProgram/uniformMessage/request"
	"github.com/gin-gonic/gin"
	"net/http"
	"power-wechat-tutorial/services"
)

// 下发小程序和公众号统一的服务消息
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/uniform-message/uniformMessage.send.html
func APIUniformMessageSend(c *gin.Context) {

	openID, exist := c.GetQuery("openID")
	if !exist {
		panic("parameter open id expected")
	}

	mpTemplateMsg := &request.MPTemplateMsg{
		AppID:       services.AppMiniProgram.GetConfig().GetString("app_id",""),
		TemplateID:  "MTvUCMmZfl-Dv66C5fVWdf4zPJkYSaRbnrtk6DXIfTQ",
		URL:         "",
		MiniProgram: "",
		Data:        nil,
	}

	rs, err := services.AppMiniProgram.UniformMessage.Send(openID, nil, mpTemplateMsg)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, rs)
}
