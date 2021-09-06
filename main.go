package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"power-wechat-tutorial/controllers/miniprogram"
	"power-wechat-tutorial/controllers/payment"
	"power-wechat-tutorial/services"
)

var Host string = ""
var Port string = "8888"

func main() {

	var err error
	services.PaymentService, err = services.NewWXPaymentService(nil)
	if err != nil || services.PaymentService == nil {
		panic(err)
	}

	services.AppMiniProgram, err = services.NewMiniMiniProgramService()
	if err != nil || services.AppMiniProgram == nil {
		panic(err)
	}

	r := gin.Default()

	apiRouterPayment := r.Group("/payment")
	{
		apiRouterPayment.GET("/order/make", payment.APIMakeOrder)

		apiRouterPayment.POST("/wx/notify", payment.CallbackWXNotify)

		apiRouterPayment.GET("/order/query", payment.APIQueryOrder)

		apiRouterPayment.GET("/order/close", payment.APICloseOrder)

		apiRouterPayment.Static("/wx/payment", "./web")
	}


	apiRouterMiniprogram := r.Group("/miniprogram")
	{
		// Handle the index route
		apiRouterMiniprogram.GET("/auth", miniprogram.APISNSSession)
		//apiRouterMiniprogram.POST("/reservation/create", reservation.ValidateRequestMakeReservation, APIMakeReservation)

	}

	log.Fatalln(r.Run(Host + ":" + Port))

}
