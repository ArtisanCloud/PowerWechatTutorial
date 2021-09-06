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

	// Payment App Router
	apiRouterPayment := r.Group("/payment")
	{
		apiRouterPayment.GET("/order/make", payment.APIMakeOrder)

		apiRouterPayment.POST("/wx/notify", payment.CallbackWXNotify)

		apiRouterPayment.GET("/order/query", payment.APIQueryOrder)

		apiRouterPayment.GET("/order/close", payment.APICloseOrder)

		apiRouterPayment.Static("/wx/payment", "./web")
	}

	// MiniProgram App Router
	routerMiniProgram := r.Group("/miniprogram")
	{
		// Handle the auth route
		routerMiniProgram.GET("/auth", miniprogram.APISNSSession)
		routerMiniProgram.GET("/auth/checkEncryptedData", miniprogram.APICheckEncryptedData)
		routerMiniProgram.GET("/auth/getPaidUnionID", miniprogram.APIGetPaidUnionID)

		// Handle the data cube analysis route
		routerMiniProgram.GET("/datacube/getDailyRetain", miniprogram.APIGetDailyRetain)
		routerMiniProgram.GET("/datacube/getMonthlyRetain", miniprogram.APIGetMonthlyRetain)
		routerMiniProgram.GET("/datacube/getWeeklyRetain", miniprogram.APIGetWeeklyRetain)
		routerMiniProgram.GET("/datacube/getDailySummary", miniprogram.APIGetDailySummary)

		routerMiniProgram.GET("/datacube/getDailyVisitTrend", miniprogram.APIGetDailyVisitTrend)
		routerMiniProgram.GET("/datacube/getMonthlyVisitTrend", miniprogram.APIGetMonthlyVisitTrend)
		routerMiniProgram.GET("/datacube/getWeeklyVisitTrend", miniprogram.APIGetWeeklyVisitTrend)
		routerMiniProgram.GET("/datacube/getPerformanceData", miniprogram.APIGetPerformanceData)
		routerMiniProgram.GET("/datacube/getUserPortrait", miniprogram.APIGetUserPortrait)
		routerMiniProgram.GET("/datacube/getVisitPage", miniprogram.APIGetVisitPage)



	}




	log.Fatalln(r.Run(Host + ":" + Port))

}
