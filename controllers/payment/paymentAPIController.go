package payment

import (
	"github.com/ArtisanCloud/power-wechat/src/kernel/power"
	"github.com/gin-gonic/gin"
	"log"
	"power-wechat-tutorial/services"
)

func APIMakeOrder(c *gin.Context) {


	// 下单
	response, err := services.PaymentService.Order.Unify(&power.HashMap{
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
	}, false)

	if err != nil {
		log.Printf("error: %s", err)
		c.JSON(400, response)
		return
	}

	payConf, err := services.PaymentService.JSSDK.BridgeConfig(response.PrepayID, true)
	if err != nil {
		panic(err)
	}

	c.JSON(200, payConf)
}

func APIQueryOrder(c *gin.Context) {

	traceNo := c.Query("traceNo")

	rs, err := services.PaymentService.Order.QueryByOutTradeNumber(traceNo)
	if err != nil {
		panic(err)
	}
	c.JSON(200, rs)

}

func APICloseOrder(c *gin.Context) {
	traceNo := c.Query("traceNo")
	log.Printf("traceNo: %s", traceNo)

	rs, err := services.PaymentService.Order.Close(traceNo)
	if err != nil {
		log.Println("出错了： ", err)
		c.String(400, err.Error())
		return
	}
	c.JSON(200, rs)

}

func CallbackWXNotify(c *gin.Context) {

	//rs, err := paymentService.Order.QueryByOutTradeNumber("商户系统的内部订单号 [out_trade_no]")
	//rs, err := paymentService.Order.QueryByTransactionId("微信支付订单号 [transaction_id]")
	_, err := services.PaymentService.HandlePaidNotify(c.Request, func(message *power.HashMap, content *power.HashMap, fail string) interface{} {
		if content == nil || (*content)["out_trade_no"] == nil {
			return "no content notify"
		}
		// 看下支付通知事件状态
		if (*message)["event_type"].(string) != services.TRANSACTION_SUCCESS {
			// 这里可能是微信支付失败的通知，所以可能需要在数据库做一些记录，然后告诉微信我处理完成了。
			return true
		}

		// 查询商户订单号
		orderNO := (*content)["out_trade_no"].(string)
		if orderNO != "" {
			// 这里对照自有数据库里面的订单做查询以及支付状态改变
			log.Printf("订单号：%s 支付成功", orderNO)
		} else {
			// 告诉微信我还没处理成功，等会它会重新发起通知
			// 如果不需要，直接返回true即可
			return "payment fail"
		}

		return true
	})

	if err != nil {
		panic(err)
	}

	c.String(200, "")

}
