package services

import (
  "github.com/ArtisanCloud/power-wechat/src/work"
  "log"
  "os"
)

var WeComApp *work.Work
var WeComContactApp *work.Work


func NewWeComService() (*work.Work, error) {
  log.Printf("wecom corp id: %s", os.Getenv("app_id"))
  AgentID, _ := getEnvInt("wecom_agent_id")

  app, err := work.NewWork(&work.UserConfig{
    CorpID: os.Getenv("app_id"), // 企业微信的app id，所有企业微信共用一个。
    AgentID:  AgentID, // 内部应用的app id
    Secret: os.Getenv("wecom_secret"), // 内部应用的app secret
    OAuth: work.OAuth{
      Callback: "https://wecom.artisan-cloud.com/callback", //
      Scopes:   nil,
    },
    HttpDebug: true,
  })

  return app, err
}

func NewWeComContactService() (*work.Work, error) {
  log.Printf("wecom corp id: %s", os.Getenv("app_id"))
  AgentID, _ := getEnvInt("wecom_agent_id")

  app, err := work.NewWork(&work.UserConfig{
    CorpID: os.Getenv("app_id"), // 企业微信的app id，所有企业微信共用一个。
    AgentID:  AgentID, // 内部应用的app id
    Secret: os.Getenv("wecom_contact_secret"), // 内部应用的app secret
    OAuth: work.OAuth{
      Callback: "https://wecom.artisan-cloud.com/callback", //
      Scopes:   nil,
    },
    HttpDebug: true,
  })

  return app, err
}


