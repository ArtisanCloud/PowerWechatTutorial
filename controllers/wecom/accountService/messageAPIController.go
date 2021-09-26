package accountService

import (
  "github.com/ArtisanCloud/power-wechat/src/kernel/power"
  "github.com/gin-gonic/gin"
  "net/http"
  "power-wechat-tutorial/services"
)

// 读取消息
// https://work.weixin.qq.com/api/doc/90000/90135/94670
func APIAccountServiceSyncMsg(c *gin.Context) {

  cursor := c.DefaultQuery("cursor", "4gw7MepFLfgF2VC5npN")
  token := c.DefaultQuery("token", "ENCApHxnGDNAVNY4AaSJKj4Tb5mwsEMzxhFmHVGcra996NR")
  limit := 1000

  res, err := services.WeComApp.AccountServiceMessage.SyncMsg(cursor, token, limit)

  if err != nil {
    panic(err)
  }

  c.JSON(http.StatusOK, res)
}

// 发送消息
// https://work.weixin.qq.com/api/doc/90000/90135/90236
func APIAccountServiceSendMsg(c *gin.Context) {

  options := &power.HashMap{
    "touser":    c.DefaultQuery("toUser", "EXTERNAL_USERID"),
    "open_kfid": c.DefaultQuery("openKFID", "kfxxxxxxxxxxxxxx"),
    "msgid":     c.DefaultQuery("msgID", "MSGID"),
    "msgtype":   "text",
    "text": &power.HashMap{
      "content": "你购买的物品已发货，可点击链接查看物流状态http://work.weixin.qq.com/xxxxxx",
    },
  }

  res, err := services.WeComApp.AccountServiceMessage.SendMsg(options)

  if err != nil {
    panic(err)
  }

  c.JSON(http.StatusOK, res)
}

// 发送消息
// https://work.weixin.qq.com/api/doc/90000/90135/95122
func APIAccountServiceSendMsgOnEvent(c *gin.Context) {

  options := &power.HashMap{

    "code":    c.DefaultQuery("code", "CODE"),
    "msgid":   c.DefaultQuery("msgID", "MSG_ID"),
    "msgtype": "text",
    "text": &power.HashMap{
      "content": "欢迎咨询",
    },
  }

  res, err := services.WeComApp.AccountServiceMessage.SendMsgOnEvent(options)

  if err != nil {
    panic(err)
  }

  c.JSON(http.StatusOK, res)
}