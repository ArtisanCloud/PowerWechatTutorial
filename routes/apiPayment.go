package routes

import (
  "github.com/gin-gonic/gin"
  "power-wechat-tutorial/controllers/payment"
  "power-wechat-tutorial/controllers/payment/redpack"
)

func InitPaymentAPIRoutes(r *gin.Engine) {

  r.Static("/wx/payment", "./web")
  r.POST("/wx/notify", payment.CallbackWXNotify)
  apiRouterPayment := r.Group("/payment")
  {
    // Handle the pay route
    apiRouterPayment.GET("/order/make", payment.APIMakeOrder)
    apiRouterPayment.GET("/order/make/native", payment.APIMakeOrderNative)
    apiRouterPayment.GET("/order/make/app", payment.APIMakeOrderApp)
    apiRouterPayment.GET("/order/query", payment.APIQueryOrder)
    apiRouterPayment.GET("/order/close", payment.APICloseOrder)
    apiRouterPayment.GET("/order/refund", payment.APIRefundOrder)

    // Handle the bill route
    apiRouterPayment.GET("/bill/downloadURL", payment.APIBillDownloadURL)


    // Handle payment route
    apiRouterPayment.POST("redpack/sendNormal", redpack.APISendNormal)
    apiRouterPayment.POST("work/sendworkwxredpack", redpack.APIWorkSendWXRedpack)

  }

}
