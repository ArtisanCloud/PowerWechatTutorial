package external_contact

import (
	"net/http"
	"power-wechat-tutorial/services"
	"strconv"

	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/power"
	request2 "github.com/ArtisanCloud/PowerWeChat/v3/src/work/externalContact/groupWelcomeTemplate/request"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/work/externalContact/messageTemplate/request"
	"github.com/gin-gonic/gin"
)

func APIExternalContactAddMsgTemplate(c *gin.Context) {

	options := &request.RequestAddMsgTemplate{
		ChatType: "single",
		ExternalUserID: []string{
			"woAJ2GCAAAXtWyujaWJHDDGi0mACAAAA",
			"wmqfasd1e1927831123109rBAAAA",
		},
		Sender: "matrix-x",
	}

	res, err := services.WeComContactApp.ExternalContactMessageTemplate.AddMsgTemplate(c.Request.Context(), options)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, res)
}

func APIExternalContactGetGroupMsgListV2(c *gin.Context) {
	options := &request.RequestGetGroupMsgListV2{
		ChatType:   "single",
		StartTime:  1605171726,
		EndTime:    1605172726,
		Creator:    power.String("matrix-x"),
		FilterType: power.Int(1),
		Limit:      50,
		Cursor:     "CURSOR",
	}
	res, err := services.WeComContactApp.ExternalContactMessageTemplate.GetGroupMsgListV2(c.Request.Context(), options)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, res)
}

func APIExternalContactGetGroupMsgTask(c *gin.Context) {

	msgID := c.DefaultQuery("msgID", "wrOgQhDgAAMYQiS5ol9G7gK9JVAAAA")
	limit := 1000
	cursor := c.DefaultQuery("cursor", "wrOgQhDgAAMYQiS5ol9G7gK9JVAAAA")

	res, err := services.WeComContactApp.ExternalContactMessageTemplate.GetGroupMsgTask(c.Request.Context(), msgID, limit, cursor)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, res)
}

func APIExternalContactGetGroupMsgSendResult(c *gin.Context) {

	msgID := c.DefaultQuery("msgID", "wrOgQhDgAAMYQiS5ol9G7gK9JVAAAA")
	userID := c.DefaultQuery("userID", "wrOgQhDgAAMYQiS5ol9G7gK9JVAAAA")
	limit := 1000
	cursor := c.DefaultQuery("cursor", "wrOgQhDgAAMYQiS5ol9G7gK9JVAAAA")

	res, err := services.WeComContactApp.ExternalContactMessageTemplate.GetGroupMsgSendResult(c.Request.Context(), msgID, userID, limit, cursor)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, res)
}

func APIExternalContactSendWelcomeMsg(c *gin.Context) {

	options := &request.RequestSendWelcomeMsg{
		WelcomeCode: c.DefaultQuery("welcomeCode", "CALLBACK_CODE"),
		Text: &request.TextOfMessage{
			Content: "文本消息内容",
		},
		//Attachments: []*power.HashMap{
		//  &power.HashMap{
		//    "msgtype": "image",
		//    "image": power.HashMap{
		//      "media_id": "MEDIA_ID",
		//      "pic_url":  "http://p.qpic.cn/pic_wework/3474110808/7a6344sdadfwehe42060/0",
		//    },
		//  }, &power.HashMap{
		//    "msgtype": "link",
		//    "link": power.HashMap{
		//      "title":  "消息标题",
		//      "picurl": "https://example.pic.com/path",
		//      "desc":   "消息描述",
		//      "url":    "https://example.link.com/path",
		//    },
		//  }, &power.HashMap{
		//    "msgtype": "miniprogram",
		//    "miniprogram": power.HashMap{
		//      "title":        "消息标题",
		//      "pic_media_id": "MEDIA_ID",
		//      "appid":        "wx8bd80126147dfAAA",
		//      "page":         "/path/index.html",
		//    },
		//  }, &power.HashMap{
		//    "msgtype": "video",
		//    "video": power.HashMap{
		//      "media_id": "MEDIA_ID",
		//    },
		//  }, &power.HashMap{
		//    "msgtype": "file",
		//    "file": power.HashMap{
		//      "media_id": "MEDIA_ID",
		//    },
		//  },
		//},
	}
	res, err := services.WeComContactApp.ExternalContactMessageTemplate.SendWelcomeMsg(c.Request.Context(), options)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, res)
}

func APIExternalContactGroupWelcomeTemplateAdd(c *gin.Context) {

	options := &request2.RequestGroupWelcomeTemplateAdd{
		Text: &request.TextOfMessage{
			Content: "亲爱的%NICKNAME%用户，你好",
		},
		Image: &request.Image{
			MediaID: "MEDIA_ID",
			PicURL:  "http://p.qpic.cn/pic_wework/3474110808/7a6344sdadfwehe42060/0",
		},
		Link: &request.Link{
			Title:  "消息标题",
			PicURL: "https://example.pic.com/path",
			Desc:   "消息描述",
			URL:    "https://example.link.com/path",
		},
		MiniProgram: &request.MiniProgram{
			Title:      "消息标题",
			PicMediaID: "MEDIA_ID",
			AppID:      "wx8bd80126147dfAAA",
			Page:       "/path/index",
		},
		File: &request.File{
			MediaID: "1Yv-zXfHjSjU-7LH-GwtYqDGS-zz6w22KmWAT5COgP7o",
		},
		Video: &request.Video{
			MediaID: "MEDIA_ID",
		},
		AgentID: 1000014,
		Notify:  1,
	}

	res, err := services.WeComContactApp.ExternalContactGroupWelcomeTemplate.AddGroupWelcomeTemplate(c.Request.Context(), options)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, res)
}

func APIExternalContactGroupWelcomeTemplateEdit(c *gin.Context) {
	options := &request2.RequestGroupWelcomeTemplateEdit{
		TemplateID: c.DefaultQuery("templateID", "TEMPLATEID"),
		Text: &request.TextOfMessage{
			Content: "亲爱的%NICKNAME%用户，你好",
		},
		Image: &request.Image{
			MediaID: "MEDIA_ID",
			PicURL:  "http://p.qpic.cn/pic_wework/3474110808/7a6344sdadfwehe42060/0",
		},
		Link: &request.Link{
			Title:  "消息标题",
			PicURL: "https://example.pic.com/path",
			Desc:   "消息描述",
			URL:    "https://example.link.com/path",
		},
		MiniProgram: &request.MiniProgram{
			Title:      "消息标题",
			PicMediaID: "MEDIA_ID",
			AppID:      "wx8bd80126147dfAAA",
			Page:       "/path/index",
		},
		File: &request.File{
			MediaID: "1Yv-zXfHjSjU-7LH-GwtYqDGS-zz6w22KmWAT5COgP7o",
		},
		Video: &request.Video{
			MediaID: "MEDIA_ID",
		},
		AgentID: 1000014,
		Notify:  1,
	}

	res, err := services.WeComContactApp.ExternalContactGroupWelcomeTemplate.EditGroupWelcomeTemplate(c.Request.Context(), options)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, res)
}

func APIExternalContactGroupWelcomeTemplateGet(c *gin.Context) {
	templateID := c.DefaultQuery("templateID", "TEMPLATEID")
	res, err := services.WeComContactApp.ExternalContactGroupWelcomeTemplate.GetGroupWelcomeTemplate(c.Request.Context(), templateID)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, res)
}

func APIExternalContactGroupWelcomeTemplateDel(c *gin.Context) {
	templateID := c.DefaultQuery("templateID", "TEMPLATEID")
	agentId := c.DefaultQuery("agentID", "TEMPLATEID")
	agentID, _ := strconv.Atoi(agentId)

	res, err := services.WeComContactApp.ExternalContactGroupWelcomeTemplate.DelGroupWelcomeTemplate(c.Request.Context(), templateID, int64(agentID))

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, res)
}
