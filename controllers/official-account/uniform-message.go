package official_account

import (
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/power"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/miniProgram/uniformMessage/request"
	"github.com/gin-gonic/gin"
	"net/http"
	"power-wechat-tutorial/services"
)

// 下发小程序和公众号统一的服务消息
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/uniform-message/uniformMessage.send.html
func UniformMessageSend(c *gin.Context) {

	openID, exist := c.GetQuery("openID")
	if !exist {
		panic("parameter open id expected")
	}
	templateID := c.DefaultQuery("templateID", "SPNl7d_nnSnssgyj7ZJ6v9B-N04ZBspts_PYvdF3D-8")

	rs, err := services.MiniProgramApp.UniformMessage.Send(c.Request.Context(), &request.RequestUniformMessageSend{
		ToUser: openID,
		MpTemplateMsg: &request.MPTemplateMsg{
			AppID:      services.OfficialAccountApp.GetConfig().GetString("app_id", ""),
			TemplateID: templateID,
			Url:        "https://powerwechat.artisan-cloud.com/",
			//MiniProgram: &request.MPTemplateMsgMiniProgram{
			//	AppID:    "",
			//	PagePath: "",
			//},
			Data: &power.HashMap{
				"first": &power.HashMap{
					"value": "恭喜你购买成功！",
					"color": "#173177",
				},
				"DateTime": &power.HashMap{
					"value": "3-5 16:22",
					"color": "#173177",
				},
				"PayAmount": &power.HashMap{
					"value": "39.8元",
					"color": "#173177",
				},
				"Location": &power.HashMap{
					"value": "上海市长宁区",
					"color": "#173177",
				},
				"remark": &power.HashMap{
					"value": "欢迎再次购买！",
					"color": "#173177",
				},
			},
		},
	})

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, rs)
}
