package payment

import (
	"context"
	"github.com/ArtisanCloud/PowerLibs/v3/fmt"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/models"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/payment/notify/request"
	request2 "github.com/ArtisanCloud/PowerWeChat/v3/src/payment/order/request"
	request3 "github.com/ArtisanCloud/PowerWeChat/v3/src/payment/refund/request"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"power-wechat-tutorial/services"
	"time"
)

func APIMakeOrder(c *gin.Context) {
	options := &request2.RequestJSAPIPrepay{
		Amount: &request2.JSAPIAmount{
			Total:    1,
			Currency: "CNY",
		},
		Attach:      "自定义数据说明",
		Description: "Image形象店-深圳腾大-QQ公仔",
		OutTradeNo:  "55197789397733956592224980201", // 这里是商户订单号，不能重复提交给微信
		Payer: &request2.JSAPIPayer{
			OpenID: "o4QEk5Mf1Do7utS7-SF5Go30s8i4", // 用户的openid， 记得也是动态的。
		},
	}

	// 如果需要覆盖掉全局的notify_url
	//options.SetNotifyUrl("https://pay.xxx.com/wx/notify")

	// 下单
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	response, err := services.PaymentApp.Order.JSAPITransaction(ctx, options)

	if err != nil {
		log.Printf("error: %s", err)
		c.JSON(400, response)
		return
	}

	payConf, err := services.PaymentApp.JSSDK.BridgeConfig(response.PrepayID, false)
	if err != nil {
		panic(err)
		return
	}

	c.JSON(200, payConf)
}

// APIMakeOrderNative 生成Native支付二维码让用户扫
func APIMakeOrderNative(c *gin.Context) {
	options := &request2.RequestNativePrepay{
		Amount: &request2.NativeAmount{
			Total:    1,
			Currency: "CNY",
		},
		Attach:      "自定义数据说明",
		Description: "Image形象店-深圳腾大-QQ公仔",
		OutTradeNo:  "55197789397733956592225981111", // 这里是商户订单号，不能重复提交给微信
	}

	response, err := services.PaymentApp.Order.TransactionNative(c.Request.Context(), options)

	if err != nil {
		log.Printf("error: %s", err)
		c.JSON(400, response)
		return
	}

	c.JSON(200, response)
}

// APIMakeOrder App下单
func APIMakeOrderApp(c *gin.Context) {
	options := &request2.RequestAppPrepay{
		//TimeExpire: time.Now().Add(30).Format("2006-01-02T15:04:05+08:00"),
		Amount: &request2.AppAmount{
			Total:    1,
			Currency: "CNY",
		},
		Attach:      "自定义数据说明",
		Description: "Image形象店-深圳腾大-QQ公仔",
		OutTradeNo:  "5519778939773395659222498001", // 这里是商户订单号，不能重复提交给微信
	}

	// 如果需要覆盖掉全局的notify_url
	//options.SetNotifyUrl("https://pay.xxx.com/wx/notify")

	// 下单
	response, err := services.PaymentApp.Order.TransactionApp(c.Request.Context(), options)

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

	rs, err := services.PaymentApp.Order.QueryByOutTradeNumber(c.Request.Context(), traceNo)
	if err != nil {
		panic(err)
	}
	c.JSON(http.StatusOK, rs)

}

func APIRevertOrderByOutTradeNumber(c *gin.Context) {

	traceNo := c.Query("traceNo")

	rs, err := services.PaymentApp.Reverse.ByOutTradeNumber(c.Request.Context(), traceNo)
	if err != nil {
		panic(err)
	}
	c.JSON(http.StatusOK, rs)

}

func APICloseOrder(c *gin.Context) {
	traceNo := c.Query("traceNo")
	log.Printf("traceNo: %s", traceNo)

	rs, err := services.PaymentApp.Order.Close(c.Request.Context(), traceNo)
	if err != nil {
		log.Println("出错了： ", err)
		c.String(400, err.Error())
		return
	}
	c.JSON(http.StatusOK, rs)

}

// APIRefundOrder 退款
func APIRefundOrder(c *gin.Context) {
	transactionID := c.DefaultQuery("transactionID", "")
	outRefundNo := c.DefaultQuery("OutRefundNo", "")

	services.PaymentApp.GetConfig().GetString("app_id", "")

	options := &request3.RequestRefund{
		TransactionID: transactionID,
		OutRefundNo:   outRefundNo,
		Reason:        "",
		//NotifyUrl:     "", // 异步接收微信支付退款结果通知的回调地址
		FundsAccount: "",
		Amount: &request3.RefundAmount{
			Refund: 1, // 退款金额，单位：分
			Total:  1, // 订单总金额，单位：分
			From: []*request3.RefundAmountFrom{
				&request3.RefundAmountFrom{
					Account: "AVAILABLE",
					Amount:  1,
				},
			}, // 退款出资账户及金额。不传仍然需要这个空数组防止微信报错
			Currency: "CNY",
		},
		GoodsDetail: []*request3.RefundGoodDetail{
			&request3.RefundGoodDetail{
				MerchantGoodsID:  "1217752501201407033233368018",
				WechatPayGoodsID: "1001",
				GoodsName:        "公仔",
				UnitPrice:        1,
				RefundAmount:     1,
				RefundQuantity:   1,
			},
		},
	}

	rs, err := services.PaymentApp.Refund.Refund(c, options)
	if err != nil {
		panic(err)
	}
	c.JSON(http.StatusOK, rs)
}

func CallbackWXNotify(c *gin.Context) {
	//rs, err := PaymentApp.Order.QueryByOutTradeNumber("商户系统的内部订单号 [out_trade_no]")
	//rs, err := PaymentApp.Order.QueryByTransactionId("微信支付订单号 [transaction_id]")
	//_, err := services.PaymentApp.HandlePaidNotify(c.Request, func(chat-bot *power.HashMap, content *power.HashMap, fail string) interface{} {
	//if content == nil || (*content)["out_trade_no"] == nil {
	//  return fail("no content notify")
	//}
	//return true
	//})

	res, err := services.PaymentApp.HandlePaidNotify(
		c.Request,
		func(message *request.RequestNotify, transaction *models.Transaction, fail func(message string)) interface{} {

			// 看下支付通知事件状态
			// 这里可能是微信支付失败的通知，所以可能需要在数据库做一些记录，然后告诉微信我处理完成了。
			if message.EventType != "TRANSACTION.SUCCESS" {
				fmt.Dump(transaction)
				return true
			}

			if transaction.OutTradeNo != "" {
				// 这里对照自有数据库里面的订单做查询以及支付状态改变
				log.Printf("订单号：%s 支付成功", transaction.OutTradeNo)
			} else {
				// 因为微信这个回调不存在订单号，所以可以告诉微信我还没处理成功，等会它会重新发起通知
				// 如果不需要，直接返回true即可
				fail("payment fail")
				return nil
			}
			return true
		},
	)

	// 这里可能是因为不是微信官方调用的，无法正常解析出transaction和message，所以直接抛错。
	if err != nil {
		res.StatusCode = 502
		err = res.Write(c.Writer)
		return
		//panic(err)
	}

	// 这里根据之前返回的是true或者fail，框架这边自动会帮你回复微信
	err = res.Write(c.Writer)

	if err != nil {
		panic(err)
	}

}
