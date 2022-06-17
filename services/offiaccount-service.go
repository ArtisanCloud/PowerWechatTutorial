package services

import (
	"github.com/ArtisanCloud/PowerWeChat/v2/src/kernel"
	"github.com/ArtisanCloud/PowerWeChat/v2/src/officialAccount"
	"log"
	"os"
	"power-wechat-tutorial/config"
)

var OfficialAccountApp *officialAccount.OfficialAccount

func NewOfficialAccountAppService(conf *config.Configuration) (*officialAccount.OfficialAccount, error) {
	log.Printf("officialAccount app_id: %s", os.Getenv("miniprogram_app_id"))

	var cache kernel.CacheInterface
	if conf.MiniProgram.RedisAddr != "" {
		cache = kernel.NewRedisClient(&kernel.RedisOptions{
			Addr: conf.MiniProgram.RedisAddr,
		})
	}

	app, err := officialAccount.NewOfficialAccount(&officialAccount.UserConfig{

		AppID:  conf.OffiAccount.AppID,     // 小程序、公众号或者企业微信的appid
		Secret: conf.OffiAccount.AppSecret, // 商户号 appID

		Token:  "xxx",
		AESKey: "xxx",

		ResponseType: os.Getenv("response_type"),
		Log: officialAccount.Log{
			Level: "debug",
			File:  "./wechat.log",
		},
		Cache:     cache,
		HttpDebug: true,
		Debug:     false,
		//"sandbox": true,
	})

	return app, err
}
