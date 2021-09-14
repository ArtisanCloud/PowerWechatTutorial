package wecom

import (
	"github.com/ArtisanCloud/power-wechat/src/kernel/power"
	"github.com/gin-gonic/gin"
	"power-wechat-tutorial/services"
)

// SendTextMsg 企业微信内部应用主动发送消息
// https://open.work.weixin.qq.com/api/doc/90000/90135/90236
func APISendTextMsg(c *gin.Context) {
	toUser := c.DefaultQuery("toUser", "walle")
	res, err := services.WeComApp.Message.Send(&power.HashMap{
		"touser":  toUser,
		"msgtype": "text",
		"agentid": 1000004,
		"text": &power.StringMap{
			"content": "你的快递已到，请携带工卡前往邮件中心领取。\n出发前可查看<a href=\"http://work.weixin.qq.com\">邮件中心视频实况</a>，聪明避开排队。",
		},
		"safe":                     0,
		"enable_id_trans":          0,
		"enable_duplicate_check":   0,
		"duplicate_check_interval": 1800,
	})

	if err != nil {
		panic(err)
	}

	c.JSON(200, res)
}

// Recall 撤回应用消息
// https://open.work.weixin.qq.com/api/doc/90000/90135/94867
func APIRecallMsg(c *gin.Context) {
	msgID := c.Query("msgID")
	res, err := services.WeComApp.Message.Recall(msgID)

	if err != nil {
		panic(err)
	}

	c.JSON(200, res)
}
