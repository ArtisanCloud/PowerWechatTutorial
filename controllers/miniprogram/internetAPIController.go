package miniprogram

import (
  "github.com/gin-gonic/gin"
  "net/http"
  "power-wechat-tutorial/services"
)

func APIInternetGetUserEncryptKey(c *gin.Context) {

  openID, exist := c.GetQuery("openID")
  if !exist {
    panic("parameter open id expected")
  }

  rs, err := services.MiniprogramApp.Internet.GetUserEncryptKey(openID, "", "hmac_sha256")

  if err != nil {
    panic(err)
  }

  c.JSON(http.StatusOK, rs)
}
