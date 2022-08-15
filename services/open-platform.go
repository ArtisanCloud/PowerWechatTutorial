package services

import (
	"github.com/ArtisanCloud/PowerWeChat/v2/src/kernel"
	"github.com/ArtisanCloud/PowerWeChat/v2/src/openPlatform"
	"log"
	"os"
	"power-wechat-tutorial/config"
)

var OpenPlatformApp *openPlatform.OpenPlatform

func NewOpenPlatformAppService(conf *config.Configuration) (*openPlatform.OpenPlatform, error) {
	log.Printf("officialAccount app_id: %s", os.Getenv("miniprogram_app_id"))

	var cache kernel.CacheInterface
	if conf.MiniProgram.RedisAddr != "" {
		cache = kernel.NewRedisClient(&kernel.RedisOptions{
			Addr: conf.MiniProgram.RedisAddr,
		})
	}

	app, err := openPlatform.NewOpenPlatform(&openPlatform.UserConfig{

		AppID:  conf.OpenPlatform.AppID,
		Secret: conf.OpenPlatform.AppSecret,

		Token:  conf.OpenPlatform.MessageToken,
		AESKey: conf.OpenPlatform.MessageAesKey,

		Log: openPlatform.Log{
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
