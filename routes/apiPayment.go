package routes

import (
  "github.com/gin-gonic/gin"
  "power-wechat-tutorial/controllers/payment"
)

func InitPaymentAPIRoutes(r *gin.Engine) {

  r.Static("/wx/payment", "./web")
  r.POST("/wx/notify", payment.CallbackWXNotify)
  apiRouterPayment := r.Group("/payment")
  {
    // Handle the pay route
    apiRouterPayment.GET("/order/make", payment.APIMakeOrder)
    apiRouterPayment.GET("/order/query", payment.APIQueryOrder)
    apiRouterPayment.GET("/order/close", payment.APICloseOrder)

    // Handle the bill route
    apiRouterPayment.GET("/bill/downloadURL", payment.APIBillDownloadURL)

  }

}
