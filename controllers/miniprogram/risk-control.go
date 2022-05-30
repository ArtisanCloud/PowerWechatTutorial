package miniprogram

import (
  "github.com/ArtisanCloud/PowerWeChat/v2/src/miniProgram/riskControl/request"
  "github.com/gin-gonic/gin"
  "net/http"
  "power-wechat-tutorial/services"
)

// 根据提交的用户信息数据获取用户的安全等级 risk_rank，无需用户授权。
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/safety-control-capability/riskControl.getUserRiskRank.html
func APIRiskControlGetUserRiskRank(c *gin.Context) {

  appID, exist := c.GetQuery("appID")
  if !exist {
    panic("parameter app id expected")
  }

  openID, exist := c.GetQuery("openID")
  if !exist {
    panic("parameter open id expected")
  }

  options := &request.RequestRiskControl{
    AppID:        appID,
    OpenID:       openID,
    Scene:        1,
    MobileNo:     "12345678",
    BankCardNo:   "******",
    CertNo:       "*******",
    ClientIP:     "******",
    EmailAddress: "***@qq.com",
    ExtendedInfo: "",
  }

  rs, err := services.MiniProgramApp.RiskControl.GetUserRiskRank(options)

  if err != nil {
    panic(err)
  }

  c.JSON(http.StatusOK, rs)
}
