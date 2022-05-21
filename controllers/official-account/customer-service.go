package official_account

import (
  "github.com/ArtisanCloud/PowerWeChat/src/kernel/messages"
  "github.com/ArtisanCloud/PowerWeChat/src/officialAccount/customerService/request"
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
  //account, _ := ctx.GetPostForm("account")
  content, _ := ctx.GetPostForm("content")

  msg := messages.NewText(content)

  result, err := services.OfficialAccountApp.CustomerService.Message(msg).SetTo(openID).Send()
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
