package miniprogram

import (
	"github.com/ArtisanCloud/power-wechat/src/miniProgram/uniformMessage/request"
	"github.com/gin-gonic/gin"
	"power-wechat-tutorial/services"
)

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

	c.JSON(200, rs)
}
