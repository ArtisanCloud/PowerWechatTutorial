package message

import (
  "github.com/ArtisanCloud/power-wechat/src/kernel/power"
  "github.com/ArtisanCloud/power-wechat/src/work/message/appChat/request"
  "github.com/gin-gonic/gin"
  "net/http"
  "power-wechat-tutorial/services"
)

// 创建群聊会话
// https://work.weixin.qq.com/api/doc/90000/90135/90245
func APIAppChatCreate(c *gin.Context) {

  options := &request.RequestAppChatCreate{
    Name:     c.DefaultQuery("name", "NAME"),
    Owner:    c.DefaultQuery("ownerUserID", "USERID"),
    UserList: []string{c.DefaultQuery("userID", "matrix-x")},
    ChatID:   c.DefaultQuery("chatID", "CHATID"),
  }

  res, err := services.WeComApp.MessageAppChat.Create(options)

  if err != nil {
    panic(err)
  }

  c.JSON(http.StatusOK, res)

}

// 修改群聊会话
// https://work.weixin.qq.com/api/doc/90000/90135/90246
func APIAppChatUpdate(c *gin.Context) {

  options := &request.RequestAppChatUpdate{
    ChatID:      c.DefaultQuery("chatID", "CHATID"),
    Name:        c.DefaultQuery("name", "NAME"),
    Owner:       c.DefaultQuery("ownerUserID", "USERID"),
    AddUserList: []string{c.DefaultQuery("userID", "matrix-x")},
    DelUserList: []string{c.DefaultQuery("userID", "walle")},
  }

  res, err := services.WeComApp.MessageAppChat.Update(options)

  if err != nil {
    panic(err)
  }

  c.JSON(http.StatusOK, res)
}

// 获取群聊会话
// https://work.weixin.qq.com/api/doc/90000/90135/90247
func APIAppChatGet(c *gin.Context) {
  chatID := c.DefaultQuery("chatID", "CHATID")
  res, err := services.WeComApp.MessageAppChat.Get(chatID)

  if err != nil {
    panic(err)
  }

  c.JSON(http.StatusOK, res)
}

// 应用推送消息
// https://work.weixin.qq.com/api/doc/90000/90135/90248
func APIAppChatSendText(c *gin.Context) {

  options := &power.HashMap{
    "chatid":  "CHATID",
    "msgtype": "text",
    "text": power.HashMap{
      "content": "你的快递已到\n请携带工卡前往邮件中心领取",
    },
    "safe": 0,
  }

  res, err := services.WeComApp.MessageAppChat.Send(options)

  if err != nil {
    panic(err)
  }

  c.JSON(http.StatusOK, res)
}

func APIAppChatSendImage(c *gin.Context) {

  options := &power.HashMap{
    "chatid":  "CHATID",
    "msgtype": "image",
    "image": power.HashMap{
      "media_id": "MEDIAID",
    },
    "safe": 0,
  }

  res, err := services.WeComApp.MessageAppChat.Send(options)

  if err != nil {
    panic(err)
  }

  c.JSON(http.StatusOK, res)
}

func APIAppChatSendVoice(c *gin.Context) {
  options := &power.HashMap{
    "chatid":  "CHATID",
    "msgtype": "voice",
    "voice": power.HashMap{
      "media_id": "MEDIA_ID",
    },
  }

  res, err := services.WeComApp.MessageAppChat.Send(options)

  if err != nil {
    panic(err)
  }

  c.JSON(http.StatusOK, res)
}

func APIAppChatSendVideo(c *gin.Context) {

  options := &power.HashMap{
    "chatid":  "CHATID",
    "msgtype": "video",
    "video": power.HashMap{
      "media_id":    "MEDIA_ID",
      "description": "Description",
      "title":       "Title",
    },
    "safe": 0,
  }

  res, err := services.WeComApp.MessageAppChat.Send(options)

  if err != nil {
    panic(err)
  }

  c.JSON(http.StatusOK, res)
}

func APIAppChatSendFile(c *gin.Context) {
  options := &power.HashMap{
    "chatid":  "CHATID",
    "msgtype": "file",
    "file": power.HashMap{
      "media_id": "1Yv-zXfHjSjU-7LH-GwtYqDGS-zz6w22KmWAT5COgP7o",
    },
    "safe": 0,
  }

  res, err := services.WeComApp.MessageAppChat.Send(options)

  if err != nil {
    panic(err)
  }

  c.JSON(http.StatusOK, res)
}

func APIAppChatSendTextcard(c *gin.Context) {
  options := &power.HashMap{
    "chatid":  "CHATID",
    "msgtype": "textcard",
    "textcard": power.HashMap{
      "title":       "领奖通知",
      "description": "<div class=\"gray\">2016年9月26日</div> <div class=\"normal\"> 恭喜你抽中iPhone 7一台，领奖码:520258</div><div class=\"highlight\">请于2016年10月10日前联系行 政同事领取</div>",
      "url":         "https://work.weixin.qq.com/",
      "btntxt":      "更多",
    },
    "safe": 0,
  }

  res, err := services.WeComApp.MessageAppChat.Send(options)

  if err != nil {
    panic(err)
  }

  c.JSON(http.StatusOK, res)
}

func APIAppChatSendNews(c *gin.Context) {
  options := &power.HashMap{
    "chatid":  "CHATID",
    "msgtype": "news",
    "news": power.HashMap{
      "articles":
      []power.HashMap{
        {
          "title":       "中秋节礼品领取",
          "description": "今年中秋节公司有豪礼相送",
          "url":         "https://work.weixin.qq.com/",
          "picurl":      "http://res.mail.qq.com/node/ww/wwopenmng/images/independent/doc/test_pic_msg1.png",
        },
      },
    },
    "safe": 0,
  }

  res, err := services.WeComApp.MessageAppChat.Send(options)

  if err != nil {
    panic(err)
  }

  c.JSON(http.StatusOK, res)
}

func APIAppChatSendMPNews(c *gin.Context) {
  options := &power.HashMap{
    "chatid":  "CHATID",
    "msgtype": "mpnews",
    "mpnews": power.HashMap{
      "articles": []power.HashMap{
        {
          "title":              "地球一小时",
          "thumb_media_id":     "biz_get(image)",
          "author":             "Author",
          "content_source_url": "https://work.weixin.qq.com",
          "content":            "3月24日20:30-21:30 \n办公区将关闭照明一小时，请各部门同事相互转告",
          "digest":             "3月24日20:30-21:30 \n办公区将关闭照明一小时",
        },
      },
    },
    "safe": 0,
  }

  res, err := services.WeComApp.MessageAppChat.Send(options)

  if err != nil {
    panic(err)
  }

  c.JSON(http.StatusOK, res)
}

func APIAppChatSendMarkdown(c *gin.Context) {
  options := &power.HashMap{
    "chatid":  "CHATID",
    "msgtype": "markdown",
    "markdown": power.HashMap{
      "content": `"您的会议室已经预定，稍后会同步到\", 邮箱\" 
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

  res, err := services.WeComApp.MessageAppChat.Send(options)

  if err != nil {
    panic(err)
  }

  c.JSON(http.StatusOK, res)
}
