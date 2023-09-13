package paymentScore

import (
	"github.com/ArtisanCloud/PowerWeChat/v3/src/payment/payScore/request"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"power-wechat-tutorial/services"
)

func APIServiceOrder(c *gin.Context) {
	//traceNo := c.Query("traceNo")
	//log.Printf("traceNo: %s", traceNo)

	para := &request.RequestServiceOrder{
		OutOrderNo:          "",
		Appid:               "",
		ServiceId:           "",
		ServiceIntroduction: "",
		PostPayments:        nil,
		PostDiscounts:       nil,
		//TimeRange:           nil,
		//Location:            nil,
		//RiskFund:            nil,
		Attach:    "",
		NotifyUrl: "",
		Openid:    "",
		//NeedUserConfirm:     nil,
	}

	rs, err := services.PaymentApp.PayScore.ServiceOrder(c.Request.Context(), para)
	if err != nil {
		log.Println("出错了： ", err)
		c.String(400, err.Error())
		return
	}
	c.JSON(http.StatusOK, rs)

}
