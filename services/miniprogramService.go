package services

import (
	"github.com/ArtisanCloud/power-wechat/src/miniProgram"
	"log"
	"os"
)

var AppMiniProgram *miniProgram.MiniProgram

const TIMEZONE = "asia/shanghai"

func NewMiniMiniProgramService() (*miniProgram.MiniProgram, error) {
	log.Printf("miniprogram app_id: %s", os.Getenv("miniprogram_app_id"))
	app, err := miniProgram.NewMiniProgram(&miniProgram.UserConfig{

		AppID:  os.Getenv("miniprogram_app_id"), // 小程序、公众号或者企业微信的appid
		Secret: os.Getenv("miniprogram_secret"), // 商户号 appID

		ResponseType: os.Getenv("array"),
		Log: miniProgram.Log{
			Level: "debug",
			File:  "./wechat.log",
		},

		HttpDebug: true,
		//Debug: true,
		//"sandbox": true,

	})

	return app, err
}
