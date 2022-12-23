package message

import (
	"github.com/ArtisanCloud/PowerWeChat/v3/src/work/message/request"
	"github.com/gin-gonic/gin"
	"net/http"
	"power-wechat-tutorial/services"
)

// 发送应用消息
// https://work.weixin.qq.com/api/doc/90000/90135/90236
func APIMessageSendText(c *gin.Context) {

	messages := &request.RequestMessageSendText{
		RequestMessageSend: request.RequestMessageSend{
			ToUser:                 "UserID1|UserID2|UserID3",
			ToParty:                "PartyID1|PartyID2",
			ToTag:                  "TagID1 | TagID2",
			MsgType:                "text",
			AgentID:                1,
			Safe:                   0,
			EnableIDTrans:          0,
			EnableDuplicateCheck:   0,
			DuplicateCheckInterval: 1800,
		},
		Text: &request.RequestText{
			Content: "你的快递已到，请携带工卡前往邮件中心领取。\n出发前可查看<a href=\"http://work.weixin.qq.com\">邮件中心视频实况</a>，聪明避开排队。",
		},
	}
	res, err := services.WeComApp.Message.SendText(messages)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, res)
}

func APIMessageSendImage(c *gin.Context) {
	messages := &request.RequestMessageSendImage{
		RequestMessageSend: request.RequestMessageSend{
			ToUser:                 "UserID1|UserID2|UserID3",
			ToParty:                "PartyID1|PartyID2",
			ToTag:                  "TagID1 | TagID2",
			MsgType:                "image",
			AgentID:                1,
			Safe:                   0,
			EnableDuplicateCheck:   0,
			DuplicateCheckInterval: 1800,
		},
		Image: &request.RequestImage{
			MediaID: "MEDIA_ID",
		},
	}
	res, err := services.WeComApp.Message.SendImage(messages)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, res)
}

func APIMessageSendVoice(c *gin.Context) {
	messages := &request.RequestMessageSendVoice{
		RequestMessageSend: request.RequestMessageSend{
			ToUser:                 "UserID1|UserID2|UserID3",
			ToParty:                "PartyID1|PartyID2",
			ToTag:                  "TagID1 | TagID2",
			MsgType:                "voice",
			AgentID:                1,
			EnableDuplicateCheck:   0,
			DuplicateCheckInterval: 1800,
		},
		Voice: &request.RequestVoice{
			MediaID: "MEDIA_ID",
		},
	}
	res, err := services.WeComApp.Message.SendVoice(messages)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, res)
}

func APIMessageSendVideo(c *gin.Context) {
	messages := &request.RequestMessageSendVideo{
		RequestMessageSend: request.RequestMessageSend{
			ToUser:                 "UserID1|UserID2|UserID3",
			ToParty:                "PartyID1|PartyID2",
			ToTag:                  "TagID1 | TagID2",
			MsgType:                "video",
			AgentID:                1,
			Safe:                   0,
			EnableDuplicateCheck:   0,
			DuplicateCheckInterval: 1800,
		},
		Video: &request.RequestVideo{
			MediaID:     "MEDIA_ID",
			Title:       "Title",
			Description: "Description",
		},
	}
	res, err := services.WeComApp.Message.SendVideo(messages)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, res)
}

func APIMessageSendFile(c *gin.Context) {
	messages := &request.RequestMessageSendFile{
		RequestMessageSend: request.RequestMessageSend{
			ToUser:                 "UserID1|UserID2|UserID3",
			ToParty:                "PartyID1|PartyID2",
			ToTag:                  "TagID1 | TagID2",
			MsgType:                "file",
			AgentID:                1,
			Safe:                   0,
			EnableDuplicateCheck:   0,
			DuplicateCheckInterval: 1800,
		},
		File: &request.RequestFile{
			MediaID: "1Yv-zXfHjSjU-7LH-GwtYqDGS-zz6w22KmWAT5COgP7o",
		},
	}
	res, err := services.WeComApp.Message.SendFile(messages)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, res)
}

func APIMessageSendTextcard(c *gin.Context) {
	messages := &request.RequestMessageSendTextCard{
		RequestMessageSend: request.RequestMessageSend{
			ToUser:                 "UserID1|UserID2|UserID3",
			ToParty:                "PartyID1 | PartyID2",
			ToTag:                  "TagID1 | TagID2",
			MsgType:                "textcard",
			AgentID:                1,
			EnableIDTrans:          0,
			EnableDuplicateCheck:   0,
			DuplicateCheckInterval: 1800,
		},
		TextCard: &request.RequestTextCard{
			Title:       "领奖通知",
			Description: "<div class=\"gray\">2016年9月26日</div> <div class=\"normal\">恭喜你抽中iPhone 7一台，领奖码：xxxx</div><div class=\"highlight\">请于2016年10月10日前联系行政同事领取</div>",
			Url:         "URL",
			BtnTXT:      "更多",
		},
	}
	res, err := services.WeComApp.Message.SendTextCard(messages)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, res)
}

func APIMessageSendNews(c *gin.Context) {
	messages := &request.RequestMessageSendNews{
		RequestMessageSend: request.RequestMessageSend{
			ToUser:  "UserID1|UserID2|UserID3",
			ToParty: "PartyID1 | PartyID2",
			ToTag:   "TagID1 | TagID2",
			MsgType: "news",
			AgentID: 1,

			EnableIDTrans:          0,
			EnableDuplicateCheck:   0,
			DuplicateCheckInterval: 1800,
		},
		News: &request.RequestNews{
			Article: []*request.RequestNewsArticle{
				{
					Title:       "中秋节礼品领取",
					Description: "今年中秋节公司有豪礼相送",
					URL:         "URL",
					PicURL:      "http://res.mail.qq.com/node/ww/wwopenmng/images/independent/doc/test_pic_msg1.png",
					AppID:       "wx123123123123123",
					PagePath:    "pages/index?userid=zhangsan&orderid=123123123",
				},
			},
		},
	}
	res, err := services.WeComApp.Message.SendNews(messages)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, res)
}

func APIMessageSendMPNews(c *gin.Context) {
	messages := &request.RequestMessageSendMPNews{
		RequestMessageSend: request.RequestMessageSend{
			ToUser:                 "UserID1|UserID2|UserID3",
			ToParty:                "PartyID1 | PartyID2",
			ToTag:                  "TagID1 | TagID2",
			MsgType:                "mpnews",
			AgentID:                1,
			Safe:                   0,
			EnableIDTrans:          0,
			EnableDuplicateCheck:   0,
			DuplicateCheckInterval: 1800,
		},
		MPNews: &request.RequestMPNews{
			Article: []*request.RequestMPNewsArticle{
				{
					Title:            "Title",
					ThumbMediaID:     "MEDIA_ID",
					Author:           "Author",
					ContentSourceURL: "URL",
					Content:          "Content",
					Digest:           "Digest description",
				},
			},
		},
	}
	res, err := services.WeComApp.Message.SendMpNews(messages)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, res)
}

func APIMessageSendMarkdown(c *gin.Context) {

	messages := &request.RequestMessageSendMarkdown{
		RequestMessageSend: request.RequestMessageSend{
			ToUser:                 "UserID1|UserID2|UserID3",
			ToParty:                "PartyID1|PartyID2",
			ToTag:                  "TagID1 | TagID2",
			MsgType:                "markdown",
			AgentID:                1,
			EnableDuplicateCheck:   0,
			DuplicateCheckInterval: 1800,
		},
		Markdown: &request.RequestMarkdown{
			Content: `"您的会议室已经预定，稍后会同步到邮箱
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
      >如需修改会议信息，请点击：[修改会议信息](https://work.weixin.qq.com)"`,
		},
	}
	res, err := services.WeComApp.Message.SendMarkdown(messages)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, res)
}

func APIMessageSendMiniProgramNotice(c *gin.Context) {
	messages := &request.RequestMessageSendMiniProgramNotice{
		RequestMessageSend: request.RequestMessageSend{
			ToUser:                 "zhangsan|lisi",
			ToParty:                "1|2",
			ToTag:                  "1|2",
			MsgType:                "miniprogram_notice",
			EnableIDTrans:          0,
			EnableDuplicateCheck:   0,
			DuplicateCheckInterval: 1800,
		},

		MiniProgramNotice: &request.MiniProgramNotice{
			Appid:             "wx123123123123123",
			Page:              "pages/index?userid=zhangsan&orderid=123123123",
			Title:             "会议室预订成功通知",
			Description:       "4月27日 16:16",
			EmphasisFirstItem: true,
			ContentItem: []*request.ContentItem{
				{
					Key:   "会议室",
					Value: "402",
				},
				{
					Key:   "会议地点",
					Value: "广州TIT-402会议室",
				},
				{
					Key:   "会议时间",
					Value: "2018年8月1日 09:00-09:30",
				},
				{
					Key:   "参与人员",
					Value: "周剑轩",
				},
			},
		},
	}
	res, err := services.WeComApp.Message.SendMiniProgramNotice(messages)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, res)
}

func APIMessageSendTemplateCard(c *gin.Context) {
	messages := &request.RequestMessageSendTemplateCard{
		RequestMessageSend: request.RequestMessageSend{
			ToUser:                 "UserID1|UserID2|UserID3",
			ToParty:                "PartyID1 | PartyID2",
			ToTag:                  "TagID1 | TagID2",
			MsgType:                "template_card",
			AgentID:                1,
			EnableIDTrans:          0,
			EnableDuplicateCheck:   0,
			DuplicateCheckInterval: 1800,
		},
		TemplateCard: &request.RequestTemplateCard{
			CardType: "text_notice",
			Source: &request.TemplateCardSource{
				IconUrl:   "图片的url",
				Desc:      "企业微信",
				DescColor: 1,
			},
			ActionMenu: &request.TemplateCardActionMenu{
				Desc: "卡片副交互辅助文本说明",
				ActionList: []*request.TemplateCardActionListItem{
					{Text: "接受推送", Key: "A"},
					{Text: "不再推送", Key: "B"},
				},
			},
			TaskID: "task_id",
			MainTitle: &request.TemplateCardMainTitle{
				Title: "欢迎使用企业微信",
				Desc:  "您的好友正在邀请您加入企业微信",
			},
			QuoteArea: &request.TemplateCardQuoteArea{
				Type:      1,
				Url:       "https://work.weixin.qq.com",
				Title:     "企业微信的引用样式",
				QuoteText: "企业微信真好用呀真好用",
			},
			EmphasisContent: &request.TemplateCardEmphasisContent{
				Title: "100",
				Desc:  "核心数据",
			},
			SubTitleText: "下载企业微信还能抢红包！",
			HorizontalContentList: []*request.TemplateCardHorizontalContentListItem{
				{
					Keyname: "邀请人",
					Value:   "张三",
				},
				{
					Type:    1,
					Keyname: "企业微信官网",
					Value:   "点击访问",
					Url:     "https://work.weixin.qq.com",
				},
				{
					Type:    2,
					Keyname: "企业微信下载",
					Value:   "企业微信.apk",
					MediaID: "文件的media_id",
				},
				{
					Type:    3,
					Keyname: "员工信息",
					Value:   "点击查看",
					UserID:  "zhangsan",
				},
			},
			JumpList: []*request.TemplateCardJumpListItem{
				{
					Type:  1,
					Title: "企业微信官网",
					Url:   "https://work.weixin.qq.com",
				},
				{
					Type:     2,
					Title:    "跳转小程序",
					AppID:    "小程序的appid",
					Pagepath: "/index.html",
				},
			},
			CardAction: &request.TemplateCardAction{
				Type:     2,
				Url:      "https://work.weixin.qq.com",
				AppID:    "小程序的appid",
				Pagepath: "/index.html",
			},
		},
	}

	res, err := services.WeComApp.Message.SendTemplateCard(messages)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, res)
}

func APIMessageRecall(c *gin.Context) {
	msgID := c.DefaultQuery("msgID", "MSGID")
	res, err := services.WeComApp.Message.Recall(msgID)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, res)

}
