package payment

import (
  "github.com/ArtisanCloud/power-wechat/src/kernel/models"
  "github.com/ArtisanCloud/power-wechat/src/kernel/power"
  "github.com/ArtisanCloud/power-wechat/src/payment/notify/request"
  "github.com/gin-gonic/gin"
  "log"
  "net/http"
  "power-wechat-tutorial/services"
)

func APIMakeOrder(c *gin.Context) {

  // 下单
  response, err := services.PaymentApp.Order.JSAPITransaction(&power.HashMap{
    "amount": &power.HashMap{
      "total":    1,
      "currency": "CNY",
    },
    "attach":       "自定义数据说明",
    "description":  "Image形象店-深圳腾大-QQ公仔",
    "mchid":        "1611854986",
    "notify_url":   "https://pay.wangchaoyi.com/wx/notify",
    "out_trade_no": "5519778939773395659222199111", // 这里是商户订单号，不能重复提交给微信
    "payer": &power.HashMap{
      "openid": "oAuaP0TRUMwP169nQfg7XCEAw3HQ", // 用户的openid， 记得也是动态的。
    },
  })

  if err != nil {

    log.Printf("error: %s", err)
    c.JSON(400, response)
    return
  }

  payConf, err := services.PaymentApp.JSSDK.BridgeConfig(response.PrepayID, true)
  if err != nil {
    panic(err)
  }

  c.JSON(200, payConf)
}

func APIQueryOrder(c *gin.Context) {

  traceNo := c.Query("traceNo")

  rs, err := services.PaymentApp.Order.QueryByOutTradeNumber(traceNo)
  if err != nil {
    panic(err)
  }
  c.JSON(http.StatusOK, rs)

}

func APICloseOrder(c *gin.Context) {
  traceNo := c.Query("traceNo")
  log.Printf("traceNo: %s", traceNo)

  rs, err := services.PaymentApp.Order.Close(traceNo)
  if err != nil {
    log.Println("出错了： ", err)
    c.String(400, err.Error())
    return
  }
  c.JSON(http.StatusOK, rs)

}

func CallbackWXNotify(c *gin.Context) {
  //rs, err := PaymentApp.Order.QueryByOutTradeNumber("商户系统的内部订单号 [out_trade_no]")
  //rs, err := PaymentApp.Order.QueryByTransactionId("微信支付订单号 [transaction_id]")
  response, err := services.PaymentApp.HandlePaidNotify(c.Request, func(message *request.RequestNotify, transaction *models.Transaction, fail func(message string)) interface{} {
    if transaction == nil || transaction.OutTradeNo == nil {
       fail("no content notify")
      return false
    }
    return true
  })

  if err != nil {
    log.Println("出错了： ", err)
    c.String(400, err.Error())
    return
  }

  response.Send(c.Writer)
}
