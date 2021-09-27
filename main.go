package main

import (
  "github.com/gin-gonic/gin"
  "log"

  "power-wechat-tutorial/routes"

  "power-wechat-tutorial/services"
)

var Host = ""
var Port = "8888"

func main() {

  var err error
  services.PaymentApp, err = services.NewWXPaymentApp(nil)
  if err != nil || services.PaymentApp == nil {
    panic(err)
  }

  services.MiniProgramApp, err = services.NewMiniMiniProgramService()
  if err != nil || services.MiniProgramApp == nil {
    panic(err)
  }

  services.OfficialAccountApp, err = services.NewOfficialAccountAppService()
  if err != nil || services.OfficialAccountApp == nil {
    panic(err)
  }

  services.WeComApp, err = services.NewWeComService()
  if err != nil || services.WeComApp == nil {
    panic(err)
  }
  services.WeComContactApp, err = services.NewWeComContactService()
  if err != nil || services.WeComContactApp == nil {
    panic(err)
  }

  r := gin.Default()

  // Initialize the routes
  routes.InitializeRoutes(r)

  log.Fatalln(r.Run(Host + ":" + Port))

}
