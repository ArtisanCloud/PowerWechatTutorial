package services

import (
	"github.com/ArtisanCloud/PowerLibs/v3/logger/drivers"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/work"
	"power-wechat-tutorial/config"
)

var WeComApp *work.Work
var WeComContactApp *work.Work

func NewWeComService(conf *config.Configuration) (*work.Work, error) {
	var cache kernel.CacheInterface
	if conf.WeCom.RedisAddr != "" {
		cache = kernel.NewRedisClient(&kernel.UniversalOptions{
			Addrs: []string{conf.MiniProgram.RedisAddr},
		})
	}

	app, err := work.NewWork(&work.UserConfig{
		CorpID:             conf.WeCom.CorpID,          // 企业微信的corp id，所有企业微信共用一个。
		AgentID:            conf.WeCom.AgentID,         // 内部应用的app id
		Secret:             conf.WeCom.Secret,          // 内部应用的app secret
		Token:              conf.WeCom.MessageToken,    // 内部应用的app token
		AESKey:             conf.WeCom.MessageAesKey,   // 内部应用的app aeskey
		CallbackURL:        conf.WeCom.MessageCallback, // 内部应用的场景回调设置
		FinanceSDKPath:     conf.WeCom.FinanceSDKPath,
		FinanceSDKPlatform: conf.WeCom.FinanceSDKPlatform,
		OAuth: work.OAuth{
			Callback: conf.WeCom.OAuthCallback, // 内部应用的app oauth url
			Scopes:   []string{"snsapi_base"},
		},
		Log: work.Log{
			Driver: &drivers.DummyLogger{},
			Level:  "debug",
			//Level: "off",
			File: "./wechat.log",
		},
		Cache:     cache,
		HttpDebug: true,
		Debug:     false,
	})

	return app, err
}

func NewWeComContactService(conf *config.Configuration) (*work.Work, error) {
	var cache kernel.CacheInterface
	if conf.MiniProgram.RedisAddr != "" {
		cache = kernel.NewRedisClient(&kernel.UniversalOptions{
			Addrs: []string{conf.MiniProgram.RedisAddr},
		})
	}

	app, err := work.NewWork(&work.UserConfig{
		CorpID:             conf.WeCom.CorpID,          // 企业微信的app id，所有企业微信共用一个。
		AgentID:            conf.WeCom.AgentID,         // 内部应用的app id
		Secret:             conf.WeCom.Secret,          // 内部应用的app secret
		Token:              conf.WeCom.ContactToken,    // 内部应用的app token
		AESKey:             conf.WeCom.ContactAESKey,   // 内部应用的app aeskey
		CallbackURL:        conf.WeCom.ContactCallback, // 内部应用的场景回调设置
		FinanceSDKPath:     conf.WeCom.FinanceSDKPath,
		FinanceSDKPlatform: conf.WeCom.FinanceSDKPlatform,

		OAuth: work.OAuth{
			Callback: conf.WeCom.OAuthCallback, // 内部应用的app oauth url
			Scopes:   nil,
		},
		HttpDebug: true,
		Debug:     false,
		Cache:     cache,
	})
	return app, err
}
