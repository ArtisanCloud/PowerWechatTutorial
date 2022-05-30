package external_contact

import (
  "github.com/ArtisanCloud/PowerWeChat/v2/src/kernel/power"
  "github.com/ArtisanCloud/PowerWeChat/v2/src/work/externalContact/groupChat/request"
  "github.com/gin-gonic/gin"
  "net/http"
  "power-wechat-tutorial/services"
  "strconv"
)

// 获取客户群列表
// https://work.weixin.qq.com/api/doc/90000/90135/92120
func APIExternalContactGroupChatList(c *gin.Context) {

  options := &request.RequestGroupChatList{
    StatusFilter: 0,
    OwnerFilter: &power.HashMap{
      "userid_list": []string{"abel"},
    },
    Cursor: c.DefaultQuery("cursor", "r9FqSqsI8fgNbHLHE5QoCP50UIg2cFQbfma3l2QsmwI"),
    Limit:  10,
  }
  res, err := services.WeComContactApp.ExternalContactGroupChat.List(options)

  if err != nil {
    panic(err)
  }

  c.JSON(http.StatusOK, res)
}

// 获取客户群详情
// https://work.weixin.qq.com/api/doc/90000/90135/92122
func APIExternalContactGroupChatGet(c *gin.Context) {

  chatID := c.DefaultQuery("chatID", "wrOgQhDgAAMYQiS5ol9G7gK9JVAAAA")
  needName := c.DefaultQuery("needName", "true")
  bNeedName, _ := strconv.ParseBool(needName)

  res, err := services.WeComContactApp.ExternalContactGroupChat.Get(chatID, bNeedName)

  if err != nil {
    panic(err)
  }

  c.JSON(http.StatusOK, res)
}

// 客户群opengid转换
// https://work.weixin.qq.com/api/doc/90000/90135/94822
func APIExternalContactOpenGIDToChatID(c *gin.Context) {
  openID := c.DefaultQuery("openID", "oAAAAAAA")

  res, err := services.WeComContactApp.ExternalContactGroupChat.OpenGIDToChatID(openID)

  if err != nil {
    panic(err)
  }

  c.JSON(http.StatusOK, res)
}
