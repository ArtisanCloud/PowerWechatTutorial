package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"power-wechat-tutorial/controllers"
	"power-wechat-tutorial/services"
)

func main() {

	PaymentService, err := services.NewWXPaymentService(nil)
	if err != nil || PaymentService == nil {
		panic(err)
	}

	r := gin.Default()

	r.GET("/order/make", controllers.APIMakeOrder)

	r.POST("/wx/notify", controllers.CallbackWXNotify)

	r.GET("/order/query", controllers.APIQueryOrder)

	r.GET("/order/close", controllers.APICloseOrder)

	r.Static("/wx/payment", "./web")

	log.Fatalln(r.Run(":8888"))

}
