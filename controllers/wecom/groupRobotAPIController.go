package wecom

import (
  "github.com/ArtisanCloud/power-wechat/src/work/groupRobot/request"
  "github.com/gin-gonic/gin"
  "net/http"
  "power-wechat-tutorial/services"
)

// GroupRobotSendText 文本类型
func GroupRobotSendText(c *gin.Context) {
  key := "693axxx6-7aoc-4bc4-97a0-0ec2sifa5aaa"
  msg := &request.GroupRobotMsgText{
    Content:             "广州今日天气：29度，大部分多云，降雨概率：60%",
    MentionedList:       []string{"wangqing", "@all"},
    MentionedMobileList: []string{"13800001111", "@all"},
  }
  res, err := services.WeComApp.GroupRobot.SendText(key, msg)
  if err != nil {
    panic(err)
  }
  c.JSON(http.StatusOK, res)
}

// GroupRobotSendMarkdown markdown类型
func GroupRobotSendMarkdown(c *gin.Context) {
  key := "693axxx6-7aoc-4bc4-97a0-0ec2sifa5aaa"
  msg := &request.GroupRobotMsgMarkdown{
    Content: `实时新增用户反馈<font color=\"warning\">132例</font>，请相关同事注意。\n
    >类型:<font color=\"comment\">用户反馈</font>
    >普通用户反馈:<font color=\"comment\">117例</font>
    >VIP用户反馈:<font color=\"comment\">15例</font>`,
  }
  res, err := services.WeComApp.GroupRobot.SendMarkdown(key, msg)
  if err != nil {
    panic(err)
  }
  c.JSON(http.StatusOK, res)
}

// GroupRobotSendImage 图片类型
func GroupRobotSendImage(c *gin.Context) {
  key := "693axxx6-7aoc-4bc4-97a0-0ec2sifa5aaa"
  msg := &request.GroupRobotMsgImage{
    Base64: "DATA",
    Md5:    "MD5",
  }
  res, err := services.WeComApp.GroupRobot.SendImage(key, msg)
  if err != nil {
    panic(err)
  }
  c.JSON(http.StatusOK, res)
}

// GroupRobotSendNewsArticles 图文类型
func GroupRobotSendNewsArticles(c *gin.Context) {
  key := "693axxx6-7aoc-4bc4-97a0-0ec2sifa5aaa"
  msg := []*request.GroupRobotMsgNewsArticles{
    {
      Title:       "中秋节礼品领取",
      Description: "今年中秋节公司有豪礼相送",
      Url:         "www.qq.com",
      PicUrl:      "https://res.mail.qq.com/node/ww/wwopenmng/images/independent/doc/test_pic_msg1.png",
    },
  }
  res, err := services.WeComApp.GroupRobot.SendNewsArticles(key, msg)
  if err != nil {
    panic(err)
  }
  c.JSON(http.StatusOK, res)
}

// GroupRobotSendFile 文件
func GroupRobotSendFile(c *gin.Context) {
  key := "693axxx6-7aoc-4bc4-97a0-0ec2sifa5aaa"
  msg := &request.GroupRobotMsgFile{
    MediaID: "3a8asd892asd8asd",
  }
  res, err := services.WeComApp.GroupRobot.SendFile(key, msg)
  if err != nil {
    panic(err)
  }
  c.JSON(http.StatusOK, res)
}
func GroupRobotSendTemplateCard(c *gin.Context) {
  key := "693axxx6-7aoc-4bc4-97a0-0ec2sifa5aaa"
  msg := &request.GroupRobotMsgTemplateCard{
    CardType: "text_notice",
    Source: request.TemplateCardSource{
      IconUrl: "https://wework.qpic.cn/wwpic/252813_jOfDHtcISzuodLa_1629280209/0",
      Desc:    "企业微信",
    },
    MainTitle: request.TemplateCardMainTitle{
      Title: "欢迎使用企业微信",
      Desc:  "您的好友正在邀请您加入企业微信",
    },
    EmphasisContent: request.TemplateCardEmphasisContent{
      Title: "100",
      Desc:  "数据含义",
    },
    SubTitleText: "下载企业微信还能抢红包！",
    HorizontalContentList: []request.TemplateCardHorizontalContentListItem{
      {
        KeyName: "邀请人",
        Value:   "张三",
      },
      {
        KeyName: "企微官网",
        Value:   "点击访问",
        Type:    1,
        Url:     "https://work.weixin.qq.com/?from=openApi",
      },
      {
        KeyName: "企微下载",
        Value:   "企业微信.apk",
        Type:    2,
        MediaID: "MEDIAID",
      },
    },
    JumpList: []request.TemplateCardJumpListItem{
      {
        Type:  1,
        Url:   "https://work.weixin.qq.com/?from=openApi",
        Title: "企业微信官网",
      },
      {
        Type:     2,
        AppID:    "APPID",
        PagePath: "PAGEPATH",
        Title:    "跳转小程序",
      },
    },
    CardAction: request.TemplateCardCardAction{
      Type:     1,
      Url:      "https://work.weixin.qq.com/?from=openApi",
      AppID:    "APPID",
      PagePath: "PAGEPATH",
    },
  }
  res, err := services.WeComApp.GroupRobot.SendTemplateCard(key, msg)
  if err != nil {
    panic(err)
  }
  c.JSON(http.StatusOK, res)
}
