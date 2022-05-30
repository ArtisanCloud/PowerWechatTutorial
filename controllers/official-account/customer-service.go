package official_account

import (
  "github.com/ArtisanCloud/PowerLibs/v2/object"
  "github.com/ArtisanCloud/PowerWeChat/v2/src/kernel/messages"
  "github.com/ArtisanCloud/PowerWeChat/v2/src/kernel/power"
  "github.com/ArtisanCloud/PowerWeChat/v2/src/officialAccount/customerService/request"
  "github.com/gin-gonic/gin"
  "net/http"
  "power-wechat-tutorial/services"
  "strconv"
)

// GetCustomerList 获取所有客服
func GetCustomerList(ctx *gin.Context) {
  data, err := services.OfficialAccountApp.CustomerService.List()
  if err != nil {
    ctx.JSON(http.StatusBadRequest, err)
    return
  }
  ctx.JSON(http.StatusOK, data)
}

// GetCustomerOnline 获取所有在线的客服
func GetCustomerOnline(ctx *gin.Context) {
  data, err := services.OfficialAccountApp.CustomerService.Online()
  if err != nil {
    ctx.JSON(http.StatusBadRequest, err)
    return
  }
  ctx.JSON(http.StatusOK, data)
}

// CustomerCreate 新建客服
func CustomerCreate(ctx *gin.Context) {
  account, _ := ctx.GetPostForm("account")
  nickname, _ := ctx.GetPostForm("nickname")

  data, err := services.OfficialAccountApp.CustomerService.Create(account, nickname)
  if err != nil {
    ctx.JSON(http.StatusBadRequest, err)
    return
  }
  ctx.JSON(http.StatusOK, data)
}

// CustomerUpdate 更新客服
func CustomerUpdate(ctx *gin.Context) {
  account, _ := ctx.GetPostForm("account")
  nickname, _ := ctx.GetPostForm("nickname")

  data, err := services.OfficialAccountApp.CustomerService.Update(account, nickname)
  if err != nil {
    ctx.JSON(http.StatusBadRequest, err)
    return
  }
  ctx.JSON(http.StatusOK, data)
}

// CustomerDelete 删除客服 err
func CustomerDelete(ctx *gin.Context) {
  account, _ := ctx.GetPostForm("account")
  data, err := services.OfficialAccountApp.CustomerService.Delete(account)
  if err != nil {
    ctx.JSON(http.StatusBadRequest, err)
    return
  }
  ctx.JSON(http.StatusOK, data)
}

// CustomerSetAvatar 设置客户头像
func CustomerSetAvatar(ctx *gin.Context) {
  account, _ := ctx.GetPostForm("account")
  avatarPath := "./resource/cloud.jpg"
  data, err := services.OfficialAccountApp.CustomerService.SetAvatar(account, avatarPath)
  if err != nil {
    ctx.JSON(http.StatusBadRequest, err)
    return
  }
  ctx.JSON(http.StatusOK, data)
}

// 获取客服与客户聊天记录
func CustomerMessages(ctx *gin.Context) {
  startTimeStr := ctx.Query("startTime") // eg: 2015-06-01
  endTimeStr := ctx.Query("endTime")

  startTime, _ := strconv.Atoi(startTimeStr)
  endTime, _ := strconv.Atoi(endTimeStr)
  msgId := 1
  number := 1000

  data, err := services.OfficialAccountApp.CustomerService.Messages(&request.RequestMessages{
    StartTime: startTime,
    EndTime:   endTime,
    MsgID:     msgId,
    Number:    number,
  })
  if err != nil {
    ctx.JSON(http.StatusBadRequest, err)
    return
  }
  ctx.JSON(http.StatusOK, data)
}

func CustomerMessageSend(ctx *gin.Context) {
  openID, _ := ctx.GetPostForm("openID")
  account, _ := ctx.GetPostForm("account")
  content, _ := ctx.GetPostForm("content")
  _ = content

  //msg := messages.NewText(content)
  msg := messages.NewVideo("RWhYObhxAi-gwnUf3ifNhcYpLJCXRbBW6Bd4n4cEdY32ksf_tCOxACYwbOHdn3bF", &power.HashMap{
    //"media_id": "UVv55kEBGGIsWD0__1DNy0c37swPHr_IOggttifIQUAmjuNaFKehwQYps8MeAdLW",
  })

  result, err := services.OfficialAccountApp.CustomerService.Message(msg).From(account).SetTo(openID).Send()
  if err != nil {
    ctx.JSON(http.StatusBadRequest, err)
    return
  }
  ctx.JSON(http.StatusOK, result)
}
func CustomerMessageSendText(ctx *gin.Context) {
  openID, _ := ctx.GetPostForm("openID")
  account, _ := ctx.GetPostForm("account")
  content, _ := ctx.GetPostForm("content")
  _ = content

  msg := messages.NewText(content)

  result, err := services.OfficialAccountApp.CustomerService.
    Message(msg).
    From(account).
    SetTo(openID).
    Send()
  if err != nil {
    ctx.JSON(http.StatusBadRequest, err)
    return
  }
  ctx.JSON(http.StatusOK, result)
}
func CustomerMessageSendImage(ctx *gin.Context) {
  openID, _ := ctx.GetPostForm("openID")
  account, _ := ctx.GetPostForm("account")
  mediaID, _ := ctx.GetPostForm("mediaID")

  msg := messages.NewImage(mediaID, &power.HashMap{})

  result, err := services.OfficialAccountApp.CustomerService.
    Message(msg).
    From(account).
    SetTo(openID).
    Send()
  if err != nil {
    ctx.JSON(http.StatusBadRequest, err)
    return
  }
  ctx.JSON(http.StatusOK, result)
}

func CustomerMessageSendVoice(ctx *gin.Context) {
  openID, _ := ctx.GetPostForm("openID")
  account, _ := ctx.GetPostForm("account")
  mediaID, _ := ctx.GetPostForm("mediaID")

  msg := messages.NewVoice(mediaID, &power.HashMap{})

  result, err := services.OfficialAccountApp.CustomerService.
    Message(msg).
    From(account).
    SetTo(openID).
    Send()
  if err != nil {
    ctx.JSON(http.StatusBadRequest, err)
    return
  }
  ctx.JSON(http.StatusOK, result)
}

func CustomerMessageSendVideo(ctx *gin.Context) {
  openID, _ := ctx.GetPostForm("openID")
  account, _ := ctx.GetPostForm("account")
  mediaID, _ := ctx.GetPostForm("mediaID")

  msg := messages.NewVideo(mediaID, &power.HashMap{
    "title":          "test title",
    "description":    "test desc...",
    "thumb_media_id": "test thumb media id",
  })

  result, err := services.OfficialAccountApp.CustomerService.
    Message(msg).
    From(account).
    SetTo(openID).
    Send()
  if err != nil {
    ctx.JSON(http.StatusBadRequest, err)
    return
  }
  ctx.JSON(http.StatusOK, result)
}

func CustomerMessageSendLink(ctx *gin.Context) {
  openID, _ := ctx.GetPostForm("openID")
  account, _ := ctx.GetPostForm("account")

  msg := messages.NewLink(&power.HashMap{
    "title":       "ArtisanCloud",
    "description": "desc...",
    "url":         "https://www.artisan-cloud.com",
    "picurl":      "https://powerwechat.artisan-cloud.com/images/icons/favicon-32x32.png",
  })

  result, err := services.OfficialAccountApp.CustomerService.
    Message(msg).
    From(account).
    SetTo(openID).
    Send()
  if err != nil {
    ctx.JSON(http.StatusBadRequest, err)
    return
  }
  ctx.JSON(http.StatusOK, result)
}

func CustomerMessageSendMusic(ctx *gin.Context) {
  openID, _ := ctx.GetPostForm("openID")
  account, _ := ctx.GetPostForm("account")

  msg := messages.NewMusic(&power.HashMap{
    "title":          "MUSIC_TITLE",
    "description":    "MUSIC_DESCRIPTION",
    "musicurl":       "MUSIC_URL",
    "hqmusicurl":     "HQ_MUSIC_URL",
    "thumb_media_id": "THUMB_MEDIA_ID",
  })

  result, err := services.OfficialAccountApp.CustomerService.
    Message(msg).
    From(account).
    SetTo(openID).
    Send()
  if err != nil {
    ctx.JSON(http.StatusBadRequest, err)
    return
  }
  ctx.JSON(http.StatusOK, result)
}

func CustomerMessageSendNews(ctx *gin.Context) {
  openID, _ := ctx.GetPostForm("openID")
  account, _ := ctx.GetPostForm("account")

  msg := messages.NewNews(&object.HashMap{
    "title":          "MUSIC_TITLE",
    "description":    "MUSIC_DESCRIPTION",
    "musicurl":       "MUSIC_URL",
    "hqmusicurl":     "HQ_MUSIC_URL",
    "thumb_media_id": "THUMB_MEDIA_ID",
  })

  result, err := services.OfficialAccountApp.CustomerService.
    Message(msg).
    From(account).
    SetTo(openID).
    Send()
  if err != nil {
    ctx.JSON(http.StatusBadRequest, err)
    return
  }
  ctx.JSON(http.StatusOK, result)
}

func CustomerMessageSendRaw(ctx *gin.Context) {
  openID, _ := ctx.GetPostForm("openID")
  account, _ := ctx.GetPostForm("account")
  content, _ := ctx.GetPostForm("content")

  msg := messages.NewRaw(`
    {
      "touser":"` + openID + `",
      "msgtype":"text",
      "text":{"content":"` + content + `"}"}}
  `)

  result, err := services.OfficialAccountApp.CustomerService.
    Message(msg).
    From(account).
    SetTo(openID).
    Send()
  if err != nil {
    ctx.JSON(http.StatusBadRequest, err)
    return
  }
  ctx.JSON(http.StatusOK, result)
}

func CustomerInvite(ctx *gin.Context) {
  account, _ := ctx.GetPostForm("account")
  wechatID, _ := ctx.GetPostForm("wechatID")
  data, err := services.OfficialAccountApp.CustomerService.Invite(account, wechatID)
  if err != nil {
    ctx.JSON(http.StatusBadRequest, err)
    return
  }
  ctx.JSON(http.StatusOK, data)
}

func CustomerSessionCreate(ctx *gin.Context) {
  account, _ := ctx.GetPostForm("account")
  openID, _ := ctx.GetPostForm("openID")
  data, err := services.OfficialAccountApp.CustomerServiceSession.Create(account, openID)
  if err != nil {
    ctx.JSON(http.StatusBadRequest, err)
    return
  }
  ctx.JSON(http.StatusOK, data)
}

func CustomerSessionClose(ctx *gin.Context) {
  account, _ := ctx.GetPostForm("account")
  openID, _ := ctx.GetPostForm("openID")
  data, err := services.OfficialAccountApp.CustomerServiceSession.Close(account, openID)
  if err != nil {
    ctx.JSON(http.StatusBadRequest, err)
    return
  }
  ctx.JSON(http.StatusOK, data)
}
func GetCustomerSession(ctx *gin.Context) {
  openID := ctx.Query("openID")
  data, err := services.OfficialAccountApp.CustomerServiceSession.Get(openID)
  if err != nil {
    ctx.JSON(http.StatusBadRequest, err)
    return
  }
  ctx.JSON(http.StatusOK, data)
}

// CustomerSessionList err
func CustomerSessionList(ctx *gin.Context) {
  account := ctx.Query("account")
  data, err := services.OfficialAccountApp.CustomerServiceSession.List(account)
  if err != nil {
    ctx.JSON(http.StatusBadRequest, err)
    return
  }
  ctx.JSON(http.StatusOK, data)
}

func CustomerSessionWaiting(ctx *gin.Context) {
  data, err := services.OfficialAccountApp.CustomerServiceSession.Waiting()
  if err != nil {
    ctx.JSON(http.StatusBadRequest, err)
    return
  }
  ctx.JSON(http.StatusOK, data)
}
