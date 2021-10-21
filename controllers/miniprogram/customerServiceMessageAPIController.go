package miniprogram

import (
  "github.com/ArtisanCloud/PowerWeChat/src/miniProgram/customerServiceMessage/request"
  "github.com/gin-gonic/gin"
  "io"
  "io/ioutil"
  "net/http"
  "power-wechat-tutorial/services"
)

// 发送客服文本消息给用户
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/customer-message/customerServiceMessage.send.html
func APICustomerServiceMessageSendText(c *gin.Context) {
  openID, exist := c.GetQuery("openID")
  if !exist {
    panic("parameter open id expected")
  }

  rs, err := services.MiniProgramApp.CustomerServiceMessage.SendText(
    openID,
    &request.CustomerServiceMsgText{
      Content: "Hello PowerWeChat",
    },
  )

  if err != nil {
    panic(err)
  }

  c.JSON(http.StatusOK, rs)
}
func APICustomerServiceMessageSendImage(c *gin.Context) {
  openID, exist := c.GetQuery("openID")
  if !exist {
    panic("parameter open id expected")
  }

  rs, err := services.MiniProgramApp.CustomerServiceMessage.SendImage(
    openID,
    &request.CustomerServiceMsgImage{
      MediaID: "MEDIA_ID",
    },
  )

  if err != nil {
    panic(err)
  }

  c.JSON(http.StatusOK, rs)
}
func APICustomerServiceMessageSendLink(c *gin.Context) {
  openID, exist := c.GetQuery("openID")
  if !exist {
    panic("parameter open id expected")
  }

  rs, err := services.MiniProgramApp.CustomerServiceMessage.SendLink(
    openID,
    &request.CustomerServiceMsgLink{
      Title:       "PowerWechat",
      Description: "PowerWechat description",
      Url:         "https://powerwechat.artisan-cloud.com",
      ThumbUrl:    "https://xxx.com/x.png",
    },
  )

  if err != nil {
    panic(err)
  }

  c.JSON(http.StatusOK, rs)
}
func APICustomerServiceMessageSendMpPage(c *gin.Context) {
  openID, exist := c.GetQuery("openID")
  if !exist {
    panic("parameter open id expected")
  }

  rs, err := services.MiniProgramApp.CustomerServiceMessage.SendMiniProgramPage(
    openID,
    &request.CustomerServiceMsgMpPage{
      Title:        "Hello PowerWechat",
      PagePath:     "/pages/index/index",
      ThumbMediaID: "thumb_media_id",
    },
  )

  if err != nil {
    panic(err)
  }

  c.JSON(http.StatusOK, rs)
}

// 下发客服当前输入状态给用户
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/customer-message/customerServiceMessage.setTyping.html
func APICustomerServiceMessageSetTyping(c *gin.Context) {

  openID, exist := c.GetQuery("openID")
  if !exist {
    panic("parameter open id expected")
  }
  command, exist := c.GetQuery("command")
  if !exist {
    panic("parameter command expected")
  }

  rs, err := services.MiniProgramApp.CustomerServiceMessage.SetTyping(openID, command)

  if err != nil {
    panic(err)
  }

  c.JSON(http.StatusOK, rs)

}

// 把媒体文件上传到微信服务器
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/customer-message/customerServiceMessage.uploadTempMedia.html
func APICustomerServiceMessageUploadTempMediaByFile(c *gin.Context) {

  mediaPath := "./resource/cloud.jpg"
  rs, err := services.MiniProgramApp.CustomerServiceMessage.UploadTempMedia("image", mediaPath, nil)

  if err != nil {
    panic(err)
  }

  c.JSON(http.StatusOK, rs)

}

// 把媒体文件上传到微信服务器
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/customer-message/customerServiceMessage.uploadTempMedia.html
func APICustomerServiceMessageUploadTempMediaByData(c *gin.Context) {

  var err error
  mediaPath := "./resource/cloud.jpg"
  value, err := ioutil.ReadFile(mediaPath)

  rs, err := services.MiniProgramApp.CustomerServiceMessage.UploadTempMedia("image", "", &power.HashMap{
    "name":  "cloud.jpg", // 请确保文件名有准确的文件类型
    "value": value,
  })

  if err != nil {
    panic(err)
  }

  c.JSON(http.StatusOK, rs)

}

// 把媒体文件上传到微信服务器
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/customer-message/customerServiceMessage.uploadTempMedia.html
func APICustomerServiceMessageGetTempMedia(c *gin.Context) {

  mediaID, exist := c.GetQuery("mediaID")
  if !exist {
    panic("parameter media id expected")
  }

  rs, err := services.MiniProgramApp.CustomerServiceMessage.GetTempMedia(mediaID)

  if err != nil {
    panic(err)
  }

  content, _ := ioutil.ReadAll(rs.Body)
  //fmt.Dump("content-type:",rs.Header.Get("Content-Type"))
  c.Header("Content-Type", rs.Header.Get("Content-Type"))
  c.Header("Content-Disposition", rs.Header.Get("attachment;filename=\""+rs.Header.Get("filename")+"\""))
  c.Data(http.StatusOK, rs.Header.Get("Content-Type"), content)
  io.Copy(c.Writer, rs.Body)
}
