package miniprogram

import (
  "github.com/ArtisanCloud/power-wechat/src/kernel/power"
  "github.com/gin-gonic/gin"
  "io/ioutil"
  "net/http"
  "power-wechat-tutorial/services"
)

// 本接口提供基于小程序的图片智能裁剪能力
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/img/img.aiCrop.html
func APIImgAICropByURL(c *gin.Context) {

  // https://cdn.pixabay.com/photo/2015/04/23/22/00/tree-736885_1280.jpg
  url, exist := c.GetQuery("url")
  if !exist {
    panic("parameter url expected")
  }

  rs, err := services.MiniprogramApp.Image.AICrop(url, nil)

  if err != nil {
    panic(err)
  }

  c.JSON(http.StatusOK, rs)
}

// 本接口提供基于小程序的图片智能裁剪能力
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/img/img.aiCrop.html
func APIImgAICropByData(c *gin.Context) {
  var err error
  mediaPath := "./resource/tree.png"
  value, err := ioutil.ReadFile(mediaPath)

  rs, err := services.MiniprogramApp.Image.AICrop("", &power.HashMap{
    "name":  "tree.png", // 请确保文件名有准确的文件类型
    "value": value,
  })

  if err != nil {
    panic(err)
  }

  c.JSON(http.StatusOK, rs)
}

// 本接口提供基于小程序的条码/二维码识别的API
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/img/img.scanQRCode.html
func APIImgScanQRCodeByURL(c *gin.Context) {
  url, exist := c.GetQuery("url")
  if !exist {
    panic("parameter url expected")
  }

  rs, err := services.MiniprogramApp.Image.ScanQRCode(url, nil)

  if err != nil {
    panic(err)
  }

  c.JSON(http.StatusOK, rs)
}

// 本接口提供基于小程序的条码/二维码识别的API
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/img/img.scanQRCode.html
func APIImgScanQRCodeByData(c *gin.Context) {
  var err error
  mediaPath := "./resource/qrcode.png"
  value, err := ioutil.ReadFile(mediaPath)

  rs, err := services.MiniprogramApp.Image.ScanQRCode("", &power.HashMap{
    "name":  "qrcode.png", // 请确保文件名有准确的文件类型
    "value": value,
  })

  if err != nil {
    panic(err)
  }

  c.JSON(http.StatusOK, rs)
}

// 本接口提供基于小程序的图片高清化能力
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/img/img.superresolution.html
func APIImgSuperResolutionByURL(c *gin.Context) {
  url, exist := c.GetQuery("url")
  if !exist {
    panic("parameter url expected")
  }

  rs, err := services.MiniprogramApp.Image.SuperResolution(url, nil)

  if err != nil {
    panic(err)
  }

  c.JSON(http.StatusOK, rs)
}

// 本接口提供基于小程序的图片高清化能力
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/img/img.superresolution.html
func APIImgSuperResolutionByData(c *gin.Context) {
  var err error
  mediaPath := "./resource/tree.png"
  value, err := ioutil.ReadFile(mediaPath)

  rs, err := services.MiniprogramApp.Image.SuperResolution("", &power.HashMap{
    "name":  "tree.png", // 请确保文件名有准确的文件类型
    "value": value,
  })

  if err != nil {
    panic(err)
  }

  c.JSON(http.StatusOK, rs)
}
