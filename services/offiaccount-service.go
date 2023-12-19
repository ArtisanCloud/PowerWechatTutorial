package services

import (
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/officialAccount"
	"github.com/ArtisanCloud/PowerWeChat/v3/test/testLogDriver"
	"os"
	"power-wechat-tutorial/config"
)

var OfficialAccountApp *officialAccount.OfficialAccount

func NewOfficialAccountAppService(conf *config.Configuration) (*officialAccount.OfficialAccount, error) {

	var cache kernel.CacheInterface
	if conf.MiniProgram.RedisAddr != "" {
		cache = kernel.NewRedisClient(&kernel.RedisOptions{
			Addr: conf.MiniProgram.RedisAddr,
		})
	}

	app, err := officialAccount.NewOfficialAccount(&officialAccount.UserConfig{

		AppID:  conf.OffiAccount.AppID,     // 小程序、公众号或者企业微信的appid
		Secret: conf.OffiAccount.AppSecret, // 商户号 appID

		Token:        conf.OffiAccount.MessageToken,
		AESKey:       conf.OffiAccount.MessageAesKey,
		ResponseType: os.Getenv("response_type"),
		Log: officialAccount.Log{
			Driver: &testLogDriver.SimpleLogger{},
			Level:  "error",
			File:   "./wechat.log",
		},
		Cache:     cache,
		HttpDebug: true,
		Debug:     false,
		//"sandbox": true,
	})

	return app, err
}
