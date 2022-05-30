package miniprogram

import (
  "github.com/ArtisanCloud/PowerWeChat/v2/src/miniProgram/soter/request"
  "github.com/gin-gonic/gin"
  "net/http"
  "power-wechat-tutorial/services"
)

// SOTER 生物认证秘钥签名验证
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/soter/soter.verifySignature.html
func APISoterVerifySignature(c *gin.Context) {

  jsonString := c.Query("jsonString")
  jsonSignature := c.Query("jsonSignature")

  openID, exist := c.GetQuery("openID")
  if !exist {
    panic("parameter open id expected")
  }

  rs, err := services.MiniProgramApp.Soter.VerifySignature(&request.RequestSoter{
    OpenID:        openID,
    JsonString:    jsonString,
    JsonSignature: jsonSignature,
  })

  if err != nil {
    panic(err)
  }

  c.JSON(http.StatusOK, rs)
}
