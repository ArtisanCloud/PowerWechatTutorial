package miniprogram

import (
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/power"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/miniProgram/uniformMessage/request"
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

	rs, err := services.MiniProgramApp.UniformMessage.Send(c.Request.Context(),
		&request.RequestUniformMessageSend{
			ToUser: openID,
			WeAppTemplateMsg: &request.WeAppTemplateMsg{
				TemplateID: "TEMPLATE_ID",
				Page:       "page/page/index",
				FormID:     "FORMID",
				Data: &power.HashMap{
					"keyword1": &power.HashMap{
						"value": "339208499",
					},
					"keyword2": &power.HashMap{
						"value": "2015年01月05日 12:30",
					},
					"keyword3": &power.HashMap{
						"value": "腾讯微信总部",
					},
					"keyword4": &power.HashMap{
						"value": "广州市海珠区新港中路397号",
					},
				},
				EmphasisKeyword: "keyword1.DATA",
			},
			MpTemplateMsg: &request.MPTemplateMsg{
				AppID:      services.MiniProgramApp.GetConfig().GetString("app_id", ""),
				TemplateID: "MTvUCMmZfl-Dv66C5fVWdf4zPJkYSaRbnrtk6DXIfTQ",
				Url:        "https://weixin.qq.com/download",
				MiniProgram: &request.MPTemplateMsgMiniProgram{
					AppID:    "",
					PagePath: "",
				},
				Data: &power.HashMap{
					"first": &power.HashMap{
						"value": "恭喜你购买成功！",
						"color": "#173177",
					},
					"keyword1": &power.HashMap{
						"value": "巧克力",
						"color": "#173177",
					},
					"keyword2": &power.HashMap{
						"value": "39.8元",
						"color": "#173177",
					},
					"keyword3": &power.HashMap{
						"value": "2014年9月22日",
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
