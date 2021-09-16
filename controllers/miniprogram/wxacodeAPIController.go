package miniprogram

import (
  "github.com/ArtisanCloud/power-wechat/src/kernel/power"
  "github.com/gin-gonic/gin"
  "io"
  "power-wechat-tutorial/services"
)

// 获取小程序二维码，适用于需要的码数量较少的业务场景
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/qr-code/wxacode.createQRCode.html
func APIWXACodeCreateQRCode(c *gin.Context) {

  path, exist := c.GetQuery("path")
  if !exist {
    panic("parameter path expected")
  }

  rs, err := services.MiniprogramApp.WXACode.CreateQRCode(path, 430)

  if err != nil {
    panic(err)
  }

  io.Copy(c.Writer, rs.Body)
}

// 获取小程序码，适用于需要的码数量较少的业务场景
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/qr-code/wxacode.get.html
func APIWXACodeGet(c *gin.Context) {

  path, exist := c.GetQuery("path")
  if !exist {
    panic("parameter path expected")
  }

  rs, err := services.MiniprogramApp.WXACode.Get(
    path,
    430,
    false,
    &power.HashMap{
      "r": 0,
      "g": 0,
      "b": 0,
    },
    false,
  )

  if err != nil {
    panic(err)
  }

  io.Copy(c.Writer, rs.Body)
}

// 获取小程序码，适用于需要的码数量极多的业务场景
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/qr-code/wxacode.getUnlimited.html
func APIWXACodeGetUnlimited(c *gin.Context) {

  page, exist := c.GetQuery("path")
  if !exist {
    panic("parameter page expected")
  }

  rs, err := services.MiniprogramApp.WXACode.GetUnlimited(
    "a=1",
    page,
    430,
    false,
    &power.HashMap{
      "r": 0,
      "g": 0,
      "b": 0,
    },
    false,
  )

  if err != nil {
    panic(err)
  }

  io.Copy(c.Writer, rs.Body)
}
