package miniprogram

import (
  "crypto/sha256"
  "fmt"
  "github.com/ArtisanCloud/PowerWeChat/src/miniProgram/base/request"
  "github.com/gin-gonic/gin"
  "log"
  "net/http"
  "power-wechat-tutorial/services"
)

func APISNSSession(c *gin.Context) {

  code, exist := c.GetQuery("code")
  if !exist {
    panic("parameter code expected")
  }

  rs, err := services.MiniProgramApp.Auth.Session(code)

  if err != nil {
    panic(err)
  }

  c.JSON(http.StatusOK, rs)

}

func APICheckEncryptedData(c *gin.Context) {
  encryptedData := c.DefaultQuery("encryptedData", "sTWzm26PrbsXlSA8AoW+GpiyNLJP0H5p2UT4dXKwLSvXv8aU4wIiJcZUcM/IzNXnoFtERY3BDRbZh6bwd0ZGENVhucqDPXmchTqseryIZnJiKsiNMHCpAkCA2Yl00q4UpOZYtGMuTX5BTuo1yB3bOOuIfDu6neHV3D158CofGB9m7TxFQ8A/JcauWzhvmEAPygfFaqCgDTEmluLu7S8wMA==")
  hashByte := sha256.Sum256([]byte(encryptedData))
  hash := hashByte[:]
  rs, err := services.MiniProgramApp.Base.CheckEncryptedData(fmt.Sprintf("%x", hash))

  if err != nil {
    panic(err)
  }

  c.JSON(http.StatusOK, rs)

}

func APIGetPaidUnionID(c *gin.Context) {
  openid := c.DefaultQuery("openid", "")
  log.Printf("openid: %s\n", openid)
  rs, err := services.MiniProgramApp.Base.GetPaidUnionID(&request.RequestGetPaidUnionID{
    OpenID:        openid,
    // TransactionID: "",
    // MchID:         "",
    // OutTradeNo:    "",
  })

  if err != nil {
    panic(err)
  }

  c.JSON(http.StatusOK, rs)

}
