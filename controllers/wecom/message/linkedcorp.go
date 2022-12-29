package message

import (
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/power"
	"github.com/gin-gonic/gin"
	"net/http"
	"power-wechat-tutorial/services"
)

// 应用推送消息
// https://work.weixin.qq.com/api/doc/90000/90135/90250
func APIAppChatLinkedCorpSendText(c *gin.Context) {

	options := &power.HashMap{
		"touser":  []string{c.DefaultQuery("userID", "matrix-x")},
		"toparty": []string{c.DefaultQuery("partyID", "PARTYID")},
		"totag":   []string{},
		"toall":   0,
		"msgtype": "text",
		"agentid": 1,
		"text": power.HashMap{
			"content": "你的快递已到，请携带工卡前往邮件中心领取。\n出发前可查看<a href=\"http://work.weixin.qq.com\">邮件中心视频实况</a>，聪明避开排队。",
		},
		"safe": 0,
	}

	res, err := services.WeComApp.MessageLinkedCorp.Send(c.Request.Context(), options)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, res)
}

func APIAppChatLinkedCorpSendImage(c *gin.Context) {

	options := &power.HashMap{
		"touser":  []string{c.DefaultQuery("userID", "matrix-x")},
		"toparty": []string{c.DefaultQuery("partyID", "PARTYID")},
		"totag":   []string{},
		"toall":   0,
		"msgtype": "image",
		"agentid": 1,
		"image": power.HashMap{
			"media_id": "MEDIA_ID",
		},
		"safe": 0,
	}

	res, err := services.WeComApp.MessageLinkedCorp.Send(c.Request.Context(), options)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, res)
}

func APIAppChatLinkedCorpSendVoice(c *gin.Context) {
	options := &power.HashMap{
		"touser":  []string{c.DefaultQuery("userID", "matrix-x")},
		"toparty": []string{c.DefaultQuery("partyID", "PARTYID")},
		"totag":   []string{},
		"toall":   0,
		"msgtype": "voice",
		"agentid": 1,
		"voice": power.HashMap{
			"media_id": "MEDIA_ID",
		},
	}

	res, err := services.WeComApp.MessageLinkedCorp.Send(c.Request.Context(), options)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, res)
}

func APIAppChatLinkedCorpSendVideo(c *gin.Context) {
	options := &power.HashMap{
		"touser":  []string{c.DefaultQuery("userID", "matrix-x")},
		"toparty": []string{c.DefaultQuery("partyID", "PARTYID")},
		"totag":   []string{},
		"toall":   0,
		"msgtype": "video",
		"agentid": 1,
		"video": power.HashMap{
			"media_id":    "MEDIA_ID",
			"title":       "Title",
			"description": "Description",
		},
		"safe": 0,
	}

	res, err := services.WeComApp.MessageLinkedCorp.Send(c.Request.Context(), options)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, res)
}

func APIAppChatLinkedCorpSendFile(c *gin.Context) {
	options := &power.HashMap{
		"touser":  []string{c.DefaultQuery("userID", "matrix-x")},
		"toparty": []string{c.DefaultQuery("partyID", "PARTYID")},
		"totag":   []string{},
		"toall":   0,
		"msgtype": "file",
		"agentid": 1,
		"file": power.HashMap{
			"media_id": "1Yv-zXfHjSjU-7LH-GwtYqDGS-zz6w22KmWAT5COgP7o",
		},
		"safe": 0,
	}

	res, err := services.WeComApp.MessageLinkedCorp.Send(c.Request.Context(), options)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, res)
}

func APIAppChatLinkedCorpSendTextcard(c *gin.Context) {
	options := &power.HashMap{
		"touser":  []string{c.DefaultQuery("userID", "matrix-x")},
		"toparty": []string{c.DefaultQuery("partyID", "PARTYID")},
		"totag":   []string{},
		"toall":   0,
		"msgtype": "textcard",
		"agentid": 1,
		"textcard": power.HashMap{
			"title":       "领奖通知",
			"description": "<div class=\"gray\">2016年9月26日</div> <div class=\"normal\">恭喜你抽中iPhone 7一台，领奖码：xxxx</div><div class=\"highlight\">请于2016年10月10日前联系行政同事领取</div>",
			"url":         "URL",
			"btntxt":      "更多",
		},
	}

	res, err := services.WeComApp.MessageLinkedCorp.Send(c.Request.Context(), options)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, res)
}

func APIAppChatLinkedCorpSendNews(c *gin.Context) {
	options := &power.HashMap{
		"touser":  []string{c.DefaultQuery("userID", "matrix-x")},
		"toparty": []string{c.DefaultQuery("partyID", "PARTYID")},
		"totag":   []string{},
		"toall":   0,
		"msgtype": "news",
		"agentid": 1,
		"news": power.HashMap{
			"articles": []power.HashMap{
				{
					"title":       "中秋节礼品领取",
					"description": "今年中秋节公司有豪礼相送",
					"url":         "URL",
					"picurl":      "http://res.mail.qq.com/node/ww/wwopenmng/images/independent/doc/test_pic_msg1.png",
					"btntxt":      "更多",
				},
			},
		},
	}

	res, err := services.WeComApp.MessageLinkedCorp.Send(c.Request.Context(), options)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, res)
}

func APIAppChatLinkedCorpSendMPNews(c *gin.Context) {
	options := &power.HashMap{
		"touser":  []string{c.DefaultQuery("userID", "matrix-x")},
		"toparty": []string{c.DefaultQuery("partyID", "PARTYID")},
		"totag":   []string{},
		"toall":   0,
		"msgtype": "mpnews",
		"agentid": 1,
		"mpnews": power.HashMap{
			"articles": []power.HashMap{
				{
					"title":              "Title",
					"thumb_media_id":     "MEDIA_ID",
					"author":             "Author",
					"content_source_url": "URL",
					"content":            "Content",
					"digest":             "Digest description",
				},
			},
		},
		"safe": 0,
	}

	res, err := services.WeComApp.MessageLinkedCorp.Send(c.Request.Context(), options)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, res)
}

func APIAppChatLinkedCorpSendMarkdown(c *gin.Context) {
	options := &power.HashMap{
		"touser":  []string{c.DefaultQuery("userID", "matrix-x")},
		"toparty": []string{c.DefaultQuery("partyID", "PARTYID")},
		"totag":   []string{},
		"toall":   0,
		"msgtype": "markdown",
		"agentid": 1,
		"markdown": power.HashMap{
			"content": `"您的会议室已经预定，稍后会同步到\'邮箱\'
      >**事项详情**
      >事　项：<font color=\"info\">开会</font>
      >组织者：@miglioguan
      >参与者：@miglioguan、@kunliu、@jamdeezhou、@kanexiong、@kisonwang
      >
      >会议室：<font color=\"info\">广州TIT 1楼 301</font>
      >日　期：<font color=\"warning\">2018年5月18日</font>
      >时　间：<font color=\"comment\">上午9:00-11:00</font>
      >
      >请准时参加会议。
      >
      >如需修改会议信息，请点击：[修改会议信息](https://work.weixin.qq.com)" `,
		},
	}

	res, err := services.WeComApp.MessageLinkedCorp.Send(c.Request.Context(), options)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, res)
}

func APIAppChatLinkedCorpSendMiniProgramNotice(c *gin.Context) {
	options := &power.HashMap{
		"touser":  []string{c.DefaultQuery("userID", "matrix-x")},
		"toparty": []string{c.DefaultQuery("partyID", "PARTYID")},
		"totag":   []string{},
		"msgtype": "miniprogram_notice",
		"miniprogram_notice": power.HashMap{
			"appid":               "wx123123123123123",
			"page":                "pages/index?userid=zhangsan&orderid=123123123",
			"title":               "会议室预订成功通知",
			"description":         "4月27日 16:16",
			"emphasis_first_item": true,
			"content_item": []power.HashMap{
				power.HashMap{
					"key":   "会议室",
					"value": "402",
				},
				power.HashMap{
					"key":   "会议地点",
					"value": "广州TIT-402会议室",
				},
				power.HashMap{
					"key":   "会议时间",
					"value": "2018年8月1日 09:00-09:30",
				},
				power.HashMap{
					"key":   "参与人员",
					"value": "周剑轩",
				},
			},
		},
	}

	res, err := services.WeComApp.MessageLinkedCorp.Send(c.Request.Context(), options)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, res)
}
