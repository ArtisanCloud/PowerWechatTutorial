package accountService

import (
  "github.com/ArtisanCloud/power-wechat/src/kernel/power"
  "github.com/ArtisanCloud/power-wechat/src/work/accountService/customer/request"
  "github.com/gin-gonic/gin"
  "net/http"
  "power-wechat-tutorial/services"
)

// 读取消息
// https://work.weixin.qq.com/api/doc/90000/90135/94670
func APIAccountServiceCustomerBatchGet(c *gin.Context) {
  externalUserIDList := []string{c.DefaultQuery("externalUserIDList", "matrix-x")}

  res, err := services.WeComApp.AccountServiceCustomer.BatchGet(externalUserIDList)

  if err != nil {
    panic(err)
  }

  c.JSON(http.StatusOK, res)
}

// 获取配置的专员与客户群
// https://work.weixin.qq.com/api/doc/90000/90135/94674
func APIAccountServiceCustomerGetUpgradeServiceConfig(c *gin.Context) {
  res, err := services.WeComApp.AccountServiceCustomer.GetUpgradeServiceConfig()

  if err != nil {
    panic(err)
  }

  c.JSON(http.StatusOK, res)
}

// 为客户升级为专员或客户群服务
// https://work.weixin.qq.com/api/doc/90000/90135/94674
func APIAccountServiceCustomerUpgradeService(c *gin.Context) {
  options := &request.RequestUpgradeService{
    OpenKFID:       c.DefaultQuery("openKFID", "kfxxxxxxxxxxxxxx"),
    ExternalUserID: c.DefaultQuery("externalUserID", "kfxxxxxxxxxxxxxx"),
    Type:           2,
    Member: &power.StringMap{
      "userid":  c.DefaultQuery("member", "matrix-x"),
      "wording": "你好，我是你的专属服务专员zhangsan",
    },
    GroupChat: nil,
  }
  res, err := services.WeComApp.AccountServiceCustomer.UpgradeService(options)

  if err != nil {
    panic(err)
  }

  c.JSON(http.StatusOK, res)
}
