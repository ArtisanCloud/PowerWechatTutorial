package partner

import (
	"context"
	request2 "github.com/ArtisanCloud/PowerWeChat/v3/src/payment/partner/request"
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
		OutTradeNo:  "5519778939773395659222498002", // 这里是商户订单号，不能重复提交给微信
		Payer: &request2.JSAPIPayer{
			SpOpenId: "o4QEk5Mf1Do7utS7-SF5Go30s8i4", // 用户的openid， 记得也是动态的。
		},
	}

	// 如果需要覆盖掉全局的notify_url
	//options.SetNotifyUrl("https://pay.xxx.com/wx/notify")

	// 下单
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	response, err := services.PaymentApp.Partner.JSAPITransaction(ctx, options)

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

	response, err := services.PaymentApp.Partner.TransactionNative(c.Request.Context(), options)

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
	response, err := services.PaymentApp.Partner.TransactionApp(c.Request.Context(), options)

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

	rs, err := services.PaymentApp.Partner.QueryByOutTradeNumber(c.Request.Context(), traceNo)
	if err != nil {
		panic(err)
	}
	c.JSON(http.StatusOK, rs)

}

func APICloseOrder(c *gin.Context) {
	traceNo := c.Query("traceNo")
	log.Printf("traceNo: %s", traceNo)

	rs, err := services.PaymentApp.Partner.Close(c.Request.Context(), traceNo)
	if err != nil {
		log.Println("出错了： ", err)
		c.String(400, err.Error())
		return
	}
	c.JSON(http.StatusOK, rs)

}
