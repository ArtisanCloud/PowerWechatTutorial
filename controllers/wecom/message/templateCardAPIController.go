package message

import (
  "github.com/ArtisanCloud/PowerWeChat/src/kernel/power"
  "github.com/gin-gonic/gin"
  "net/http"
  "power-wechat-tutorial/services"
)

// 发送应用消息
// https://work.weixin.qq.com/api/doc/90000/90135/90236
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

  messages := &power.HashMap{
    "touser":  "UserID1|UserID2|UserID3",
    "toparty": "PartyID1 | PartyID2",
    "totag":   "TagID1 | TagID2",
    "msgtype": "template_card",
    "agentid": 1,
    "template_card": power.HashMap{
      "card_type": "button_interaction",
      "source": power.HashMap{
        "icon_url": "图片的url",
        "desc":     "企业微信",
      },
      "main_title": power.HashMap{
        "title": "欢迎使用企业微信",
        "desc":  "您的好友正在邀请您加入企业微信",
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
      "card_action": power.HashMap{
        "type":     2,
        "url":      "https://work.weixin.qq.com",
        "appid":    "小程序的appid",
        "pagepath": "/index.html",
      },
      "task_id": "task_id",
      "button_list": []power.HashMap{
        {
          "text":  "按钮1",
          "style": 1,
          "key":   "button_key_1",
        },
        {
          "text":  "按钮2",
          "style": 2,
          "key":   "button_key_2",
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

func APIMessageSendTemplateCardVoteInteraction(c *gin.Context) {

  messages := &power.HashMap{
    "touser":  "UserID1|UserID2|UserID3",
    "toparty": "PartyID1 | PartyID2",
    "totag":   "TagID1 | TagID2",
    "msgtype": "template_card",
    "agentid": 1,
    "template_card": power.HashMap{
      "card_type": "vote_interaction",
      "source": power.HashMap{
        "icon_url": "图片的url",
        "desc":     "企业微信",
      },
      "main_title": power.HashMap{
        "title": "欢迎使用企业微信",
        "desc":  "您的好友正在邀请您加入企业微信",
      },
      "task_id": "task_id",
      "checkbox": power.HashMap{
        "question_key": "question_key1",
        "option_list": []power.HashMap{
          power.HashMap{
            "id":         "option_id1",
            "text":       "选择题选项1",
            "is_checked": true,
          },
          power.HashMap{
            "id":         "option_id2",
            "text":       "选择题选项2",
            "is_checked": false,
          },
        },
        "mode": 1,
      },
      "submit_button": power.HashMap{
        "text": "提交",
        "key":  "key",
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

func APIMessageSendTemplateCardMultipleInteraction(c *gin.Context) {

  messages := &power.HashMap{
    "touser":  "UserID1|UserID2|UserID3",
    "toparty": "PartyID1 | PartyID2",
    "totag":   "TagID1 | TagID2",
    "msgtype": "template_card",
    "agentid": 1,
    "template_card": power.HashMap{
      "card_type": "multiple_interaction",
      "source": power.HashMap{
        "icon_url": "图片的url",
        "desc":     "企业微信",
      },
      "main_title": power.HashMap{
        "title": "欢迎使用企业微信",
        "desc":  "您的好友正在邀请您加入企业微信",
      },
      "task_id": "task_id",
      "select_list": []power.HashMap{
        {
          "question_key": "question_key1",
          "title":        "选择器标签1",
          "selected_id":  "selection_id1",
          "option_list": []power.HashMap{
            power.HashMap{
              "id":   "selection_id1",
              "text": "选择器选项1",
            },
            power.HashMap{
              "id":   "selection_id2",
              "text": "选择器选项2",
            },
          },
        },
        {
          "question_key": "question_key2",
          "title":        "选择器标签2",
          "selected_id":  "selection_id3",
          "option_list": []power.HashMap{
            power.HashMap{
              "id":   "selection_id3",
              "text": "选择器选项3",
            },
            power.HashMap{
              "id":   "selection_id4",
              "text": "选择器选项4",
            },
          },
        },
      },
      "submit_button": power.HashMap{
        "text": "提交",
        "key":  "key",
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

//更新模版卡片消息
// https://work.weixin.qq.com/api/doc/90000/90135/94888
func APIMessageUpdateButton(c *gin.Context) {

  messages := &power.HashMap{
    "userids":       []string{"userid1", "userid2"},
    "partyids":      []int{2, 3},
    "tagids":        []int{44, 55},
    "atall":         0,
    "agentid":       1,
    "response_code": "response_code",
    "button": power.HashMap{
      "replace_name": "replace_name",
    },
  }

  res, err := services.WeComApp.Message.UpdateTemplateCard(messages)

  if err != nil {
    panic(err)
  }

  c.JSON(http.StatusOK, res)
}

func APIMessageUpdateTemplateCardTextNotice(c *gin.Context) {
  messages := &power.HashMap{
    "userids":       []string{"userid1", "userid2"},
    "partyids":      []int{2, 3},
    "agentid":       1,
    "response_code": "response_code",
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
        power.HashMap{
          "keyname": "邀请人",
          "value":   "张三",
        },
        power.HashMap{
          "type":    1,
          "keyname": "企业微信官网",
          "value":   "点击访问",
          "url":     "https://work.weixin.qq.com",
        },
        power.HashMap{
          "type":     2,
          "keyname":  "企业微信下载",
          "value":    "企业微信.apk",
          "media_id": "文件的media_id",
        },
      },
      "jump_list": []power.HashMap{
        power.HashMap{
          "type":  1,
          "title": "企业微信官网",
          "url":   "https://work.weixin.qq.com",
        },
        power.HashMap{
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
  }

  res, err := services.WeComApp.Message.UpdateTemplateCard(messages)

  if err != nil {
    panic(err)
  }

  c.JSON(http.StatusOK, res)
}

func APIMessageUpdateTemplateCardNewsNotice(c *gin.Context) {

  messages := &power.HashMap{
    "userids":       []string{"userid1", "userid2"},
    "partyids":      []int{2, 3},
    "agentid":       1,
    "response_code": "response_code",
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
  }

  res, err := services.WeComApp.Message.UpdateTemplateCard(messages)

  if err != nil {
    panic(err)
  }

  c.JSON(http.StatusOK, res)
}

func APIMessageUpdateTemplateCardButtonInteraction(c *gin.Context) {

  messages := &power.HashMap{
    "userids":       []string{"userid1", "userid2"},
    "partyids":      []int{2, 3},
    "agentid":       1,
    "response_code": "response_code",
    "template_card": power.HashMap{
      "card_type": "button_interaction",
      "source": power.HashMap{
        "icon_url": "图片的url",
        "desc":     "企业微信",
      },
      "main_title": power.HashMap{
        "title": "欢迎使用企业微信",
        "desc":  "您的好友正在邀请您加入企业微信",
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
      "card_action": power.HashMap{
        "type":     2,
        "url":      "https://work.weixin.qq.com",
        "appid":    "小程序的appid",
        "pagepath": "/index.html",
      },
      "button_list": []power.HashMap{
        {
          "text":  "按钮1",
          "style": 1,
          "key":   "button_key_1",
        },
        {
          "text":  "按钮2",
          "style": 2,
          "key":   "button_key_2",
        },
      },
      "replace_text": "已提交",
    },
  }
  res, err := services.WeComApp.Message.UpdateTemplateCard(messages)

  if err != nil {
    panic(err)
  }

  c.JSON(http.StatusOK, res)
}

func APIMessageUpdateTemplateCardVoteInteraction(c *gin.Context) {

  messages := &power.HashMap{
    "userids":       []string{"userid1", "userid2"},
    "partyids":      []int{2, 3},
    "agentid":       1,
    "response_code": "response_code",
    "template_card": power.HashMap{
      "card_type": "vote_interaction",
      "source": power.HashMap{
        "icon_url": "图片的url",
        "desc":     "企业微信",
      },
      "main_title": power.HashMap{
        "title": "欢迎使用企业微信",
        "desc":  "您的好友正在邀请您加入企业微信",
      },
      "checkbox": power.HashMap{
        "question_key": "question_key1",
        "option_list": []power.HashMap{
          {
            "id":         "option_id1",
            "text":       "选择题选项1",
            "is_checked": true,
          },
          {
            "id":         "option_id2",
            "text":       "选择题选项2",
            "is_checked": false,
          },
        },
        "disable": false,
        "mode":    1,
      },
      "submit_button": power.HashMap{
        "text": "提交",
        "key":  "key",
      },
      "replace_text": "已提交",
    },
  }
  res, err := services.WeComApp.Message.UpdateTemplateCard(messages)

  if err != nil {
    panic(err)
  }

  c.JSON(http.StatusOK, res)

}

func APIMessageUpdateTemplateCardMultipleInteraction(c *gin.Context) {
  messages := &power.HashMap{

    "userids":       []string{"userid1", "userid2"},
    "partyids":      []int{2, 3},
    "tagids":        []int{44, 55},
    "atall":         0,
    "agentid":       1,
    "response_code": "response_code",
    "template_card": power.HashMap{
      "card_type": "multiple_interaction",
      "source": power.HashMap{
        "icon_url": "图片的url",
        "desc":     "企业微信",
      },
      "main_title": power.HashMap{
        "title": "欢迎使用企业微信",
        "desc":  "您的好友正在邀请您加入企业微信",
      },
      "select_list": []power.HashMap{
        {
          "question_key": "question_key1",
          "title":        "选择器标签1",
          "selected_id":  "selection_id1",
          "disable":      false,
          "option_list": []power.HashMap{
            {
              "id":   "selection_id1",
              "text": "选择器选项1",
            },
            {
              "id":   "selection_id2",
              "text": "选择器选项2",
            },
          },
        },
        {
          "question_key": "question_key2",
          "title":        "选择器标签2",
          "selected_id":  "selection_id3",
          "disable":      false,
          "option_list": []power.HashMap{
            {
              "id":   "selection_id3",
              "text": "选择器选项3",
            },
            {
              "id":   "selection_id4",
              "text": "选择器选项4",
            },
          },
        },
      },
      "submit_button": power.HashMap{
        "text": "提交",
        "key":  "key",
      },
      "replace_text": "已提交",
    },
  }

  res, err := services.WeComApp.Message.UpdateTemplateCard(messages)

  if err != nil {
    panic(err)
  }

  c.JSON(http.StatusOK, res)
}
