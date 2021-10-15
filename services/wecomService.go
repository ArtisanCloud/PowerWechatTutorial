package services

import (
	"github.com/ArtisanCloud/power-wechat/src/work"
	"log"
	"os"
)

var WeComApp *work.Work
var WeComContactApp *work.Work

func NewWeComService() (*work.Work, error) {
	log.Printf("wecom corp id: %s", os.Getenv("corp_id"))
	AgentID, _ := getEnvInt("wecom_agent_id")

	app, err := work.NewWork(&work.UserConfig{
		CorpID:      os.Getenv("corp_id"),                  // 企业微信的corp id，所有企业微信共用一个。
		AgentID:     AgentID,                                    // 内部应用的app id
		Secret:      os.Getenv("wecom_secret"),             // 内部应用的app secret
		Token:       os.Getenv("app_message_token"),        // 内部应用的app token
		AESKey:      os.Getenv("app_message_aes_key"),      // 内部应用的app aeskey
		CallbackURL: os.Getenv("app_message_callback_url"), // 内部应用的场景回调设置
		OAuth: work.OAuth{
			Callback: os.Getenv("app_oauth_callback_url"), // 内部应用的app oauth url
			Scopes:   []string{"snsapi_base"},
		},
		HttpDebug: true,
	})

	return app, err
}

func NewWeComContactService() (*work.Work, error) {
	log.Printf("wecom corp id: %s", os.Getenv("corp_id"))
	AgentID, _ := getEnvInt("wecom_agent_id")

	app, err := work.NewWork(&work.UserConfig{
		CorpID:      os.Getenv("corp_id"),              // 企业微信的app id，所有企业微信共用一个。
		AgentID:     AgentID,                           // 内部应用的app id
		Secret:      os.Getenv("contact_secret"),       // 内部应用的app secret
		Token:       os.Getenv("contact_token"),        // 内部应用的app token
		AESKey:      os.Getenv("contact_aes_key"),      // 内部应用的app aeskey
		CallbackURL: os.Getenv("contact_callback_url"), // 内部应用的场景回调设置
		OAuth: work.OAuth{
			Callback: os.Getenv("app_oauth_callback_url"), // 内部应用的app oauth url
			Scopes:   nil,
		},
		HttpDebug: true,
	})

	return app, err
}
