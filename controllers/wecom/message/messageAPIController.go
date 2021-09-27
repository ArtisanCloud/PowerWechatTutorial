package message

import (
  "github.com/ArtisanCloud/power-wechat/src/kernel/power"
  "github.com/gin-gonic/gin"
  "net/http"
  "power-wechat-tutorial/services"
)

// 发送应用消息
// https://work.weixin.qq.com/api/doc/90000/90135/90236
func APIMessageSendText(c *gin.Context) {

  messages := &power.HashMap{
    "touser":  "UserID1|UserID2|UserID3",
    "toparty": "PartyID1|PartyID2",
    "totag":   "TagID1 | TagID2",
    "msgtype": "text",
    "agentid": 1,
    "text": power.HashMap{
      "content": "你的快递已到，请携带工卡前往邮件中心领取。\n出发前可查看<a href=\"http://work.weixin.qq.com\">邮件中心视频实况</a>，聪明避开排队。",
    },
    "safe":                     0,
    "enable_id_trans":          0,
    "enable_duplicate_check":   0,
    "duplicate_check_interval": 1800,
  }
  res, err := services.WeComApp.Message.Send(messages)

  if err != nil {
    panic(err)
  }

  c.JSON(http.StatusOK, res)
}

func APIMessageSendImage(c *gin.Context) {
  messages := &power.HashMap{
    "touser":  "UserID1|UserID2|UserID3",
    "toparty": "PartyID1|PartyID2",
    "totag":   "TagID1 | TagID2",
    "msgtype": "image",
    "agentid": 1,
    "image": power.HashMap{
      "media_id": "MEDIA_ID",
    },
    "safe":                     0,
    "enable_duplicate_check":   0,
    "duplicate_check_interval": 1800,
  }
  res, err := services.WeComApp.Message.Send(messages)

  if err != nil {
    panic(err)
  }

  c.JSON(http.StatusOK, res)
}

func APIMessageSendVoice(c *gin.Context) {
  messages := &power.HashMap{
    "touser":  "UserID1|UserID2|UserID3",
    "toparty": "PartyID1|PartyID2",
    "totag":   "TagID1 | TagID2",
    "msgtype": "voice",
    "agentid": 1,
    "voice": power.HashMap{
      "media_id": "MEDIA_ID",
    },
    "enable_duplicate_check":   0,
    "duplicate_check_interval": 1800,
  }
  res, err := services.WeComApp.Message.Send(messages)

  if err != nil {
    panic(err)
  }

  c.JSON(http.StatusOK, res)
}

func APIMessageSendVideo(c *gin.Context) {
  messages := &power.HashMap{
    "touser":  "UserID1|UserID2|UserID3",
    "toparty": "PartyID1|PartyID2",
    "totag":   "TagID1 | TagID2",
    "msgtype": "video",
    "agentid": 1,
    "video": power.HashMap{
      "media_id":    "MEDIA_ID",
      "title":       "Title",
      "description": "Description",
    },
    "safe":                     0,
    "enable_duplicate_check":   0,
    "duplicate_check_interval": 1800,
  }
  res, err := services.WeComApp.Message.Send(messages)

  if err != nil {
    panic(err)
  }

  c.JSON(http.StatusOK, res)
}

func APIMessageSendFile(c *gin.Context) {
  messages := &power.HashMap{
    "touser":  "UserID1|UserID2|UserID3",
    "toparty": "PartyID1|PartyID2",
    "totag":   "TagID1 | TagID2",
    "msgtype": "file",
    "agentid": 1,
    "file": power.HashMap{
      "media_id": "1Yv-zXfHjSjU-7LH-GwtYqDGS-zz6w22KmWAT5COgP7o",
    },
    "safe":                     0,
    "enable_duplicate_check":   0,
    "duplicate_check_interval": 1800,
  }
  res, err := services.WeComApp.Message.Send(messages)

  if err != nil {
    panic(err)
  }

  c.JSON(http.StatusOK, res)
}

func APIMessageSendTextcard(c *gin.Context) {
  messages := &power.HashMap{
    "touser":  "UserID1|UserID2|UserID3",
    "toparty": "PartyID1 | PartyID2",
    "totag":   "TagID1 | TagID2",
    "msgtype": "textcard",
    "agentid": 1,
    "textcard": power.HashMap{
      "title":       "领奖通知",
      "description": "<div class=\"gray\">2016年9月26日</div> <div class=\"normal\">恭喜你抽中iPhone 7一台，领奖码：xxxx</div><div class=\"highlight\">请于2016年10月10日前联系行政同事领取</div>",
      "url":         "URL",
      "btntxt":      "更多",
    },
    "enable_id_trans":          0,
    "enable_duplicate_check":   0,
    "duplicate_check_interval": 1800,
  }
  res, err := services.WeComApp.Message.Send(messages)

  if err != nil {
    panic(err)
  }

  c.JSON(http.StatusOK, res)
}


func APIMessageSendNews(c *gin.Context) {

  messages := &power.HashMap{
    "touser":  "UserID1|UserID2|UserID3",
    "toparty": "PartyID1 | PartyID2",
    "totag":   "TagID1 | TagID2",
    "msgtype": "news",
    "agentid": 1,
    "news": power.HashMap{
      "articles": []power.HashMap{
        {
          "title":       "中秋节礼品领取",
          "description": "今年中秋节公司有豪礼相送",
          "url":         "URL",
          "picurl":      "http://res.mail.qq.com/node/ww/wwopenmng/images/independent/doc/test_pic_msg1.png",
          "appid":       "wx123123123123123",
          "pagepath":    "pages/index?userid=zhangsan&orderid=123123123",
        },
      },
      "enable_id_trans":          0,
      "enable_duplicate_check":   0,
      "duplicate_check_interval": 1800,
    },
  }
  res, err := services.WeComApp.Message.Send(messages)

  if err != nil {
    panic(err)
  }

  c.JSON(http.StatusOK, res)
}

func APIMessageSendMPNews(c *gin.Context) {
  messages := &power.HashMap{
    "touser":  "UserID1|UserID2|UserID3",
    "toparty": "PartyID1 | PartyID2",
    "totag":   "TagID1 | TagID2",
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
    "safe":                     0,
    "enable_id_trans":          0,
    "enable_duplicate_check":   0,
    "duplicate_check_interval": 1800,
  }
  res, err := services.WeComApp.Message.Send(messages)

  if err != nil {
    panic(err)
  }

  c.JSON(http.StatusOK, res)
}

func APIMessageSendMarkdown(c *gin.Context) {

  messages := &power.HashMap{
    "touser":  "UserID1|UserID2|UserID3",
    "toparty": "PartyID1|PartyID2",
    "totag":   "TagID1 | TagID2",
    "msgtype": "markdown",
    "agentid": 1,
    "markdown": power.HashMap{
      "content": `"您的会议室已经预定，稍后会同步到邮箱
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
    "enable_duplicate_check":   0,
    "duplicate_check_interval": 1800,
  }
  res, err := services.WeComApp.Message.Send(messages)

  if err != nil {
    panic(err)
  }

  c.JSON(http.StatusOK, res)
}


func APIMessageSendMiniProgramNotice(c *gin.Context) {
  messages := &power.HashMap{
    "touser":  "zhangsan|lisi",
    "toparty": "1|2",
    "totag":   "1|2",
    "msgtype": "miniprogram_notice",
    "miniprogram_notice": power.HashMap{
      "appid":               "wx123123123123123",
      "page":                "pages/index?userid=zhangsan&orderid=123123123",
      "title":               "会议室预订成功通知",
      "description":         "4月27日 16:16",
      "emphasis_first_item": true,
      "content_item": []power.HashMap{
        {
          "key":   "会议室",
          "value": "402",
        },
        {
          "key":   "会议地点",
          "value": "广州TIT-402会议室",
        },
        {
          "key":   "会议时间",
          "value": "2018年8月1日 09:00-09:30",
        },
        {
          "key":   "参与人员",
          "value": "周剑轩",
        },
      },
    },
    "enable_id_trans":          0,
    "enable_duplicate_check":   0,
    "duplicate_check_interval": 1800,
  }
  res, err := services.WeComApp.Message.Send(messages)

  if err != nil {
    panic(err)
  }

  c.JSON(http.StatusOK, res)
}

func APIMessageRecall(c *gin.Context) {
  msgID:=c.DefaultQuery("msgID","MSGID")
  res, err := services.WeComApp.Message.Recall(msgID)

  if err != nil {
    panic(err)
  }

  c.JSON(http.StatusOK, res)

}
