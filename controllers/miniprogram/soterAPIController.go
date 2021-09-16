package miniprogram

import (
  "github.com/gin-gonic/gin"
  "net/http"
  "power-wechat-tutorial/services"
)

// SOTER 生物认证秘钥签名验证
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/soter/soter.verifySignature.html
func APISoterVerifySignature(c *gin.Context) {

  openID, exist := c.GetQuery("openID")
  if !exist {
    panic("parameter open id expected")
  }

  rs, err := services.MiniProgramApp.Soter.VerifySignature(openID, "", "")

  if err != nil {
    panic(err)
  }

  c.JSON(http.StatusOK, rs)
}
