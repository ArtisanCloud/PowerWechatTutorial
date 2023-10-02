package services

import (
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/response"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/miniProgram"
	"github.com/ArtisanCloud/PowerWeChat/v3/test/testLogDriver"
	"power-wechat-tutorial/config"
)

var MiniProgramApp *miniProgram.MiniProgram

const TIMEZONE = "asia/shanghai"
const DATETIME_FORMAT = "20060102"

func NewMiniMiniProgramService(conf *config.Configuration) (*miniProgram.MiniProgram, error) {
	var cache kernel.CacheInterface
	if conf.MiniProgram.RedisAddr != "" {
		cache = kernel.NewRedisClient(&kernel.RedisOptions{
			Addr: conf.MiniProgram.RedisAddr,
		})
	}
	app, err := miniProgram.NewMiniProgram(&miniProgram.UserConfig{
		AppID:        conf.MiniProgram.AppID,  // 小程序、公众号或者企业微信的appid
		Secret:       conf.MiniProgram.Secret, // 商户号 appID
		ResponseType: response.TYPE_MAP,
		Token:        conf.MiniProgram.MessageToken,
		AESKey:       conf.MiniProgram.MessageAesKey,

		AppKey:  conf.MiniProgram.VirtualPayAppKey,
		OfferID: conf.MiniProgram.VirtualPayOfferID,

		Log: miniProgram.Log{
			Driver: &testLogDriver.SimpleLogger{},
			Level:  "debug",
			File:   "./wechat.log",
		},
		//"sandbox": true,
		Cache:     cache,
		HttpDebug: true,
		Debug:     false,
	})

	return app, err
}
