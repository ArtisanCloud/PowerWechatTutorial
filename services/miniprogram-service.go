package services

import (
	"github.com/ArtisanCloud/PowerLibs/v3/logger/drivers"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/response"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/miniProgram"
	"power-wechat-tutorial/config"
)

var MiniProgramApp *miniProgram.MiniProgram

const TIMEZONE = "asia/shanghai"
const DATETIME_FORMAT = "20060102"

func NewMiniMiniProgramService(conf *config.Configuration) (*miniProgram.MiniProgram, error) {
	var cache kernel.CacheInterface
	if conf.MiniProgram.RedisAddr != "" {
		cache = kernel.NewRedisClient(&kernel.UniversalOptions{
			Addrs: []string{conf.MiniProgram.RedisAddr},
			//Addrs: []string{
			//	"47.108.182.200:7000",
			//	"47.108.182.200:7001",
			//	"47.108.182.200:7002",
			//},
			//Username: "michaelhu",
			//Password: "111111",
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
		Http:    miniProgram.Http{},
		Log: miniProgram.Log{
			Driver: &drivers.SimpleLogger{},
			Level:  "debug",
			File:   "./wechat.log",
		},
		//"sandbox": true,
		StableTokenMode: true,
		ForceRefresh:    false,
		Cache:           cache,
		HttpDebug:       true,
		Debug:           false,
	})

	return app, err
}
