package externalContact

import (
  "github.com/ArtisanCloud/power-wechat/src/kernel/power"
  request2 "github.com/ArtisanCloud/power-wechat/src/work/externalContact/groupWelcomeTemplate/request"
  "github.com/ArtisanCloud/power-wechat/src/work/externalContact/messageTemplate/request"
  "github.com/gin-gonic/gin"
  "net/http"
  "power-wechat-tutorial/services"
  "strconv"
)

func APIExternalContactAddMsgTemplate(c *gin.Context) {

  options := &request.RequestAddMsgTemplate{
    ChatType: "single",
    ExternalUserID: []string{
      "woAJ2GCAAAXtWyujaWJHDDGi0mACAAAA",
      "wmqfasd1e1927831123109rBAAAA",
    },
    Sender: "matrix-x",
    Text: request.TextOfMessage{
      "文本消息内容",
    },
    Attachments: []*power.HashMap{
      {
        "msgtype": "image",
        "image": power.HashMap{
          "media_id": "MEDIA_ID",
          "pic_url":  "http://p.qpic.cn/pic_wework/3474110808/7a6344sdadfwehe42060/0",
        },
      }, {
        "msgtype": "link",
        "link": power.HashMap{
          "title":  "消息标题",
          "picurl": "https://example.pic.com/path",
          "desc":   "消息描述",
          "url":    "https://example.link.com/path",
        },
      }, {
        "msgtype": "miniprogram",
        "miniprogram": power.HashMap{
          "title":        "消息标题",
          "pic_media_id": "MEDIA_ID",
          "appid":        "wx8bd80126147dfAAA",
          "page":         "/path/index.html",
        },
      }, {
        "msgtype": "video",
        "video": power.HashMap{
          "media_id": "MEDIA_ID",
        },
      }, {
        "msgtype": "file",
        "file": power.HashMap{
          "media_id": "MEDIA_ID",
        },
      },
    },
  }

  res, err := services.WeComContactApp.ExternalContactMessageTemplate.AddMsgTemplate(options)

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
    Creator:    "matrix-x",
    FilterType: 1,
    Limit:      50,
    Cursor:     "CURSOR",
  }
  res, err := services.WeComContactApp.ExternalContactMessageTemplate.GetGroupMsgListV2(options)

  if err != nil {
    panic(err)
  }

  c.JSON(http.StatusOK, res)
}

func APIExternalContactGetGroupMsgTask(c *gin.Context) {

  msgID := c.DefaultQuery("msgID", "wrOgQhDgAAMYQiS5ol9G7gK9JVAAAA")
  limit := 1000
  cursor := c.DefaultQuery("cursor", "wrOgQhDgAAMYQiS5ol9G7gK9JVAAAA")

  res, err := services.WeComContactApp.ExternalContactMessageTemplate.GetGroupMsgTask(msgID, limit, cursor)

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

  res, err := services.WeComContactApp.ExternalContactMessageTemplate.GetGroupMsgSendResult(msgID, userID, limit, cursor)

  if err != nil {
    panic(err)
  }

  c.JSON(http.StatusOK, res)
}

func APIExternalContactSendWelcomeMsg(c *gin.Context) {

  options := &request.RequestSendWelcomeMsg{
    WelcomeCode: c.DefaultQuery("welcomeCode", "CALLBACK_CODE"),
    Text: &request.TextOfMessage{
      "文本消息内容",
    },
    Attachments: []*power.HashMap{
      &power.HashMap{
        "msgtype": "image",
        "image": power.HashMap{
          "media_id": "MEDIA_ID",
          "pic_url":  "http://p.qpic.cn/pic_wework/3474110808/7a6344sdadfwehe42060/0",
        },
      }, &power.HashMap{
        "msgtype": "link",
        "link": power.HashMap{
          "title":  "消息标题",
          "picurl": "https://example.pic.com/path",
          "desc":   "消息描述",
          "url":    "https://example.link.com/path",
        },
      }, &power.HashMap{
        "msgtype": "miniprogram",
        "miniprogram": power.HashMap{
          "title":        "消息标题",
          "pic_media_id": "MEDIA_ID",
          "appid":        "wx8bd80126147dfAAA",
          "page":         "/path/index.html",
        },
      }, &power.HashMap{
        "msgtype": "video",
        "video": power.HashMap{
          "media_id": "MEDIA_ID",
        },
      }, &power.HashMap{
        "msgtype": "file",
        "file":
        power.HashMap{
          "media_id": "MEDIA_ID",
        },
      },
    },
  }
  res, err := services.WeComContactApp.ExternalContactMessageTemplate.SendWelcomeMsg(options)

  if err != nil {
    panic(err)
  }

  c.JSON(http.StatusOK, res)
}

func APIExternalContactGroupWelcomeTemplateAdd(c *gin.Context) {

  options := &request2.RequestGroupWelcomeTemplateAdd{
    Text: &request.TextOfMessage{
      "亲爱的%NICKNAME%用户，你好",
    },
    Image: &request.Image{
      "MEDIA_ID",
      "http://p.qpic.cn/pic_wework/3474110808/7a6344sdadfwehe42060/0",
    },
    Link: &request.Link{
      "消息标题",
      "https://example.pic.com/path",
      "消息描述",
      "https://example.link.com/path",
    },
    MiniProgram: &request.MiniProgram{
      "消息标题",
      "MEDIA_ID",
      "wx8bd80126147dfAAA",
      "/path/index",
    },
    File: &request.File{
      "1Yv-zXfHjSjU-7LH-GwtYqDGS-zz6w22KmWAT5COgP7o",
    },
    Video: &request.Video{
      "MEDIA_ID",
    },
    AgentID: 1000014,
    Notify:  1,
  }

  res, err := services.WeComContactApp.ExternalContactGroupWelcomeTemplate.AddGroupWelcomeTemplate(options)

  if err != nil {
    panic(err)
  }

  c.JSON(http.StatusOK, res)
}

func APIExternalContactGroupWelcomeTemplateEdit(c *gin.Context) {
  options := &request2.RequestGroupWelcomeTemplateEdit{
    TemplateID: c.DefaultQuery("templateID", "TEMPLATEID"),
    Text: &request.TextOfMessage{
      "亲爱的%NICKNAME%用户，你好",
    },
    Image: &request.Image{
      "MEDIA_ID",
      "http://p.qpic.cn/pic_wework/3474110808/7a6344sdadfwehe42060/0",
    },
    Link: &request.Link{
      "消息标题",
      "https://example.pic.com/path",
      "消息描述",
      "https://example.link.com/path",
    },
    MiniProgram: &request.MiniProgram{
      "消息标题",
      "MEDIA_ID",
      "wx8bd80126147dfAAA",
      "/path/index",
    },
    File: &request.File{
      "1Yv-zXfHjSjU-7LH-GwtYqDGS-zz6w22KmWAT5COgP7o",
    },
    Video: &request.Video{
      "MEDIA_ID",
    },
    AgentID: 1000014,
    Notify:  1,
  }

  res, err := services.WeComContactApp.ExternalContactGroupWelcomeTemplate.EditGroupWelcomeTemplate(options)

  if err != nil {
    panic(err)
  }

  c.JSON(http.StatusOK, res)
}

func APIExternalContactGroupWelcomeTemplateGet(c *gin.Context) {
  templateID := c.DefaultQuery("templateID", "TEMPLATEID")
  res, err := services.WeComContactApp.ExternalContactGroupWelcomeTemplate.GetGroupWelcomeTemplate(templateID)

  if err != nil {
    panic(err)
  }

  c.JSON(http.StatusOK, res)
}

func APIExternalContactGroupWelcomeTemplateDel(c *gin.Context) {
  templateID := c.DefaultQuery("templateID", "TEMPLATEID")
  agentId := c.DefaultQuery("agentID", "TEMPLATEID")
  agentID, _ := strconv.Atoi(agentId)

  res, err := services.WeComContactApp.ExternalContactGroupWelcomeTemplate.DelGroupWelcomeTemplate(templateID, int64(agentID))

  if err != nil {
    panic(err)
  }

  c.JSON(http.StatusOK, res)
}
