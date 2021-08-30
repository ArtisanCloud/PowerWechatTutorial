package main

import (
	"github.com/ArtisanCloud/go-libs/object"
	"github.com/ArtisanCloud/power-wechat/src/payment"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
)

const TRANSACTION_SUCCESS = "TRANSACTION.SUCCESS"
const TRANSACTION_FAILED = "TRANSACTION.FAILED"

func NewWXPaymentService(r *http.Request) (*payment.Payment, error) {
	log.Printf("app_id: %s", os.Getenv("app_id"))
	Payment, err := payment.NewPayment(&payment.UserConfig{
		//"corp_id":        os.Getenv("corp_id"),
		//"secret":         os.Getenv("secret"),
		AppID:       os.Getenv("app_id"),  // 小程序、公众号或者企业微信的appid
		MchID:       os.Getenv("mch_id"),  // 商户号 appID
		MchApiV3Key: os.Getenv("mch_api_v3_key"), //
		Key:         os.Getenv("key"),
		CertPath:    os.Getenv("wx_cert_path"),
		KeyPath:     os.Getenv("wx_key_path"),
		SerialNo:    os.Getenv("serial_no"),

		ResponseType: os.Getenv("array"),
		Log: payment.Log{
			Level: "debug",
			File:  "./wechat.log",
		},
		Http: payment.Http{
			Timeout: 30.0,
			BaseURI: "https://api.mch.weixin.qq.com",
		},

		NotifyURL: os.Getenv("notify_url"),
		HttpDebug: true,
		//Debug: true,
		//"sandbox": true,

		// server config
		//Token:            os.Getenv("token"),
		//AESKey:           os.Getenv("aes_key"),
	}, r)

	return Payment, err
}

func main() {
	r := gin.Default()

	r.GET("/order/make", func(c *gin.Context) {
		paymentService, err := NewWXPaymentService(c.Request)
		if err != nil {
			panic(err)
		}

		// 下单
		response, err := paymentService.Order.Unify(&object.HashMap{
			"amount": &object.HashMap{
				"total":    1,
				"currency": "CNY",
			},
			"attach":       "自定义数据说明",
			"description":  "Image形象店-深圳腾大-QQ公仔",
			"mchid":        "1611854986",
			"notify_url":   "https://pay.wangchaoyi.com/wx/notify",
			"out_trade_no": "5519778939773395659222199111", // 这里是商户订单号，不能重复提交给微信
			"payer": &object.HashMap{
				"openid": "oAuaP0TRUMwP169nQfg7XCEAw3HQ",  // 用户的openid， 记得也是动态的。
			},
		}, false)

		if err != nil {
			log.Printf("error: %s", err)
			c.JSON(400, response)
			return
		}

		payConf, err := paymentService.JSSDK.BridgeConfig(response.PrepayID, true)
		if err != nil {
			panic(err)
		}


		c.JSON(200, payConf)
	})

	r.POST("/wx/notify", func(c *gin.Context) {
		paymentService, err := NewWXPaymentService(c.Request)
		if err != nil {
			panic(err)
		}
		//rs, err := paymentService.Order.QueryByOutTradeNumber("商户系统的内部订单号 [out_trade_no]")
		//rs, err := paymentService.Order.QueryByTransactionId("微信支付订单号 [transaction_id]")
		_, err = paymentService.HandlePaidNotify(func(message *object.HashMap, content *object.HashMap, fail string) interface{} {
			if content == nil || (*content)["out_trade_no"] == nil {
				return "no content notify"
			}
			// 看下支付通知事件状态
			if (*message)["event_type"].(string) != TRANSACTION_SUCCESS {
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
	})

	r.GET("/order/query", func(c *gin.Context) {
		paymentService, err := NewWXPaymentService(c.Request)
		traceNo := c.Query("traceNo")
		if err != nil {
			panic(err)
		}
		rs, err := paymentService.Order.QueryByOutTradeNumber(traceNo)
		c.JSON(200, rs)
	})

	r.GET("/order/close", func(c *gin.Context) {
		paymentService, err := NewWXPaymentService(c.Request)
		traceNo := c.Query("traceNo")
		log.Printf("traceNo: %s", traceNo)
		if err != nil {
			panic(err)
		}
		rs, err := paymentService.Order.Close(traceNo)
		if err != nil {
			log.Println("出错了： ", err)
			c.String(400, err.Error())
			return
		}
		c.JSON(200, rs)
	})

	r.Static("/wx/payment", "./web")

	log.Fatalln(r.Run(":8888"))

}