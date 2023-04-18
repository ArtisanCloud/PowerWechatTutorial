package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"power-wechat-tutorial/config"
	"power-wechat-tutorial/routes"
	"power-wechat-tutorial/services"
)

var Host = ""
var Port = "8887"

func main() {
	conf := config.Get()

	var err error
	services.PaymentApp, err = services.NewWXPaymentApp(conf)
	if err != nil || services.PaymentApp == nil {
		panic(err)
	}

	services.MiniProgramApp, err = services.NewMiniMiniProgramService(conf)
	if err != nil || services.MiniProgramApp == nil {
		panic(err)
	}

	services.OfficialAccountApp, err = services.NewOfficialAccountAppService(conf)
	if err != nil || services.OfficialAccountApp == nil {
		panic(err)
	}

	services.WeComApp, err = services.NewWeComService(conf)
	if err != nil || services.WeComApp == nil {
		panic(err)
	}

	services.WeComContactApp, err = services.NewWeComContactService(conf)
	if err != nil || services.WeComContactApp == nil {
		panic(err)
	}

	services.OpenPlatformApp, err = services.NewOpenPlatformAppService(conf)
	if err != nil || services.OpenPlatformApp == nil {
		panic(err)
	}

	services.RobotChatApp, err = services.NewRobotChatService(conf)
	if err != nil || services.RobotChatApp == nil {
		panic(err)
	}

	r := gin.Default()

	// Initialize the routes
	routes.InitializeRoutes(r)

	log.Fatalln(r.Run(Host + ":" + Port))

}
