package message

import (
  "github.com/ArtisanCloud/power-wechat/src/kernel/power"
  "github.com/gin-gonic/gin"
  "net/http"
  "power-wechat-tutorial/services"
)

func APIMessageSendTemplateCardTextNotice(c *gin.Context) {
  messages := &power.HashMap{
    "touser":  "UserID1|UserID2|UserID3",
    "toparty": "PartyID1 | PartyID2",
    "totag":   "TagID1 | TagID2",
    "msgtype": "template_card",
    "agentid": 1,
    "template_card": power.HashMap{
      "card_type": "text_notice",
      "source": power.HashMap{
        "icon_url": "图片的url",
        "desc":     "企业微信",
      },
      "main_title": power.HashMap{
        "title": "欢迎使用企业微信",
        "desc":  "您的好友正在邀请您加入企业微信",
      },
      "emphasis_content": power.HashMap{
        "title": "100",
        "desc":  "核心数据",
      },
      "sub_title_text": "下载企业微信还能抢红包！",
      "horizontal_content_list": []power.HashMap{
        {
          "keyname": "邀请人",
          "value":   "张三",
        },
        {
          "type":    1,
          "keyname": "企业微信官网",
          "value":   "点击访问",
          "url":     "https://work.weixin.qq.com",
        },
        {
          "type":     2,
          "keyname":  "企业微信下载",
          "value":    "企业微信.apk",
          "media_id": "文件的media_id",
        },
      },
      "jump_list": []power.HashMap{
        {
          "type":  1,
          "title": "企业微信官网",
          "url":   "https://work.weixin.qq.com",
        },
        {
          "type":     2,
          "title":    "跳转小程序",
          "appid":    "小程序的appid",
          "pagepath": "/index.html",
        },
      },
      "card_action": power.HashMap{
        "type":     2,
        "url":      "https://work.weixin.qq.com",
        "appid":    "小程序的appid",
        "pagepath": "/index.html",
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

func APIMessageSendTemplateCardNewsNotice(c *gin.Context) {
  messages := &power.HashMap{
    "touser":  "UserID1|UserID2|UserID3",
    "toparty": "PartyID1 | PartyID2",
    "totag":   "TagID1 | TagID2",
    "msgtype": "template_card",
    "agentid": 1,
    "template_card": power.HashMap{
      "card_type": "news_notice",
      "source": power.HashMap{
        "icon_url": "图片的url",
        "desc":     "企业微信",
      },
      "main_title": power.HashMap{
        "title": "欢迎使用企业微信",
        "desc":  "您的好友正在邀请您加入企业微信",
      },
      "card_image": power.HashMap{
        "url":          "图片的url",
        "aspect_ratio": 1.3,
      },
      "vertical_content_list": []power.HashMap{
        {
          "title": "惊喜红包等你来拿",
          "desc":  "下载企业微信还能抢红包！",
        },
      },
      "horizontal_content_list": []power.HashMap{
        {
          "keyname": "邀请人",
          "value":   "张三",
        },
        {
          "type":    1,
          "keyname": "企业微信官网",
          "value":   "点击访问",
          "url":     "https://work.weixin.qq.com",
        },
        {
          "type":     2,
          "keyname":  "企业微信下载",
          "value":    "企业微信.apk",
          "media_id": "文件的media_id",
        },
      },
      "jump_list": []power.HashMap{
        {
          "type":  1,
          "title": "企业微信官网",
          "url":   "https://work.weixin.qq.com",
        },
        {
          "type":     2,
          "title":    "跳转小程序",
          "appid":    "小程序的appid",
          "pagepath": "/index.html",
        },
      },
      "card_action": power.HashMap{
        "type":     2,
        "url":      "https://work.weixin.qq.com",
        "appid":    "小程序的appid",
        "pagepath": "/index.html",
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

func APIMessageSendTemplateCardButtonInteraction(c *gin.Context) {

}

func APIMessageSendTemplateCardVoteInteraction(c *gin.Context) {

}

func APIMessageSendTemplateCardMultipleInteraction(c *gin.Context) {

}

func APIMessageUpdateButton(c *gin.Context) {

}

func APIMessageUpdateTemplateCardTextNotice(c *gin.Context) {

}

func APIMessageUpdateTemplateCardNewsNotice(c *gin.Context) {

}

func APIMessageUpdateTemplateCardButtonInteraction(c *gin.Context) {

}

func APIMessageUpdateTemplateCardVoteInteraction(c *gin.Context) {

}

func APIMessageUpdateTemplateCardMultipleInteraction(c *gin.Context) {

}
